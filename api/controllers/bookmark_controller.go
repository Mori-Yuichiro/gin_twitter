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

type IBookmarkController interface {
	CreateBookmark(c *gin.Context)
	DeleteBookmark(c *gin.Context)
}

type bookmarkController struct {
	bu usecases.IBookmarkUsecase
}

func NewBookmarkController(bu usecases.IBookmarkUsecase) IBookmarkController {
	return &bookmarkController{bu}
}

func (bc *bookmarkController) CreateBookmark(c *gin.Context) {
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

	bookmark := models.Bookmark{
		UserId:  uint(userId.(float64)),
		TweetId: uint(tweetId),
	}
	if err := bc.bu.CreateBookmark(bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (bc *bookmarkController) DeleteBookmark(c *gin.Context) {
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

	if err := bc.bu.DeleteBookmark(uint(userId.(float64)), uint(tweetId)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
