package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"gin-twitter/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	csrf "github.com/utrack/gin-csrf"
)

type IUserController interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	LogOut(c *gin.Context)
	CsrfToken(c *gin.Context)
	GetUserIdByToken(c *gin.Context)
	GetUserByUserId(c *gin.Context)
	UpdateUser(c *gin.Context)
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

func (uc *userController) LogOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(
		"token",
		"",
		-1,
		"/",
		os.Getenv("API_DOMAIN"),
		true,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *userController) CsrfToken(c *gin.Context) {
	token := csrf.GetToken(c)
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("_csrf", token, 3600, "/", os.Getenv("API_DOMAIN"), true, true)
	c.JSON(http.StatusOK, gin.H{"csrf_token": token})
}

func (uc *userController) GetUserIdByToken(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (uc *userController) GetUserByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userRes, err := uc.uu.GetUserByUserId(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userRes)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid parameter"})
		return
	}
	userId, _ := strconv.Atoi(id)

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := uc.uu.UpdateUser(user, uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
