package main

import (
	"log"

	"github.com/spf13/viper"

	"todo2"
	"todo2/internal/handler"
	"todo2/internal/repository"
	"todo2/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error inisializtion %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo2.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
