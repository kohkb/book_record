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

func (h *BooksHandler) CreateTask(c *gin.Context) {
	title, _ := c.GetPostForm("title")
	status, _ := c.GetPostForm("status")
	istatus, _ := strconv.ParseUint(status, 10, 32)
	h.Db.Create(&models.Book{Title: title, Status: istatus})
	c.Redirect(http.StatusMovedPermanently, "/books")
}

func (h *BooksHandler) EditTask(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"book": book,
	})
}
func (h *BooksHandler) UpdateTask(c *gin.Context) {
	Book := models.Book{}
	id := c.Param("id")
	title, _ := c.GetPostForm("title")
	status, _ := c.GetPostForm("status")
	istatus, _ := strconv.ParseUint(status, 10, 32)
	h.Db.First(&Book, id)
	Book.Title = title
	Book.Status = istatus
	h.Db.Save(&Book)
	c.Redirect(http.StatusMovedPermanently, "/books")
}
func (h *BooksHandler) DeleteTask(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	h.Db.Delete(&book)
	c.Redirect(http.StatusMovedPermanently, "/books")
}
