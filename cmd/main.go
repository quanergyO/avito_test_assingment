package main

import (
	"avito_test_assingment/internal/cache"
	"avito_test_assingment/internal/handler"
	"avito_test_assingment/internal/repository"
	"avito_test_assingment/internal/repository/postgres"
	"avito_test_assingment/internal/service"
	"avito_test_assingment/server"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// @title Avito Test Assignment Banner Service
// @version 1.0
// @description This is a sample server for Banner Service API.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		slog.Error("Error: loading env variables", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}

	redisDB, err := cache.NewRedis(cache.Config{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.DB"),
	})
	if err != nil {
		slog.Error("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisDB)
	handlers := handler.NewHandler(services)

	serv := new(server.Server)
	go func() {
		if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			slog.Error("Error: failed to start server on port:", viper.GetString("port"), err.Error())
			os.Exit(1)
		}
	}()

	slog.Info("Start server")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = serv.ShutDown(context.Background()); err != nil {
		slog.Error("error occured on server shutting down:", err)
		os.Exit(1)
	}

	if err = db.Close(); err != nil {
		slog.Error("error occured on close db connection:", err)
		os.Exit(1)
	}
	slog.Info("Server shutting down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
