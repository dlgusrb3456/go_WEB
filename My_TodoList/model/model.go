package model

import "time"

type Todo struct {
	ID        int       `json:"id`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at`
}

var TodoMap map[int]*Todo

func init() {
	TodoMap = make(map[int]*Todo)
}

func GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range TodoMap {
		list = append(list, v)
	}
	return list
}

var count int = 0

func AddTodo(name string) *Todo {
	tempTodo := &Todo{Name: name, ID: count, Completed: false}
	TodoMap[count] = tempTodo
	count += 1
	return tempTodo
}

func DeleteTodo(id int) bool {
	_, ok := TodoMap[id]
	if ok {
		delete(TodoMap, id)
		return true
	} else {
		return false
	}
}

func CompleteTodo(id int) int { // 1 : true to false 2: false to true 3: non-exist
	v, ok := TodoMap[id]
	if ok {
		complition := v.Completed
		if complition == true {
			v.Completed = false
			return 1
		} else {
			v.Completed = true
			return 2
		}
	} else {
		return 3
	}

	return 1
}

func GetInfo(id int) (*Todo, bool) {
	v, ok := TodoMap[id]
	if ok {
		return v, true
	} else {
		return nil, false
	}
}
