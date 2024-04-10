package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	bannerTable = "banners"
	userTable   = "users"
)

func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	err = makeMigration()
	slog.Info("Success connect to DB")
	return db, err
}

func makeMigration() error {
	slog.Info("Start Migrations")
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("Can't get current working directory")
		return err
	}
	migrationsPath := filepath.Join(wd, "schema")

	m, err := migrate.New(
		"file://"+migrationsPath,
		"postgres://postgres:qwewrty@localhost:5436/postgres?sslmode=disable")
	if err != nil {
		slog.Error("Can't connect to db")
		return err
	}
	err = m.Up()

	return nil
}
