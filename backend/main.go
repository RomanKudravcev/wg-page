package main

import "github.com/gin-gonic/gin"

func main() {
    // Create a Gin router with default middleware: logger and recovery (crash-free) middleware
    router := gin.Default()

    // Define a route
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })

    // Start the server on port 8080
    router.Run(":8080")
}