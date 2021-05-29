package dbutil

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mhewedy/go-conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func New() *gorm.DB {

	var (
		host     = conf.Get("db.host")
		port     = conf.GetInt("db.port")
		user     = conf.Get("db.user")
		password = conf.Get("db.password")
		dbname   = conf.Get("db.dbname")
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// see this wonderful blog https://www.alexedwards.net/blog/configuring-sqldb
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}
