package db

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var m sync.Mutex

func InitDB() {
	dsn := "postgres://facebook_user:facebook@localhost:5432/facebook"
	var err error
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	db.AutoMigrate(&SCHEDULE{})
}

func GetDB() *gorm.DB {
	m.Lock()
	return db
}

func ReleaseDB() {
	m.Unlock()
}

// ReorderRows shuffles the rows in the specified table
// O(N log N) ?
func ReorderRows(db *gorm.DB, table string) error {
	tmp := table + "_tmp"
	return db.Transaction(func(tx *gorm.DB) error {
		if err := db.Raw("CREATE TABLE ? SELECT * FROM ? ORDER BY RAND()", tmp, table).Error; err != nil {
			return err
		}
		if err := db.Raw("DELETE FROM ?", table).Error; err != nil {
			return err
		}
		if err := db.Raw("INSERT INTO ? SELECT * FROM ? ORDER BY RAND()", table, tmp).Error; err != nil {
			return err
		}
		if err := db.Raw("DROP TABLE ?", tmp).Error; err != nil {
			return err
		}
		return nil
	})
}