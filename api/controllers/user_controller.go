package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
}

type userController struct {
	uu usecases.IUserUsecase
}

func NewUserController(uu usecases.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := uc.uu.SignUp(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Created"})
}

func (uc *userController) LogIn(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// cookie := new(http.Cookie)
	// cookie.Name = "token"
	// cookie.Value = tokenString
	// cookie.Expires = time.Now().Add(12 * time.Hour)
	// cookie.Path = "/"
	// cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	// cookie.HttpOnly = true
	// cookie.SameSite = http.SameSiteNoneMode

	c.SetSameSite(http.SameSiteNoneMode)
	// SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
	c.SetCookie(
		"token",
		tokenString,
		3600,
		"/",
		os.Getenv("API_DOMAIN"),
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
