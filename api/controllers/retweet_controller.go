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

type IRetweetController interface {
	CreateRetweet(c *gin.Context)
	DeleteRetweet(c *gin.Context)
}

type retweetController struct {
	ru usecases.IRetweetUsecase
}

func NewRetweetController(ru usecases.IRetweetUsecase) IRetweetController {
	return &retweetController{ru}
}

func (rc *retweetController) CreateRetweet(c *gin.Context) {
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

	retweet := models.Retweet{
		UserId:  uint(userId.(float64)),
		TweetId: uint(tweetId),
	}
	if err := rc.ru.CreateRetweet(retweet); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (rc *retweetController) DeleteRetweet(c *gin.Context) {
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

	if err := rc.ru.DeleteRetweet(uint(userId.(float64)), uint(tweetId)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
