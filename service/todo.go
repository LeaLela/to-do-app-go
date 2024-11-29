package service

import (
	"fmt"
	"time"
	"todo-cc/port"
)

type Todo struct {
	persistence port.PersistencePort
}

func NewTodo(todoPersistence port.PersistencePort) Todo {
	return Todo{
		todoPersistence,
	}
}

func (t Todo) CreateNewTask(title, description string, deadline time.Time, completed bool) error {
	err := t.persistence.NewTask(title, description, deadline, completed)
	if err != nil {
		return fmt.Errorf("error while saving task: %v+", err.Error())
	}

	return nil
}

func (t Todo) GetTask(id int) (*port.TaskDTO, error) {
	task, err := t.persistence.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("error while getting task with id(%d): %v+", id, err.Error())
	}

	return task, nil
}

func (t Todo) GetAllTasks() ([]port.TaskDTO, error) {
	tasks, err := t.persistence.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("error while retrieving tasks: %v+", err.Error())
	}

	return tasks, nil
}

func (t Todo) DeleteTask(taskId int) error {
	err := t.persistence.DeleteTask(taskId)
	if err != nil {
		return fmt.Errorf("error while deleting task: %v", err.Error())
	}

	return nil
}

func (t Todo) CompleteTask(taskId int) error {
	err := t.persistence.CompleteTask(taskId)
	if err != nil {
		return fmt.Errorf("error while deleting task: %v", err.Error())
	}

	return nil
}
