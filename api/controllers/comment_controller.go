package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"gin-twitter/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ICommentController interface {
	CreateComment(c *gin.Context)
}

type commentController struct {
	cu usecases.ICommentUsecase
}

func NewCommentController(cu usecases.ICommentUsecase) ICommentController {
	return &commentController{cu}
}

func (cc *commentController) CreateComment(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	comment := models.Comment{}
	if err := c.Bind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	comment.UserId = uint(userId.(float64))
	if err := cc.cu.CreateComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
