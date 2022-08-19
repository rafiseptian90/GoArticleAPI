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
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(1)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}
