package service

import "microservice_go/domain"

type UserService interface {
	GetAll() ([]domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}
