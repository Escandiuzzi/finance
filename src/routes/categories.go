package routes

import (
	"finance/src/controllers"
	"log"

	"github.com/gin-gonic/gin"

	"database/sql"
)

func Categories(router *gin.RouterGroup, db *sql.DB) {

	results, err := db.Query("SELECT * FROM categories;")

	if err != nil {
		log.Println("failed to execute query", err)
		return
	}

	log.Println("query result", results)

	router.GET("/", controllers.GetCategories)
}
