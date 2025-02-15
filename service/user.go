package service

import (
	"intikom-test-go/model"
	"intikom-test-go/repository"
	"intikom-test-go/utils"
)

type UserServiceType interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(id uint) (model.User, error)
	Create(user model.RegisterRequest) (model.User, error)
	Update(user model.User, request model.UpdateUserRequest) (model.User, error)
	Delete(id uint) error
}

type UserService struct {
	UserRepository repository.UserRepository
}

func (s *UserService) FindAll() ([]model.User, error) {
	return s.UserRepository.FindAll()
}

func (s *UserService) FindByEmail(email string) (model.User, error) {
	return s.UserRepository.FindByEmail(email)
}

func (s *UserService) FindById(id uint) (model.User, error) {
	return s.UserRepository.FindById(id)
}

func (s *UserService) Create(user model.RegisterRequest) (model.User, error) {

	encryptedPassword := utils.GeneratePassword(user.Password)
	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: encryptedPassword,
	}

	createdUser, err := s.UserRepository.Create(userModel)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

func (s *UserService) Update(user model.User, request model.UpdateUserRequest) (model.User, error) {
	if request.Name != nil {
		user.Name = *request.Name
	}
	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Password != nil {
		user.Password = utils.GeneratePassword(*request.Password)
	}

	updatedUser, err := s.UserRepository.Update(user)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (s *UserService) Delete(id uint) (model.User, error) {
	user, err := s.UserRepository.FindById(id)
	if err != nil {
		return model.User{}, err
	}

	err = s.UserRepository.Delete(user.ID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
