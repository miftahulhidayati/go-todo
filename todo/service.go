package todo

import "errors"

type Service interface {
	GetTodos() ([]Todo, error)
	GetTodosByUserID(userID int) ([]Todo, error)
	GetTodoByID(input GetTodoInput) (Todo, error)
	DeleteTodo(input GetTodoInput) error
	CreateTodo(input CreateTodoInput) (Todo, error)
	UpdateTodo(inputID GetTodoInput, inputData CreateTodoInput) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) GetTodos() ([]Todo, error) {
	todos, err := s.repository.GetAll()

	if err != nil {
		return todos, err
	}
	return todos, nil
}
func (s *service) GetTodosByUserID(userID int) ([]Todo, error) {
	if userID != 0 {
		todos, err := s.repository.GetByUserID(userID)
		if err != nil {
			return todos, err
		}
		return todos, nil
	}

	todos, err := s.repository.GetAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}
func (s *service) GetTodoByID(input GetTodoInput) (Todo, error) {
	todo, err := s.repository.GetByID(input.ID)
	if err != nil {
		return todo, err
	}
	return todo, nil
}
func (s *service) DeleteTodo(input GetTodoInput) error {
	err := s.repository.Delete(input.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	todo := Todo{}
	todo.Name = input.Name
	todo.UserID = input.User.ID
	todo.Complete = false

	newTodo, err := s.repository.Create(todo)
	if err != nil {
		return newTodo, err
	}
	return newTodo, nil
}
func (s *service) UpdateTodo(inputID GetTodoInput, inputData CreateTodoInput) (Todo, error) {
	todo, err := s.repository.GetByID(inputID.ID)
	if err != nil {
		return todo, err
	}
	if todo.UserID != inputData.User.ID {
		return todo, errors.New("UserID does not match")
	}
	todo.Name = inputData.Name
	todo.Complete = inputData.Complete

	updateTodo, err := s.repository.Update(todo)
	if err != nil {
		return updateTodo, err
	}
	return updateTodo, nil
}
