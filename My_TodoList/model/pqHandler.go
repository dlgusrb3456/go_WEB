package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type pqHandler struct {
	db *sql.DB
}

func (p *pqHandler) GetTodos(sessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := p.db.Query("SELECT id,name,completed, createdAt FROM todos WHERE sessionID=$1", sessionId)
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
func (p *pqHandler) AddTodo(name string, sessionId string) *Todo {
	statement, err := p.db.Prepare("INSERT INTO todos (sessionID,name,completed,createdAt) VALUES ($1,$2,$3,now())")
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
func (p *pqHandler) DeleteTodo(id int) bool {
	stmt, err := p.db.Prepare("DELETE FROM todos WHERE id=$1")
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
func (p *pqHandler) CompleteTodo(id int) int {
	checkTodo, ok := p.GetInfo(id)
	if ok {
		stmt, err := p.db.Prepare("UPDATE todos SET completed=$1 WHERE id=$2")
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
func (p *pqHandler) GetInfo(id int) (*Todo, bool) {
	todos := &Todo{}
	stmt, err := p.db.Prepare("SELECT id,name,completed, createdAt FROM todos WHERE id = $1")
	if err != nil {
		panic(err)
	}
	errs := stmt.QueryRow(id).Scan(&todos.ID, &todos.Name, &todos.Completed, &todos.CreatedAt) //SELECT문 prepare후 Scan으로 값 담기
	if errs != nil {
		panic(err)
	}
	return todos, true
}

func (p *pqHandler) CloseDB() {
	p.db.Close()
}

func newPQHandler(dbConn string) DBHandler {
	fmt.Println("it's work 4")
	fmt.Println("dbConn::", dbConn)
	database, err := sql.Open("postgres", dbConn)
	fmt.Println("it's work 5")

	if err != nil {
		panic("what")
	}
	if err = database.Ping(); err != nil {
		fmt.Println("ping Error")
		panic(err)
	}

	fmt.Println("it's work 7")
	statement, errs := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id        SERIAL PRIMARY KEY,
			sessionID VARCHAR(256),
			name      TEXT,
			completed BOOLEAN,
			createdAt TIMESTAMP
		);
	`)
	if errs != nil {
		fmt.Println("prepare error")
		panic(errs)
	}
	_, errs = statement.Exec()
	if errs != nil {
		fmt.Println("exec error")
		panic(errs)
	}
	statement, errs = database.Prepare(
		`CREATE INDEX IF NOT EXISTS sessionIDIndexOnTodos ON todos(
			sessionID ASC
		);
	`)

	if errs != nil {
		panic(errs)
	}
	_, errs = statement.Exec()
	if errs != nil {
		panic(errs)
	}

	return &pqHandler{db: database}
}
