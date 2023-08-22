package routes

import (
	categoryController "finance/src/controllers/category"
	"log"

	"github.com/gin-gonic/gin"

	"database/sql"
)

func Categories(router *gin.RouterGroup, db *sql.DB) {

	_, err := db.Query("SELECT * FROM categories;")

	if err != nil {
		log.Println("failed to execute query", err)
		return
	}

	router.GET("/", categoryController.GetCategories)
}
