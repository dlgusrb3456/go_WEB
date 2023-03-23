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

func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id,name,completed, createdAt FROM todos WHERE sessionID=?", sessionId)
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
func (s *sqliteHandler) AddTodo(name string, sessionId string) *Todo {
	statement, err := s.db.Prepare("INSERT INTO todos (sessionID,name,completed,createdAt) VALUES (?,?,?,datetime('now'))")
	if err != nil {
		panic(err)
	}
	result, err := statement.Exec(sessionId, name, false)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()

	todo := &Todo{int(id), name, false, time.Now()}
	return todo
}
func (s *sqliteHandler) DeleteTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, errs := stmt.Exec(id)
	if errs != nil {
		panic(errs)
	}
	cnt, _ := rst.RowsAffected()
	if cnt == 1 {
		return true
	}
	return false
}
func (s *sqliteHandler) CompleteTodo(id int) int {
	checkTodo, ok := s.GetInfo(id)
	if ok {
		stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?")
		if err != nil {
			panic(err)
		}
		complition := checkTodo.Completed
		if complition == true {
			rst, errs := stmt.Exec(false, id)
			if errs != nil {
				panic(errs)
			}
			cnt, _ := rst.RowsAffected()
			fmt.Println("affect count by completed: ", cnt)
			return 1
		} else {
			rst, errs := stmt.Exec(true, id)
			if errs != nil {
				panic(errs)
			}
			cnt, _ := rst.RowsAffected()
			fmt.Println("affect count by completed: ", cnt)
			return 2
		}

	} else {
		return 3
	}

}
func (s *sqliteHandler) GetInfo(id int) (*Todo, bool) {
	todos := &Todo{}
	stmt, err := s.db.Prepare("SELECT id,name,completed, createdAt FROM todos WHERE id = ?")
	if err != nil {
		panic(err)
	}
	errs := stmt.QueryRow(id).Scan(&todos.ID, &todos.Name, &todos.Completed, &todos.CreatedAt) //SELECT문 prepare후 Scan으로 값 담기
	if errs != nil {
		panic(err)
	}
	return todos, true
}

func (s *sqliteHandler) CloseDB() {
	s.db.Close()
}

func newSqliteHandler(filepath string) DBHandler {
	fmt.Println("it's work 4")

	database, err := sql.Open("sqlite3", filepath)
	fmt.Println("it's work 5")

	if err != nil {

		panic("what")
	}
	fmt.Println("it's work 7")
	statement, errs := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			sessionID STRING,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		);
		CREATE INDEX IF NOT EXISTS sessionIDIndexOnTodos ON todos(
			sessionID ASC
		);`)
	if errs != nil {
		panic(errs)
	}
	statement.Exec()
	return &sqliteHandler{db: database}
}
