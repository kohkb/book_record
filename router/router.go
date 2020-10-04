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
	r.POST("/books", booksHandler.CreateTask)
	r.GET("/books/:id", booksHandler.EditTask)
	r.POST("/books/edit/:id", booksHandler.UpdateTask)
	r.POST("/books/delete/:id", booksHandler.DeleteTask)
	apiV1 := r.Group("/api/v1")
	{
		apiBooksHandler := v1.BooksHandler{
			Db: dbConn,
		}
		apiV1.GET("/books", apiBooksHandler.GetAll)
		apiV1.POST("/books", apiBooksHandler.CreateTask)
		apiV1.GET("/books/:id", apiBooksHandler.EditTask)
		apiV1.PUT("/books/:id", apiBooksHandler.UpdateTask)
		apiV1.DELETE("/books/:id", apiBooksHandler.DeleteTask)
	}

	r.Run()
}
