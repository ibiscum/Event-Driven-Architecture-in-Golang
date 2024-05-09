package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/baskets"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/customers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/depot"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/config"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/logger"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/monolith"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/rpc"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/waiter"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/internal/web"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/notifications"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/ordering"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/payments"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter03/stores"
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
	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
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
	)

	return m.waiter.Wait()
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
