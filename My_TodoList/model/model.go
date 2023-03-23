package model

import "time"

type Todo struct {
	ID        int       `json:"id`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at`
}

type DBHandler interface {
	GetTodos(sessionId string) []*Todo
	AddTodo(name string, sessionId string) *Todo
	DeleteTodo(id int) bool
	CompleteTodo(id int) int
	GetInfo(id int) (*Todo, bool)
	CloseDB()
}

type memoryHandler struct {
	TodoMap map[int]*Todo
	Count   int
}

func NewDBHandler(filepath string) DBHandler {
	return newSqliteHandler(filepath)
}

// func init() {
// 	//handler = newMemoryHandler()
// 	handler = newSqliteHandler()
// }

// func GetTodos() []*Todo {
// 	return handler.getTodos()
// }

// func AddTodo(name string) *Todo {
// 	return handler.addTodo(name)
// }

// func DeleteTodo(id int) bool {
// 	return handler.deleteTodo(id)
// }

// func CompleteTodo(id int) int {
// 	return handler.completeTodo(id)
// }

// func GetInfo(id int) (*Todo, bool) {
// 	return handler.getInfo(id)
// }

// func Close() {
// 	handler.closeDB()
// }

// -----------------------------------------

// func GetTodos() []*Todo {
// 	list := []*Todo{}
// 	for _, v := range TodoMap {
// 		list = append(list, v)
// 	}
// 	return list
// }

// func AddTodo(name string) *Todo {
// 	tempTodo := &Todo{Name: name, ID: count, Completed: false}
// 	TodoMap[count] = tempTodo
// 	count += 1
// 	return tempTodo
// }

// func DeleteTodo(id int) bool {
// 	_, ok := TodoMap[id]
// 	if ok {
// 		delete(TodoMap, id)
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func CompleteTodo(id int) int { // 1 : true to false 2: false to true 3: non-exist
// 	v, ok := TodoMap[id]
// 	if ok {
// 		complition := v.Completed
// 		if complition == true {
// 			v.Completed = false
// 			return 1
// 		} else {
// 			v.Completed = true
// 			return 2
// 		}
// 	} else {
// 		return 3
// 	}

// 	return 1
// }

// func GetInfo(id int) (*Todo, bool) {
// 	v, ok := TodoMap[id]
// 	if ok {
// 		return v, true
// 	} else {
// 		return nil, false
// 	}
// }
