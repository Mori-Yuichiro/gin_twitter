package controllers

import (
	"gin-twitter/models"
	"gin-twitter/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IImageController interface {
	UploadImage(c *gin.Context)
}

type imageController struct {
	iu usecases.IImageUsecase
}

func NewImageController(iu usecases.IImageUsecase) IImageController {
	return &imageController{iu}
}

func (ic *imageController) UploadImage(c *gin.Context) {
	image := models.Image{}
	if err := c.Bind(&image); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	imageRes, err := ic.iu.UploadImage(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, imageRes)
}
