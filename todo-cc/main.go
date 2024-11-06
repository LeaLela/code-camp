package main

import (
	"fmt"
	"learajic/code-camp/1/persistence/sqlite"
	
)


func main() {
	db, err := sqlite.NewSqliteDatabase()
	if err != nil {
		panic(fmt.Sprintf("error while initializing database: %s", err.Error()))
	}
	db.GetDb().Ping()
	if err != nil {
		fmt.Printf("error while pinging db: %s", err.Error())
	}
	fmt.Println("app entrypoint")
}