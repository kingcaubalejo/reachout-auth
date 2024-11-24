package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"reachout-auth/adapter"
	"reachout-auth/service"
	"reachout-auth/repository"
)



func main() {
	fmt.Println("Server Running a port 8080")

	userRepository := repository.InitMemoryDatabase()
	userService := service.NewUserService(userRepository)
	userAdapter := adapter.NewUserAdapter(userService)

	router := gin.Default()
	router.GET("/users", userAdapter.GetAllUsers)
	router.GET("/users/:id", userAdapter.ReadUser)
	router.POST("/user", userAdapter.RegisterUser)
	
	router.Run("localhost:8080")
}