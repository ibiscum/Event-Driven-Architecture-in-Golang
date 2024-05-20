package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/baskets"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/cosec"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/depot"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/config"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/system"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/web"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/migrations"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/notifications"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/ordering"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/payments"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/search"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/stores"
)

type monolith struct {
	*system.System
	modules []system.Module
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("mallbots exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}
	s, err := system.NewSystem(cfg)
	if err != nil {
		return err
	}
	m := monolith{
		System: s,
		modules: []system.Module{
			&baskets.Module{},
			&customers.Module{},
			&depot.Module{},
			&notifications.Module{},
			&ordering.Module{},
			&payments.Module{},
			&stores.Module{},
			&cosec.Module{},
			&search.Module{},
		},
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.DB())
	err = m.MigrateDB(migrations.FS)
	if err != nil {
		return err
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.Mux().Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started mallbots application")
	defer fmt.Println("stopped mallbots application")

	m.Waiter().Add(
		m.WaitForWeb,
		m.WaitForRPC,
		m.WaitForStream,
	)

	// go func() {
	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		m.logger.Debug().Msgf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	return m.Waiter().Wait()
}

func (m *monolith) startupModules() error {
	for _, module := range m.modules {
		ctx := m.Waiter().Context()
		if err := module.Startup(ctx, m); err != nil {
			return err
		}
	}

	return nil
}
