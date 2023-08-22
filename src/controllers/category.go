package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"finance/src/models"
)

var categories = []models.Category{
	{ID: "1", Title: "Food", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "2", Title: "Clothes", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, categories)
}
