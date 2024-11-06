package sqlite

import (
	"database/sql"
	"fmt"
	"learajic/code-camp/1/persistence"
	_ "github.com/mattn/go-sqlite3"
)	

type Db struct {
	database *sql.DB
}

func NewSqliteDatabase() (*Db, error) {
	db,  err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return nil, &persistence.DbConnectionError{
			Message: fmt.Sprintf("error while connecting to database: %v+", err.Error()),
		}
	}

	dbInstance := &Db{database: db}

	return dbInstance, nil
}

func (sl *Db) GetDb() *sql.DB {
	return sl.database
}