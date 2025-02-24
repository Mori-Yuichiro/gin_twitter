package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
)

type IBookmarkUsecase interface {
	CreateBookmark(bookmark models.Bookmark) error
	DeleteBookmark(userId, tweetId uint) error
}

type bookmarkUsecase struct {
	br repositories.IBookmarkRepository
}

func NewBookmarkUsecase(br repositories.IBookmarkRepository) IBookmarkUsecase {
	return &bookmarkUsecase{br}
}

func (bu *bookmarkUsecase) CreateBookmark(bookmark models.Bookmark) error {
	if err := bu.br.CreateBookmark(&bookmark); err != nil {
		return err
	}
	return nil
}

func (bu *bookmarkUsecase) DeleteBookmark(userId, tweetId uint) error {
	if err := bu.br.DeleteBookmark(userId, tweetId); err != nil {
		return err
	}
	return nil
}
