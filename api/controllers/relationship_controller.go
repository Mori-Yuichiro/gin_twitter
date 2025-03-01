package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"gin-twitter/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type IRelationshipController interface {
	CreateRelationship(c *gin.Context)
	DeleteRelationship(c *gin.Context)
}

type relationshipController struct {
	ru usecases.IRelationshipUsecase
}

func NewRelationshipController(ru usecases.IRelationshipUsecase) IRelationshipController {
	return &relationshipController{ru}
}

func (rc *relationshipController) CreateRelationship(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	followerId := uint(claims["userId"].(float64))

	id, ok := c.Params.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid parameter"})
		return
	}
	followedId, _ := strconv.Atoi(id)

	db := &gorm.DB{}
	relationship := models.Relationship{
		FollowerId: followerId,
		FollowedId: uint(followedId),
	}
	if err := relationship.BeforeCreate(db); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := rc.ru.CreateRelationship(relationship); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (rc *relationshipController) DeleteRelationship(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	token, _ := utils.ParseToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	followerId := uint(claims["userId"].(float64))

	id, ok := c.Params.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid parameter"})
		return
	}
	followedId, _ := strconv.Atoi(id)

	if err := rc.ru.DeleteRelationship(followerId, uint(followedId)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
