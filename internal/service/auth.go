package service

import (
	"auth/internal/custom-errors"
	"auth/internal/hashPassword"
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/gookit/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenClaims struct {
	Email  string `json:"email"`
	UserID int    `json:"ID"`
	Exp    int64  `json:"exp"`
	Key    []byte `json:"key"`
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

func (aS *AuthService) Register(ctx context.Context, params AuthParams) (Tokens, model.User, error) {
	path := "internal.service.auth.CreateUser"
	password, err := hashPassword.HashPassword(params.Password)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+".HashPassword, error: {%w}", err)
	}

	user, err := aS.userRepository.CreateUser(ctx, model.User{Email: params.Email, Password: password})
	if err != nil {
		if errors.Is(err, custom_errors.ErrAlreadyExists) {
			return Tokens{}, model.User{}, custom_errors.ErrAlreadyExists
		}
		return Tokens{}, model.User{}, fmt.Errorf(path+".CreateUser, error: {%w}", err)
	}
	tokens, user, err := aS.GenerateTokens(ctx, params)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return Tokens{}, model.User{}, custom_errors.ErrUserNotFound
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return Tokens{}, model.User{}, custom_errors.ErrWrongCredetianls
		}
		slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
		return Tokens{}, model.User{}, err
	}
	return tokens, user, nil
}

func (aS *AuthService) GenerateTokens(ctx context.Context, params AuthParams) (Tokens, model.User, error) {
	path := "internal.service.auth.GenerateTokens"

	user, err := aS.userRepository.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+"GetUserByEmail, error: {%w}", err)
	}

	claims := TokenClaims{
		Email:  user.Email,
		UserID: user.ID,
		Exp:    time.Now().Add(time.Second * 30).Unix(),
		Key:    aS.SecretKeyAccess,
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

func (aS *AuthService) Refresh(ctx context.Context, token string) (Tokens, model.User, error) {
	path := "internal.service.auth.RefreshToken"
	c := &ClientService{}
	tokenClaims, err := c.VerifyToken(token)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".VerifyToken, error: {%w}", err).Error())
		return Tokens{}, model.User{}, err
	}
	_, err = aS.tokenRepository.GetToken(ctx, tokenClaims.UserID)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserUnauthorized) {
			return Tokens{}, model.User{}, fmt.Errorf(path+".RefreshToken, error: {%w}", custom_errors.ErrUserUnauthorized)
		}
		return Tokens{}, model.User{}, fmt.Errorf(path+".RefreshToken, error: {%w}", err)
	}

	authParams := AuthParams{Email: tokenClaims.Email}
	tokens, user, err := aS.GenerateTokens(ctx, authParams)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return Tokens{}, model.User{}, custom_errors.ErrUserNotFound
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return Tokens{}, model.User{}, custom_errors.ErrWrongCredetianls
		}
		slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
		return Tokens{}, model.User{}, err
	}
	return tokens, user, nil
}
func (aS *AuthService) Login(ctx context.Context, params AuthParams) (Tokens, model.User, error) {
	path := "internal.service.auth.Login"
	user, err := aS.userRepository.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+"GetUserByEmail, error: {%w}", err)
	}
	ok := hashPassword.CheckPasswordHash(params.Password, user.Password)
	if !ok {
		return Tokens{}, model.User{}, fmt.Errorf(path+".CheckPasswordHash, error: {%w}", custom_errors.ErrWrongCredetianls)
	}
	tokens, user, err := aS.GenerateTokens(ctx, params)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return Tokens{}, model.User{}, custom_errors.ErrUserNotFound
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return Tokens{}, model.User{}, custom_errors.ErrWrongCredetianls
		}
		slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
		return Tokens{}, model.User{}, err
	}
	return tokens, user, nil
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
//eyJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsIklEIjo1OSwiZXhwIjoxNzEyMTc2MDcyLCJrZXkiOiJjMlZqY21WMFgzSmxabkpsYzJoZmEyVjUifQ.
//mBYIbVwQ8XEDpNOhA3SMDk9AQjlL8q2H-o9hNt16IRU
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
//eyJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsIklEIjo1OSwiZXhwIjoxNzEyMTc2MTA2LCJrZXkiOiJjMlZqY21WMFgzSmxabkpsYzJoZmEyVjUifQ.
//Qh81W6upmGAKK5FUEVVYixsKfjgez4Ym-q_AkZ4kQZs
