package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/nats-io/nats.go"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/baskets"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/cosec"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/customers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/depot"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/config"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/logger"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/rpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/waiter"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/internal/web"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/migrations"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/notifications"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/ordering"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/payments"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/search"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter10/stores"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	// parse config/env/...
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	// init infrastructure...
	// init db
	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)
	// migration database
	err = migrateDB(m.db)
	if err != nil {
		return err
	}
	// init nats & jetstream
	m.nc, err = nats.Connect(cfg.Nats.URL)
	if err != nil {
		return err
	}
	defer m.nc.Close()
	m.js, err = initJetStream(cfg.Nats, m.nc)
	if err != nil {
		return err
	}
	m.logger = initLogger(cfg)
	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		&baskets.Module{},
		&customers.Module{},
		&depot.Module{},
		&notifications.Module{},
		&ordering.Module{},
		&payments.Module{},
		&stores.Module{},
		&cosec.Module{},
		&search.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started mallbots application")
	defer fmt.Println("stopped mallbots application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
		m.waitForStream,
	)

	// go func() {
	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		m.logger.Debug().Msgf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	return m.waiter.Wait()
}

func migrateDB(db *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}

func initLogger(cfg config.AppConfig) zerolog.Logger {
	return logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}

func initJetStream(cfg config.NatsConfig, nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     cfg.Stream,
		Subjects: []string{fmt.Sprintf("%s.>", cfg.Stream)},
	})

	return js, err
}
