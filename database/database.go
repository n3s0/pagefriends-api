package database

import (
	"fmt"
	"os"

	"github.com/n3s0/pagefriends-api/p/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteConnection() *gorm.DB {
	path := "pagefriends.db"
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		fmt.Printf("%v\n", err)
		os.Exit(-1)
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	db.AutoMigrate(
		&model.Book{},
		&model.Author{},
	)

	return db
}
