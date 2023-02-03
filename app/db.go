package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB     *sql.DB
	GormDB *gorm.DB
)

func connectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}

// InitDatabase For Vanilla SQL
func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connectionString())
	if err != nil {
		return nil, err
	}
	log.Infof("Database connection was created")

	return db, nil
}

func InitGormDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB, err = db.DB()
	if err != nil {
		return nil, err
	}

	GormDB = db
	return db, nil
}

func RunMigrations() error {
	err := goose.Up(DB, "./migrations")
	if err != nil {
		return err
	}
	return nil
}
