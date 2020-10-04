package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kohkb/book_record/models"
)

type BooksHandler struct {
	Db *gorm.DB
}

func (h *BooksHandler) GetAll(c *gin.Context) {
	var books []models.Book
	h.Db.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"books": books,
	})
}

func (h *BooksHandler) NewBook(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{})
}

func (h *BooksHandler) CreateBook(c *gin.Context) {
	title, _ := c.GetPostForm("title")
	category, _ := c.GetPostForm("category")
	status, _ := c.GetPostForm("status")
	icategory, _ := strconv.ParseUint(category, 10, 32)
	istatus, _ := strconv.ParseUint(status, 10, 32)
	h.Db.Create(&models.Book{Title: title, Category: icategory, Status: istatus})
	c.Redirect(http.StatusMovedPermanently, "/books")
}

func (h *BooksHandler) EditBook(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"book": book,
	})
}
func (h *BooksHandler) UpdateBook(c *gin.Context) {
	Book := models.Book{}
	id := c.Param("id")
	title, _ := c.GetPostForm("title")
	category, _ := c.GetPostForm("category")
	status, _ := c.GetPostForm("status")
	icategory, _ := strconv.ParseUint(category, 10, 32)
	istatus, _ := strconv.ParseUint(status, 10, 32)
	h.Db.First(&Book, id)
	Book.Title = title
	Book.Category = icategory
	Book.Status = istatus
	h.Db.Save(&Book)
	c.Redirect(http.StatusMovedPermanently, "/books")
}
func (h *BooksHandler) DeleteBook(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	h.Db.Delete(&book)
	c.Redirect(http.StatusMovedPermanently, "/books")
}
