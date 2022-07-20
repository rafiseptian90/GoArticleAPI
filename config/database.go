package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func DBConnection() (db *gorm.DB) {
	dsn := "host=" + os.Getenv("DB_HOST") + " " +
		"user=" + os.Getenv("DB_USERNAME") + " " +
		"password=" + os.Getenv("DB_PASSWORD") + " " +
		"dbname=" + os.Getenv("DB_NAME") + " " +
		"port=" + os.Getenv("DB_PORT") + " " +
		"sslmode=disable TimeZone=" + os.Getenv("APP_TIMEZONE")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	return
}
