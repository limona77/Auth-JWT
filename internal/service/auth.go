package service

import (
	"auth/internal/custom-errors"
	"auth/internal/hashPassword"
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gookit/slog"

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

	userModel := model.User{Email: params.Email, Password: password}
	_, err = aS.userRepository.CreateUser(ctx, userModel)
	if err != nil {
		if errors.Is(err, custom_errors.ErrAlreadyExists) {
			return Tokens{}, model.User{}, custom_errors.ErrAlreadyExists
		}
		return Tokens{}, model.User{}, fmt.Errorf(path+".CreateUser, error: {%w}", err)
	}
	tokens, user, err := aS.GenerateTokens(ctx, params)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			return Tokens{}, model.User{}, custom_errors.ErrUserNotFound
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			return Tokens{}, model.User{}, custom_errors.ErrWrongCredetianls
		}
		return Tokens{}, model.User{}, fmt.Errorf(path+".GenerateTokens, error: {%w}", err)
	}
	tokenModel := model.Token{
		RefreshToken: tokens.RefreshToken,
		UserID:       user.ID,
	}
	_, err = aS.tokenRepository.SaveToken(ctx, tokenModel)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+".SaveToken, error: {%w}", err)
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
		Exp:    time.Now().Add(time.Second * 15).Unix(),
		Key:    aS.SecretKeyAccess,
	}
	tokenA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := tokenA.SignedString(claims.Key)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+".tokenA.SignedString, error: {%w}", err)
	}

	claims.Exp = time.Now().Add(time.Second * 30).Unix()
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
			return Tokens{}, model.User{}, custom_errors.ErrUserNotFound
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			return Tokens{}, model.User{}, custom_errors.ErrWrongCredetianls
		}
		return Tokens{}, model.User{}, fmt.Errorf(path+".GenerateTokens, error: {%w}", err)
	}
	tokenModel := model.Token{
		RefreshToken: tokens.RefreshToken,
		UserID:       user.ID,
	}
	_, err = aS.tokenRepository.SaveToken(ctx, tokenModel)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+".SaveToken, error: {%w}", err)
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
			return Tokens{}, model.User{}, custom_errors.ErrUserNotFound
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			return Tokens{}, model.User{}, custom_errors.ErrWrongCredetianls
		}
		return Tokens{}, model.User{}, fmt.Errorf(path+".GenerateTokens, error: {%w}", err)
	}
	tokenModel := model.Token{
		RefreshToken: tokens.RefreshToken,
		UserID:       user.ID,
	}
	_, err = aS.tokenRepository.SaveToken(ctx, tokenModel)
	if err != nil {
		return Tokens{}, model.User{}, fmt.Errorf(path+".SaveToken, error: {%w}", err)
	}
	return tokens, user, nil
}

func (aS *AuthService) Logout(ctx context.Context, token string) (int, error) {
	path := "internal.service.auth.Logout"
	userID, err := aS.tokenRepository.RemoveToken(ctx, token)
	if err != nil {
		return 0, fmt.Errorf(path+".RemoveToken, error: {%w}", err)
	}

	return userID, nil
}
