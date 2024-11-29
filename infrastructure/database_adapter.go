package infrastructure

import (
	"database/sql"
	"fmt"
	"time"
	"todo-cc/port"
)

type SqliteAdapter struct {
	dbClient *sql.DB
}

func NewPersistenceAdapter(dbClient *sql.DB) *SqliteAdapter {
	return &SqliteAdapter{
		dbClient: dbClient,
	}
}

func (a *SqliteAdapter) GetTask(id int) (*port.TaskDTO, error) {
	findTaskSqlStatement := `
  SELECT title, description, deadline, completed, deleted FROM task WHERE id = ?;
`
	statement, err := a.dbClient.Prepare(findTaskSqlStatement)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare query: %v", err.Error())
	}
	defer statement.Close()

	var TaskDTO port.TaskDTO
	err = statement.
		QueryRow(id).
		Scan(&TaskDTO.Title, &TaskDTO.Description, &TaskDTO.Deadline, &TaskDTO.Completed, &TaskDTO.Deleted)
	if err != nil {
		return nil, fmt.Errorf("unable to set ID into statement: %v", err.Error())
	}

	return &TaskDTO, nil
}

func (a *SqliteAdapter) NewTask(title, description string, deadline time.Time, completed bool) error {
	createTaskSql := `INSERT INTO task(title, description, deadline, completed) values(?, ?, ?, ?)`

	stmt, err := a.dbClient.Prepare(createTaskSql)
	if err != nil {
		return fmt.Errorf("unable to prepare insert statement: %v", err.Error())
	}

	_, err = stmt.Exec(title, description, deadline, completed)
	if err != nil {
		return fmt.Errorf("unable to execute insert statement: %v", err.Error())
	}
	return nil
}

func (a *SqliteAdapter) GetAllTasks() (tasks []port.TaskDTO, err error) {
	getAllTasksSql := `SELECT title, description, deadline, completed, deleted FROM task;`

	rows, err := a.dbClient.Query(getAllTasksSql)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare get all tasks statement: %v", err.Error())
	}
	defer rows.Close() // Ensure rows are closed after processing

	for rows.Next() {
		var task port.TaskDTO
		if err := rows.Scan(&task.Title, &task.Description, &task.Deadline, &task.Completed, &task.Deleted); err != nil {
			return nil, fmt.Errorf("error while scanning one of the tasks")
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (a *SqliteAdapter) DeleteTask(taskId int) (err error) {
	deleteTaskSql := `UPDATE task SET deleted = true WHERE id = ?`

	stmt, err := a.dbClient.Prepare(deleteTaskSql)
	if err != nil {
		return fmt.Errorf("unable to prepare delete task statement: %v", err.Error())
	}

	_, err = stmt.Exec(taskId)
	if err != nil {
		return fmt.Errorf("unable to execute delete task statement: %v", err.Error())
	}

	return err
}

func (a *SqliteAdapter) CompleteTask(taskId int) (err error) {
	completeTask := `UPDATE task SET completed = true WHERE id = ?`

	stmt, err := a.dbClient.Prepare(completeTask)
	if err != nil {
		return fmt.Errorf("unable to prepare complete task statement: %v", err.Error())
	}

	_, err = stmt.Exec(taskId)
	if err != nil {
		return fmt.Errorf("unable to execute complete task statement: %v", err.Error())
	}

	return err
}
