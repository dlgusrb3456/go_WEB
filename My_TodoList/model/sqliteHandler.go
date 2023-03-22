package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id,name,completed, createdAt FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tempTodo Todo
		rows.Scan(&tempTodo.ID, &tempTodo.Name, &tempTodo.Completed, &tempTodo.CreatedAt)
		todos = append(todos, &tempTodo)
	}
	return todos
}
func (s *sqliteHandler) AddTodo(name string) *Todo {
	statement, err := s.db.Prepare("INSERT INTO todos (name,completed,createdAt) VALUES (?,?,datetime('now'))")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(name, false)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()

	todo := &Todo{int(id), name, false, time.Now()}
	return todo
}
func (s *sqliteHandler) DeleteTodo(id int) bool {
	return false
}
func (s *sqliteHandler) CompleteTodo(id int) int {
	return 1
}
func (s *sqliteHandler) GetInfo(id int) (*Todo, bool) {
	return nil, false
}

func (s *sqliteHandler) CloseDB() {
	s.db.Close()
}

func newSqliteHandler() DBHandler {
	fmt.Println("it's work 4")

	database, err := sql.Open("sqlite3", "./test.db")
	fmt.Println("it's work 5")

	if err != nil {

		panic("what")
	}
	fmt.Println("it's work 7")
	statement, errs := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	if errs != nil {
		panic(errs)
	}
	fmt.Println("it's work 8")
	statement.Exec()
	fmt.Println("it's work 9")
	return &sqliteHandler{db: database}
}
