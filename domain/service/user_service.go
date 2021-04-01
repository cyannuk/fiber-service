package service

import (
	"fiber-service/domain/model"
	"fiber-service/interface/repository"
	"fiber-service/interface/service"
)

type userService struct {
	repository.UserRepository
}

func (service userService) GetUser(id int64) (*model.User, error) {
	return service.UserRepository.GetUser(id)
}

func (service userService) DeleteUser(id int64) error {
	return service.UserRepository.DeleteUser(id)
}

func (service userService) GetUsers(offset int64, limit int64) ([]model.User, error) {
	return service.UserRepository.GetUsers(offset, limit)
}

func (service userService) GetUserOrders(offset int64, limit int64) ([]model.UserOrder, error) {
	return service.UserRepository.GetUserOrders(offset, limit)
}

func (service userService) CreateUser(user *model.User) error {
	return service.UserRepository.CreateUser(user)
}

func (service userService) FindUser(email string) (*model.User, error) {
	return service.UserRepository.FindUser(email)
}

func NewUserService(repository repository.UserRepository) service.UserService {
	return userService{repository}
}
