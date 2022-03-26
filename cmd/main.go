package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pmokeev/covid-statistic/internal"
	"github.com/pmokeev/covid-statistic/internal/controller"
	"github.com/pmokeev/covid-statistic/internal/router"
	"github.com/pmokeev/covid-statistic/internal/service"
	"github.com/spf13/viper"
)

func initConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfigFile(); err != nil {
		log.Fatalf("Error while init config %s", err.Error())
	}

	service := service.NewService()
	controller := controller.NewController(service)
	router := router.NewRouter(controller)
	server := internal.NewServer()

	go func() {
		if err := server.Run(viper.GetString("port"), router.InitRouter()); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Listen: %s\n", err)
		}
	}()

	log.Println("API started")

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down API...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("API forced to shutdown:", err)
	}

	log.Println("API exiting")
}
