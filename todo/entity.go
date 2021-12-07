package todo

import (
	"time"

	"github.com/miftahulhidayati/go-todo/user"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int
	UserID    int
	Name      string
	Complete  bool
	User      user.User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
