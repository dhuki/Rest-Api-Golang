package config

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type postgresServer interface {
	Start(env string) (*gorm.DB, error)
}

type postgres struct {
	logger log.Logger
}

func NewDatabase(logger log.Logger) postgresServer {
	return postgres{
		logger: logger,
	}
}

func (p postgres) Start(env string) (*gorm.DB, error) {
	var db *gorm.DB
	{
		err := godotenv.Load(env)
		if err != nil {
			return nil, err
		}

		dbHost := os.Getenv("db.host")
		dbPort := os.Getenv("db.port")
		dbName := os.Getenv("db.name")
		dbUsername := os.Getenv("db.username")
		dbPassword := os.Getenv("db.password")

		// by default sslmode=enable, so you have to connect with ssl
		// since your server doesn't provide it
		// just use sslmode=disable
		dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbName, dbPassword)
		db, err = gorm.Open("postgres", dbURI)
		if err != nil {
			return nil, err
		}
		db.LogMode(false)
	}
	return db, nil
}
