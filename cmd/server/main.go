package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	go_blog "go-blog"
	//_ "go-blog/cmd/server/docs"
	"go-blog/pkg/handler"
	"go-blog/pkg/repository"
	"go-blog/pkg/repository/postgres"
	"go-blog/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

//	@title Swagger Go blog API
// 	@version 1.0
// 	@description This is simple Blog-post REST API written in Go.
// 	@host localhost:8080
// 	@BasePath /

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in Header
//	@name Authorization
// 	@description Use Bearer token for authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("couldn't load config: %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("couldn't load env file: %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("couldn't connect to DB: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(go_blog.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("failed to run server: %s", err.Error())
		}
	}()

	logrus.Printf("Go blog app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("Go blog shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred while shutting down the server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred while closing db connection: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
