package v1

import (
	"net/http"

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
	c.JSON(http.StatusOK, books)
}

func (h *BooksHandler) CreateBook(c *gin.Context) {
	book := models.Book{}
	err := c.BindJSON(&book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Db.Create(&book)
	c.JSON(http.StatusOK, &book)
}

func (h *BooksHandler) EditBook(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	c.JSON(http.StatusOK, book)
}

func (h *BooksHandler) UpdateBook(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	err := c.BindJSON(&book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.Db.Save(&book)
	c.JSON(http.StatusOK, &book)
}

func (h *BooksHandler) DeleteBook(c *gin.Context) {
	book := models.Book{}
	id := c.Param("id")
	h.Db.First(&book, id)
	err := h.Db.First(&book, id).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	h.Db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
