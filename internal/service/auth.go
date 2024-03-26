package service

import (
	custom_errros "auth/internal/custom-errors"
	"auth/internal/hashPassword"
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenClaims struct {
	Email string `json:"email"`
	ID    int    `json:"ID"`
	Exp   int64  `json:"exp"`
	Key   []byte `json:"key"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepository   repository.User
	tokenRepository  repository.Token
	SecretKeyAccess  []byte
	SecretKeyRefresh []byte
}

func NewAuthService(uR repository.User, tR repository.Token, secretKeyAccess, secretKeyRefresh []byte) *AuthService {
	return &AuthService{uR, tR, secretKeyAccess, secretKeyRefresh}
}

func (aS *AuthService) CreateUser(ctx context.Context, params AuthParams) (model.User, error) {
	path := "internal.service.auth.CreateUser"
	password, err := hashPassword.HashPassword(params.Password)
	if err != nil {
		return model.User{}, fmt.Errorf(path+".HashPassword, error: {%w}", err)
	}

	user, err := aS.userRepository.CreateUser(ctx, model.User{Email: params.Email, Password: password})
	if err != nil {
		if errors.Is(err, custom_errros.ErrAlreadyExists) {
			return model.User{}, custom_errros.ErrAlreadyExists
		}
		return model.User{}, fmt.Errorf(path+".CreateUser, error: {%w}", err)
	}
	return user, nil
}

func (aS *AuthService) GenerateTokens(ctx context.Context, params AuthParams) (Tokens, model.User, error) {
	path := "internal.service.auth.GenerateTokens"

	user, err := aS.userRepository.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+"GetUserByEmail, error: {%w}", err)
	}
	ok := hashPassword.CheckPasswordHash(params.Password, user.Password)
	if !ok {
		return Tokens{}, model.User{}, fmt.Errorf(path+".CheckPasswordHash, error: {%w}", custom_errros.ErrWrongCredetianls)
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
		return Tokens{}, model.User{}, fmt.Errorf(path+".tokenA.SignedString, error: {%w}", err)
	}

	claims.Exp = time.Now().Add(time.Hour * 30).Unix()
	claims.Key = aS.SecretKeyRefresh
	tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := tokenR.SignedString(claims.Key)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+".tokenR.SignedString, error: {%w}", err)
	}
	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, user, nil
}

func (aS *AuthService) SaveToken(ctx context.Context, token model.Token) (model.Token, error) {
	path := "internal.service.auth.SaveToken"

	t, err := aS.tokenRepository.SaveToken(ctx, token)
	if err != nil {
		return model.Token{}, fmt.Errorf(path+".SaveToken, error: {%w}", err)
	}
	return t, nil
}
