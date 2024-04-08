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
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

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

	if err != nil {
		slog.Info("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, redisDB)

	serv := new(server.Server)
	go func() {
		if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatal("Error: failed to start server on port: %s", viper.GetString("port"), err)
		}
	}()

	slog.Info("Start server")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := serv.ShutDown(context.Background()); err != nil {
		slog.Error("error occured on server shutting down: %s", err.Error())
		os.Exit(1)
	}

	if err := db.Close(); err != nil {
		slog.Error("error occured on close db connection: %s", err.Error())
		os.Exit(1)
	}
	slog.Info("Server shutting down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
