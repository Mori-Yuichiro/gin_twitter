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

type ITweetController interface {
	CreateTweet(c *gin.Context)
	GetAllTweet(c *gin.Context)
	DeleteTweet(c *gin.Context)
}

type tweetController struct {
	tu usecases.ITweetUsecase
}

func NewTweetController(tu usecases.ITweetUsecase) ITweetController {
	return &tweetController{tu}
}

func (tc *tweetController) CreateTweet(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	tweet := models.Tweet{}
	if err := c.Bind(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tweet.UserId = uint(userId.(float64))
	if err := tc.tu.CreateTweet(tweet); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (tc *tweetController) GetAllTweet(c *gin.Context) {
	tweetRes, err := tc.tu.GetAllTweet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tweetRes)
}

func (tc *tweetController) DeleteTweet(c *gin.Context) {
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

	err := tc.tu.DeleteTweet(uint(tweetId), uint(userId.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
