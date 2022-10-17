package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/alexm24/ticket-booking/internal/config"
	"github.com/alexm24/ticket-booking/internal/handler"
	"github.com/alexm24/ticket-booking/internal/repository"
	"github.com/alexm24/ticket-booking/internal/server"
	"github.com/alexm24/ticket-booking/internal/service"
)

func App(configPath string) {
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Panicf("error read config: %s", err.Error())
	}
	log.Println("config read successfully")

	db, err := repository.NewPostgresDB(cfg.DBConfig)
	if err != nil {
		log.Panicf("failed to initialize postgres db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		log.Printf("start http server on port %s", cfg.Port)
		err = srv.Run(cfg.Port, handlers.InitRoutes(cfg.Path))
		if err != http.ErrServerClosed {
			log.Panicf("error occurred while running http server: %s", err.Error())
		}
	}()

	signalLisner := make(chan os.Signal, 1)
	signal.Notify(signalLisner,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	stop := <-signalLisner
	log.Printf("Shutting Down app: %s", stop)

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		log.Printf("error on db connection close: %s", err.Error())
	}
}
