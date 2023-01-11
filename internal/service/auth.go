package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
	"todo-list/internal/repository"
)

var (
	accessTokenExpires, _  = strconv.Atoi(os.Getenv("ACCESS_EXPIRES"))
	refreshTokenExpires, _ = strconv.Atoi(os.Getenv("REFRESH_EXPIRES"))
	signingKey             = os.Getenv("SECRET_KEY")
	accessTokenTTL         = time.Minute * time.Duration(accessTokenExpires)
	refreshTokenTTL        = time.Minute * time.Duration(refreshTokenExpires)
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

func (s *JWTAuthorizationService) GenerateAccessRefreshTokens(id int) (string, string, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		return "", "", err
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
		},
		user.Id})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
		},
		user.Id})
	accessTokenString, err := accessToken.SignedString([]byte(signingKey))
	refreshTokenString, err := refreshToken.SignedString([]byte(signingKey))
	return accessTokenString, refreshTokenString, err
}

func (s *JWTAuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if !token.Valid {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaims")
	}
	return claims.UserId, nil
}
