package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
)

type IImageUsecase interface {
	UploadImage(image models.Image) (models.ImageResponse, error)
}

type imageUsecase struct {
	ir repositories.IImageRepository
}

func NewImageUsecase(ir repositories.IImageRepository) IImageUsecase {
	return &imageUsecase{ir}
}

func (iu *imageUsecase) UploadImage(image models.Image) (models.ImageResponse, error) {
	uploadResult, err := iu.ir.UploadImage(&image)
	if err != nil {
		return models.ImageResponse{}, err
	}
	resImageUrl := models.ImageResponse{
		Data: uploadResult.URL,
	}
	return resImageUrl, nil
}
