package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClientService struct {
	userRepository repository.User
}

func NewClientService(uR repository.User) *ClientService {
	return &ClientService{uR}
}

func (cS *ClientService) VerifyToken(token string) (TokenClaims, error) {
	path := "internal.service.auth.ParseToken"

	var tokenClaims TokenClaims

	t, err := jwt.ParseWithClaims(token, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return tokenClaims.Key, nil
	})
	if err != nil {
		return tokenClaims, fmt.Errorf(path+".ParseWithClaims, error: {%w}", err)
	}

	if !t.Valid {
		return TokenClaims{}, fmt.Errorf(path+".Valid, error: {%w}", err)
	}
	if time.Now().Unix() > tokenClaims.Exp {
		return TokenClaims{}, fmt.Errorf("Token expired")
	}
	return tokenClaims, nil
}

func (cS *ClientService) GetUserByEmail(ctx context.Context, params AuthParams) (model.User, error) {
	path := "internal.service.auth.GetUserByEmail"
	user, err := cS.userRepository.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return model.User{}, fmt.Errorf(path+".GetUserByEmail, error: {%w}", err)
	}
	user.Password = ""
	return user, nil
}
