package service

import "auth/internal/repository"

type Auth interface {
	CreateUser()
}

type Services struct {
	Auth
}

type ServicesDeps struct {
	Repository *repository.Repositories
}

func NewServices(deps ServicesDeps) *Services {
	return &Services{Auth: NewAuthService(deps.Repository.User)}

}
