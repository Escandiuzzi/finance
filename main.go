package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"finance/src/models"
)

var categories = []models.Category{
	{Title: "Food"},
	{Title: "Clothes"},
}

func main() {
	//db := db.Connect()

	//initializeTables(db)

	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.GET("/", getHome)
	// v1 := router.Group("/v1")
	// {
	// 	routes.Categories(v1.Group("/categories"), db)
	// }

	router.Run(fmt.Sprintf(":%s", port))
}

func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello, World!")
}

func initializeTables(db *sql.DB) {
	var exists bool
	if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'categories' );").Scan(&exists); err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	if !exists {
		results, err := db.Query("CREATE TABLE categories (id SERIAL PRIMARY KEY, title VARCHAR(100) NOT NULL, created_at TIMESTAMP DEFAULT current_timestamp, updated_at TIMESTAMP DEFAULT current_timestamp);")
		if err != nil {
			fmt.Println("failed to execute query", err)
			return
		}

		fmt.Println("Table created successfully")

		for _, category := range categories {
			queryStmt := `INSERT INTO categories (title) VALUES ($1) RETURNING id;`

			log.Println(category)

			row := db.QueryRow(queryStmt, category.Title)

			var name string

			err := row.Scan(&name)

			if err != nil && err != sql.ErrNoRows {
				log.Println("failed to execute query", queryStmt, &category.Title, err)
				return
			}
		}
		fmt.Println("Mock Categories included in Table", results)
	} else {
		fmt.Println("Table 'categories' already exists ")
	}
}
