package port

import "time"

type PersistencePort interface {
	GetTask(id int) (*TaskDTO, error)
	NewTask(title, description string, deadline time.Time, completed bool) error
	GetAllTasks() (tasks []TaskDTO, err error)
	DeleteTask(taskId int) (err error)
	CompleteTask(taskId int) (err error)
}

type TaskDTO struct {
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
	Deleted     bool
}
