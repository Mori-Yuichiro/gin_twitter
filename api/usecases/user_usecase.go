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
	GetUserByUserId(userId uint) (models.UserResponse, error)
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
		Name:     user.Name,
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

func (uu *userUsecase) GetUserByUserId(userId uint) (models.UserResponse, error) {
	user := models.User{}
	if err := uu.ur.GetUserByUserId(&user, userId); err != nil {
		return models.UserResponse{}, err
	}

	tweets := []models.TweetResponse{}
	if len(user.Tweets) > 0 {
		for _, v := range user.Tweets {
			tweetUser := models.UserResponse{
				ID:           v.User.ID,
				Name:         v.User.Name,
				Email:        v.User.Email,
				Password:     v.User.Password,
				Avator:       v.User.Avator,
				DisplayName:  v.User.DisplayName,
				ProfileImage: v.User.ProfileImage,
				Bio:          v.User.Bio,
				Location:     v.User.Location,
				Website:      v.User.Website,
				CreatedAt:    v.User.CreatedAt,
				UpdatedAt:    v.User.UpdatedAt,
			}

			retweets := []models.RetweetResponse{}
			if len(v.Retweets) > 0 {
				for _, ret := range v.Retweets {
					retweet := models.RetweetResponse{
						ID:        ret.ID,
						UserId:    ret.UserId,
						TweetId:   ret.TweetId,
						CreatedAt: ret.CreatedAt,
						UpdatedAt: ret.UpdatedAt,
					}
					retweets = append(retweets, retweet)
				}
			}

			favorites := []models.FavoriteResponse{}
			if len(v.Favorites) > 0 {
				for _, fav := range v.Favorites {
					favorite := models.FavoriteResponse{
						ID:        fav.ID,
						UserId:    fav.UserId,
						TweetId:   fav.TweetId,
						CreatedAt: fav.CreatedAt,
						UpdatedAt: fav.UpdatedAt,
					}
					favorites = append(favorites, favorite)
				}
			}

			tweet := models.TweetResponse{
				ID:        v.ID,
				Content:   v.Content,
				UserId:    v.UserId,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				User:      tweetUser,
				Retweets:  retweets,
				Favorites: favorites,
			}

			tweets = append(tweets, tweet)
		}
	}

	resUser := models.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Avator:       user.Avator,
		DisplayName:  user.DisplayName,
		ProfileImage: user.ProfileImage,
		Bio:          user.Bio,
		Location:     user.Location,
		Website:      user.Website,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Tweets:       tweets,
	}

	return resUser, nil
}
