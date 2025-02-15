package service

import (
	"errors"
	"intikom-test-go/model"
	"intikom-test-go/repository"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskServiceType interface {
	FindAll(userId uint) ([]model.Task, error)
	Create(userId uint, request model.CreateTaskRequest) (model.Task, error)
	FindByUserTaskId(userId uint, taskId uint) (model.Task, error)
	FindById(taskId uint) (model.Task, error)
	Update(userId uint, taskId uint, request model.UpdateTaskRequest) (model.Task, error)
	Delete(userId uint, taskId uint) (model.Task, error)
}

type TaskService struct {
	UserRepository repository.UserRepository
	TaskRepository repository.TaskRepository
}

func NewTaskService() TaskServiceType {
	return &TaskService{
		UserRepository: repository.NewUserRepository(),
		TaskRepository: repository.NewTaskRepository(),
	}
}

func (s *TaskService) FindAll(userId uint) ([]model.Task, error) {
	tasks, err := s.TaskRepository.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) Create(userId uint, request model.CreateTaskRequest) (model.Task, error) {
	task := model.Task{
		UserID:      userId,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}

	createdTask, err := s.TaskRepository.Create(task)
	if err != nil {
		return model.Task{}, err
	}

	return createdTask, nil
}

func (s *TaskService) FindByUserTaskId(userId uint, taskId uint) (model.Task, error) {
	task, err := s.TaskRepository.FindByUserTaskId(userId, taskId)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (s *TaskService) FindById(taskId uint) (model.Task, error) {
	task, err := s.TaskRepository.FindById(taskId)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (s *TaskService) Update(userId uint, taskId uint, request model.UpdateTaskRequest) (model.Task, error) {
	task, err := s.TaskRepository.FindByUserTaskId(userId, taskId)
	if err != nil {
		return model.Task{}, err
	}

	if request.Title != nil {
		task.Title = *request.Title
	}
	if request.Description != nil {
		task.Description = *request.Description
	}
	if request.Status != nil {
		task.Status = *request.Status
	}

	updatedTask, err := s.TaskRepository.Update(task)
	if err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (s *TaskService) Delete(userId uint, taskId uint) (model.Task, error) {
	task, err := s.TaskRepository.FindById(taskId)
	if err != nil {
		return model.Task{}, err
	}

	if task.UserID != userId {
		return model.Task{}, ErrTaskNotFound
	}

	err = s.TaskRepository.Delete(taskId)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}
