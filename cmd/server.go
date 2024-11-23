package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"reachout-auth/adapter"
)



func main() {
	fmt.Println("Server Running a port 8080")

	router := gin.Default()
	router.GET("/users", adapter.GetAllUsers)
	router.GET("/users/:id", adapter.ReadUser)
	router.POST("/user", adapter.RegisterUser)
	
	router.Run("localhost:8080")
}