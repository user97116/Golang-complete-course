package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Define routes and handlers
    r.GET("/api/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "List of users"})
    })

    r.POST("/api/users", func(c *gin.Context) {
        // Handle user creation
        c.JSON(http.StatusCreated, gin.H{"message": "User created"})
    })

    r.PUT("/api/users/:id", func(c *gin.Context) {
        // Handle user update
        c.JSON(http.StatusOK, gin.H{"message": "User updated"})
    })

    r.DELETE("/api/users/:id", func(c *gin.Context) {
        // Handle user deletion
        c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
    })

    // Start the server
    r.Run(":8080")
}
// Use a tool like curl or Postman to make requests to the API endpoints.
