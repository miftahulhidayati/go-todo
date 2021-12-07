package todo

type TodoFormatter struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func FormatTodo(todo Todo) TodoFormatter {
	todoFormatter := TodoFormatter{
		ID:       todo.ID,
		UserID:   todo.UserID,
		Name:     todo.Name,
		Complete: todo.Complete,
	}
	return todoFormatter
}
func FormatTodos(todos []Todo) []TodoFormatter {
	todosFormatter := []TodoFormatter{}
	for _, todo := range todos {
		todoFormatter := FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}
	return todosFormatter
}
