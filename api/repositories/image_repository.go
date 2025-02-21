package repositories

import (
	"context"
	"gin-twitter/models"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type IImageRepository interface {
	UploadImage(image *models.Image) (*uploader.UploadResult, error)
}

type imageRepository struct {
	cld *cloudinary.Cloudinary
}

func NewImageRepository(cld *cloudinary.Cloudinary) IImageRepository {
	return &imageRepository{cld}
}

func (ir *imageRepository) UploadImage(image *models.Image) (*uploader.UploadResult, error) {
	ctx := context.Background()
	uploadResult, err := ir.cld.Upload.Upload(
		ctx,
		image.Data,
		uploader.UploadParams{PublicID: ""},
	)
	if err != nil {
		return &uploader.UploadResult{}, err
	}

	return uploadResult, nil
}
