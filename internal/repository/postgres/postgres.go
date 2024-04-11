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
	err = makeMigration(cfg)

	return db, err
}

func makeMigration(cfg Config) error {
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("Can't get current working directory")
		return err
	}
	migrationsPath := filepath.Join(wd, "schema")
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	m, err := migrate.New(
		"file://"+migrationsPath,
		databaseURL)
	if err != nil {
		slog.Error("Can't connect to db")
		return err
	}
	err = m.Up()

	return nil
}
