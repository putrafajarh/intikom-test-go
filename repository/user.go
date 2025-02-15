package repository

import (
	"errors"
	"intikom-test-go/database"
	"intikom-test-go/model"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")
var ErrEmailExists = errors.New("email already exists")

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db := database.GetDB()
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindById(id uint) (model.User, error) {
	var user model.User
	result := r.db.Find(&user, id)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	emailExists, _ := r.FindByEmail(user.Email)

	if emailExists.ID != 0 {
		return model.User{}, ErrEmailExists
	}
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
