package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	v1 "github.com/kohkb/book_record/api/v1"
	"github.com/kohkb/book_record/controllers"
)

func Router(dbConn *gorm.DB) {
	booksHandler := controllers.BooksHandler{
		Db: dbConn,
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.LoadHTMLGlob("templates/*")
	r.GET("/books", booksHandler.GetAll)
	r.GET("/books/new", booksHandler.NewBook)
	r.POST("/books/new", booksHandler.CreateBook)
	r.GET("/books/edit/:id", booksHandler.EditBook)
	r.POST("/books/edit/:id", booksHandler.UpdateBook)
	r.POST("/books/delete/:id", booksHandler.DeleteBook)
	apiV1 := r.Group("/api/v1")
	{
		apiBooksHandler := v1.BooksHandler{
			Db: dbConn,
		}
		apiV1.GET("/books", apiBooksHandler.GetAll)
		apiV1.POST("/books", apiBooksHandler.CreateBook)
		apiV1.GET("/books/:id", apiBooksHandler.EditBook)
		apiV1.PUT("/books/:id", apiBooksHandler.UpdateBook)
		apiV1.DELETE("/books/:id", apiBooksHandler.DeleteBook)
	}

	r.Run()
}
