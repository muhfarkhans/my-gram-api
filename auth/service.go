package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var JWT_SECRET_KEY = []byte(viper.GetString("jwt.secret_key"))

func init() {
	if os.Getenv("APP_ENV") == "production" {
		JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
	}
}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claim)
	signedToken, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return token, errors.New("invalid token")
	}

	return token, nil
}
