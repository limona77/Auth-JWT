package service

import (
	"auth/internal/hashPassword"
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenClaims struct {
	Email string `json:"email"`
	ID    uint32 `json:"ID"`
	Exp   int64  `json:"exp"`
	Key   []byte `json:"key"`
	jwt.RegisteredClaims
}

type AuthService struct {
	authRepository   repository.User
	SecretKeyAccess  []byte
	SecretKeyRefresh []byte
}

func NewAuthService(aR repository.User, secretKeyAccess, secretKeyRefresh []byte) *AuthService {
	return &AuthService{aR, secretKeyAccess, secretKeyRefresh}
}

func (aS *AuthService) CreateUser(ctx context.Context, params AuthParams) error {
	path := "internal.service.auth.CreateUser"
	password, err := hashPassword.HashPassword(params.Password)
	if err != nil {
		return fmt.Errorf(path+".HashPassword, error: {%w}", err)
	}

	_, err = aS.authRepository.CreateUser(ctx, model.User{Email: params.Email, Password: password})
	if err != nil {
		return fmt.Errorf(path+".CreateUser, error: {%w}", err)
	}
	return nil
}

func (aS *AuthService) GenerateTokens(ctx context.Context, params AuthParams) (Tokens, error) {
	user, err := aS.authRepository.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return Tokens{}, err
	}

	claims := TokenClaims{
		Email: user.Email,
		ID:    user.ID,
		Exp:   time.Now().Add(time.Second * 30).Unix(),
		Key:   aS.SecretKeyAccess,
	}
	tokenA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := tokenA.SignedString(claims.Key)
	if err != nil {
		return Tokens{}, err
	}

	claims.Exp = time.Now().Add(time.Hour * 30).Unix()
	claims.Key = aS.SecretKeyRefresh
	tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := tokenR.SignedString(claims.Key)
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
