package db

import (
	"github.com/jinzhu/gorm"
	"github.com/kohkb/book_record/models"
	_ "github.com/mattn/go-sqlite3"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	db.LogMode(true)
	db.AutoMigrate((&models.Book{}))
	return db
}
