package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
	"gin-twitter/validators"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user models.User) error
	LogIn(user models.User) (string, error)
}

type userUsecase struct {
	ur repositories.IUserRepository
	uv validators.IUserValidator
}

func NewUserUsecase(ur repositories.IUserRepository, uv validators.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user models.User) error {
	if err := uu.uv.SignUpValidator(user); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	newUser := models.User{
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) LogIn(user models.User) (string, error) {
	if err := uu.uv.LogInValidator(user); err != nil {
		return "", err
	}

	storeUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(storeUser.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": storeUser.ID,
			"exp":    time.Now().Add(time.Hour * 5).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, nil
}
