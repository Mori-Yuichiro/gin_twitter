package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"gin-twitter/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type IFavoriteController interface {
	CreateFavorite(c *gin.Context)
	DeleteFavorite(c *gin.Context)
}

type favoriteController struct {
	fu usecases.IFavoriteUsecase
}

func NewFavoriteController(fu usecases.IFavoriteUsecase) IFavoriteController {
	return &favoriteController{fu}
}

func (fc *favoriteController) CreateFavorite(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]
	id, ok := c.Params.Get("tweetId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid parameter"})
		return
	}
	tweetId, _ := strconv.Atoi(id)

	favorite := models.Favorite{
		UserId:  uint(userId.(float64)),
		TweetId: uint(tweetId),
	}
	if err := fc.fu.CreateFavorite(favorite); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (fc *favoriteController) DeleteFavorite(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]
	id, ok := c.Params.Get("tweetId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid parameter"})
		return
	}
	tweetId, _ := strconv.Atoi(id)

	if err := fc.fu.DeleteFavorite(uint(userId.(float64)), uint(tweetId)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
