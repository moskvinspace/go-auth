package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
}

func ConnectDB() {
	cfg := newConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.host,
		cfg.user,
		cfg.password,
		cfg.dbname,
		cfg.port,
		cfg.sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	DB = db
}

func defaultConfig() *dbConfig {
	return &dbConfig{
		host:     "localhost",
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		port:     "5432",
		sslmode:  "disable",
	}
}

func newConfig() *dbConfig {
	cfg := defaultConfig()

	host := os.Getenv("HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	switch {
	case host != "":
		cfg.host = host
		fallthrough
	case user != "":
		cfg.user = user
		fallthrough
	case password != "":
		cfg.password = password
		fallthrough
	case dbname != "":
		cfg.dbname = dbname
	}

	return cfg
}
