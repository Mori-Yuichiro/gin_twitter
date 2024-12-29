package router

import (
	"gin-twitter/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controllers.IUserController) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	r.POST("/signup", uc.SignUp)
	r.POST("/login", uc.LogIn)

	return r
}
