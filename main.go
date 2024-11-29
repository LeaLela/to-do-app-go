package main

import (
	"fmt"
	"todo-cc/database"
	"todo-cc/infrastructure"
	"todo-cc/service"
)

func main() {

	db, err := database.NewSqliteDatabase()
	if err != nil {
		panic(fmt.Sprintf("error while initializing database: %s", err.Error()))
	}
	err = db.MigrateDB()
	if err != nil {
		panic(err.Error())
	}

	taskPersistence := infrastructure.NewPersistenceAdapter(db.GetDb())

	todoService := service.NewTodo(taskPersistence)
	restController := infrastructure.NewRestController(todoService)

	restController.Run()
}
