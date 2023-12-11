package main

import (
	"log"
	C "main/controller"
	"main/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type NewItem struct {
    Item   string `json:"item" binding:"required"`
    Bought bool   `json:"bought"`
}


// CORSMiddleware sets up CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
    db.Init();
    database := db.GetDB();
    if database == nil {
        log.Fatal("Failed to connect to the database")
    }

	shoppingItem := new(C.ShoppingItemController)
	router.POST("/items", shoppingItem.Create)
	router.GET("/items", shoppingItem.All)

	router.Run(":8080")
}