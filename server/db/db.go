package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	// Carl's: "postgres://postgres:facebook@localhost:5432/postgres"
	// dsn := "postgres://postgres:facebook@localhost:5432/postgres"
	dsn := os.Getenv("POSTGRES_CONNECTION")

	var err error
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	// db.AutoMigrate(&SCHEDULE{})
	// db.AutoMigrate(&CHECKLIST{})

	db.AutoMigrate(&CHECKLIST{}, &SCHEDULE{})

}

func GetDB() *gorm.DB {
	return db
}
