package main

import (
	"fmt"
	"time"
	"learajic/code-camp/1/database"
	"learajic/code-camp/1/infrastructure"
	"learajic/code-camp/1/service"
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

	//TODO remove this, test purpose only.
	taskPersistence.NewTask("Jozo Test", "DESCRIPTION TIE", time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC), false)
	_, _ = taskPersistence.GetTask(1)
	restController.Run()
}
