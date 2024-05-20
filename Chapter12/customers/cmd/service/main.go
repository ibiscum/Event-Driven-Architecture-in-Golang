package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/customers/migrations"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/config"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/system"
	"github.com/ibiscum/Event-Driven-Architecture-in-Golang/Chapter12/internal/web"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("customers exitted abnormally: %s\n", err)
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
	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			return
		}
	}(s.DB())
	if err = s.MigrateDB(migrations.FS); err != nil {
		return err
	}
	s.Mux().Mount("/", http.FileServer(http.FS(web.WebUI)))
	// call the module composition root
	if err = customers.Root(s.Waiter().Context(), s); err != nil {
		return err
	}

	fmt.Println("started customers service")
	defer fmt.Println("stopped customers service")

	s.Waiter().Add(
		s.WaitForWeb,
		s.WaitForRPC,
		s.WaitForStream,
	)

	// go func() {
	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		m.logger.Debug().Msgf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	return s.Waiter().Wait()
}
