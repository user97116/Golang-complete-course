package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })
    r.Run() // Starts the server
}
// Expected Output: Access "http://localhost:8080" in a web browser to see a JSON response.
