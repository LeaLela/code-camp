package service

import (
	"learajic/code-camp/1/port"
	"time"
)
type Todo struct {
	persistence persistence.Port
}
func NewTodo(todoPersistence persistence.Port) Todo {
	return Todo{
		todoPersistence,
	}
}
func (t Todo) CreateNewTask(title, description string, deadline time.Time, completed bool) {
	err := t.persistence.NewTask(title, description, deadline, completed)
	if err != nil {
		
	}
}