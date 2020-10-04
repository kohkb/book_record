package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kohkb/book_record/controllers"
)

func Router(dbConn *gorm.DB) {
	booksHandler := controllers.BooksHandler{
		Db: dbConn,
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/books", booksHandler.GetAll)                 // 一覧画面
	r.POST("/books", booksHandler.CreateTask)            // 新規作成
	r.GET("/books/:id", booksHandler.EditTask)           // 編集画面
	r.POST("/books/edit/:id", booksHandler.UpdateTask)   // 更新
	r.POST("/books/delete/:id", booksHandler.DeleteTask) // 削除
	r.Run()
}
