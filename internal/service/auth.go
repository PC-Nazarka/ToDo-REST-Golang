package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo-list/internal/repository"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = "SomeKey"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type JWTAuthorizationService struct {
	userRepo repository.User
}

func NewJWTAuthorizationService(userRepo repository.User) *JWTAuthorizationService {
	return &JWTAuthorizationService{userRepo: userRepo}
}

func (s *JWTAuthorizationService) GenerateToken(id int) (string, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id})
	return token.SignedString([]byte(signingKey))
}

func (s *JWTAuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaims")
	}
	return claims.UserId, nil
}
