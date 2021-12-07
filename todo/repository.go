package todo

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	GetByUserID(UserID int) ([]Todo, error)
	Create(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(id int) error
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return todos, err
	}
	return todos, nil
}
func (r *repository) GetByID(id int) (Todo, error) {
	var todo Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}
func (r *repository) GetByUserID(UserID int) ([]Todo, error) {
	var todos []Todo
	err := r.db.Where("user_id = ?", UserID).Find(&todos).Error
	if err != nil {
		return todos, err
	}
	return todos, nil
}
func (r *repository) Create(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}
func (r *repository) Update(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}
func (r *repository) Delete(id int) error {
	err := r.db.Delete(&Todo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
