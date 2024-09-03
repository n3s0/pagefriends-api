package p

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
}

type AuthorModel struct {
	gorm.Model
	FirstName     string
	MiddleInitial string
	LastName      string
}

var db *gorm.DB

func SQLiteConnection() {
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

	db.AutoMigrate(&Book{})
}

func DbCreateBook(db *gorm.DB, BookModel *Book) (err error) {
	err = db.Create(BookModel).Error
	if err != nil {
		return err
	}
	return nil
}

func DbGetBooks(db *gorm.DB, BookModel *[]Book) (err error) {
	err = db.Find(BookModel).Error
	if err != nil {
		return err
	}

	return nil
}

func DbGetBook(db *gorm.DB, BookModel *Book, id int) (err error) {
	err = db.Where("id = ?", id).First(BookModel).Error
	if err != nil {
		return err
	}

	return nil
}

func DbUpdateBook(db *gorm.DB, BookModel *Book, id int) (err error) {
	db.Save(BookModel)
	return nil
}

func DbDeleteBook(db *gorm.DB, BookModel *Book, id int) (err error) {
	db.Where("id = ?", id).Delete(BookModel)
	return nil
}

func DbCreateAuthor(db *gorm.DB, AuthorModel *AuthorModel) (err error) {
	err = db.Create(AuthorModel).Error
	if err != nil {
		return err
	}
	return nil
}

func DbGetAuthors(db *gorm.DB, AuthorModel *[]AuthorModel) (err error) {
	err = db.Find(AuthorModel).Error
	if err != nil {
		return err
	}

	return nil
}

func DbGetAuthor(db *gorm.DB, AuthorModel *AuthorModel, id int) (err error) {
	err = db.Where("id = ?", id).First(AuthorModel).Error
	if err != nil {
		return err
	}

	return nil
}

func DbUpdateAuthor(db *gorm.DB, AuthorModel *AuthorModel, id int) (err error) {
	db.Save(AuthorModel)
	return nil
}

func DbDeleteAuthor(db *gorm.DB, AuthorModel *AuthorModel, id int) (err error) {
	db.Where("id = ?", id).Delete(AuthorModel)
	return nil
}
