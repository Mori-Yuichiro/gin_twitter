package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
)

type IRelationshipUsecase interface {
	CreateRelationship(relationship models.Relationship) error
	DeleteRelationship(followerId, followedId uint) error
}

type relationshipUsecase struct {
	rr repositories.IRelationshipRepository
}

func NewRelationshipUsecase(rr repositories.IRelationshipRepository) IRelationshipUsecase {
	return &relationshipUsecase{rr}
}

func (ru *relationshipUsecase) CreateRelationship(relationship models.Relationship) error {
	if err := ru.rr.CreateRelationship(&relationship); err != nil {
		return err
	}
	return nil
}

func (ru *relationshipUsecase) DeleteRelationship(followerId, followedId uint) error {
	if err := ru.rr.DeleteRelationship(followerId, followedId); err != nil {
		return err
	}
	return nil
}
