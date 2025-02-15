package repository

import (
	"errors"
	"intikom-test-go/database"
	"intikom-test-go/model"

	"gorm.io/gorm"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	FindAllByUserId(userId uint) ([]model.Task, error)
	FindById(id uint) (model.Task, error)
	FindByUserTaskId(userId uint, taskId uint) (model.Task, error)
	Create(task model.Task) (model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() TaskRepository {
	db := database.GetDB()
	return &taskRepository{db: db}
}

func (r *taskRepository) FindAllByUserId(userId uint) ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) FindByUserTaskId(userId uint, id uint) (model.Task, error) {
	var task model.Task
	result := r.db.Where("user_id = ?", userId).Find(&task, id)

	if result.Error != nil {
		return model.Task{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (r *taskRepository) FindById(id uint) (model.Task, error) {
	var task model.Task
	result := r.db.Find(&task, id)

	if result.Error != nil {
		return model.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (r *taskRepository) Create(task model.Task) (model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) Update(task model.Task) (model.Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
