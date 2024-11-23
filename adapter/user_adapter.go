package adapter

import (

	"net/http"
	"github.com/gin-gonic/gin"

	"reachout-auth/service"
	"reachout-auth/model"
)

func GetAllUsers(c *gin.Context) {
	users := service.ReadAllUser()
	c.IndentedJSON(http.StatusOK, users)
}

func RegisterUser(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user := service.SaveUser(&newUser)
	c.IndentedJSON(http.StatusCreated, user)
}

func ReadUser(c *gin.Context) {
	id := c.Param("id")

	a := service.ReadUser(id)
	if a != nil {
		c.IndentedJSON(http.StatusOK, a)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}