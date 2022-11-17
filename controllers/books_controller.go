package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfirme-sys/books-api/database"
	"github.com/jfirme-sys/books-api/models"
)

func ShowBook(context *gin.Context) {
	id := context.Param("id")

	integerID, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDB()

	var book models.Book
	err = db.First(book, integerID).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}
