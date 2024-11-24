package adapter

import (

	"net/http"
	"github.com/gin-gonic/gin"

	"reachout-auth/service"
	"reachout-auth/model"
)

type UserAdapter struct {
	userService service.UserService
}

func NewUserAdapter(userService service.UserService) *UserAdapter {
	return &UserAdapter{userService: userService}
}

func (u *UserAdapter) GetAllUsers(c *gin.Context) {
	users := u.userService.ReadAllUser()
	c.IndentedJSON(http.StatusOK, users)
}

func (u *UserAdapter) RegisterUser(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user := u.userService.CreateUser(&newUser)
	c.IndentedJSON(http.StatusCreated, user)
}

func (u *UserAdapter) ReadUser(c *gin.Context) {
	id := c.Param("id")

	a := u.userService.ReadUser(id)
	if a != nil {
		c.IndentedJSON(http.StatusOK, a)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}