package todo

import "errors"

type Service interface {
	GetTodos() ([]Todo, error)
	GetTodosByUserID(userID int) ([]Todo, error)
	GetTodoByID(userID int, input GetTodoInput) (Todo, error)
	DeleteTodo(userID int, input GetTodoInput) error
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
func (s *service) GetTodoByID(userID int, input GetTodoInput) (Todo, error) {

	todo, err := s.repository.GetByID(input.ID)
	if todo.UserID != userID {
		return todo, errors.New("Not Found")
	}
	if err != nil {
		return todo, err
	}
	return todo, nil
}
func (s *service) DeleteTodo(inputID int, input GetTodoInput) error {
	todo, err := s.repository.GetByID(input.ID)
	if err != nil {
		return err
	}
	if todo.UserID != inputID {
		return errors.New("Failed to delete, UserID does not match")
	}
	err = s.repository.Delete(input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	todo := Todo{}
	todo.Name = input.Name
	todo.UserID = input.User.ID
	todo.Complete = input.Complete

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
