package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type IUserController interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	CsrfToken(c *gin.Context)
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

	c.SetSameSite(http.SameSiteNoneMode)
	// SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
	c.SetCookie(
		"token",
		tokenString,
		3600,
		"/",
		os.Getenv("API_DOMAIN"),
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (uc *userController) CsrfToken(c *gin.Context) {
	token := csrf.GetToken(c)
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("_csrf", token, 3600, "/", os.Getenv("API_DOMAIN"), true, true)
	c.JSON(http.StatusOK, gin.H{"csrf_token": token})
}
