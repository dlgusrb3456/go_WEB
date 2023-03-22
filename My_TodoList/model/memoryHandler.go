package model

func (m *memoryHandler) addTodo(name string) *Todo {
	tempTodo := &Todo{Name: name, ID: m.Count, Completed: false}
	m.TodoMap[m.Count] = tempTodo
	m.Count += 1
	return tempTodo
}

func (m *memoryHandler) deleteTodo(id int) bool {
	_, ok := m.TodoMap[id]
	if ok {
		delete(m.TodoMap, id)
		return true
	} else {
		return false
	}
}

func (m *memoryHandler) completeTodo(id int) int { // 1 : true to false 2: false to true 3: non-exist
	v, ok := m.TodoMap[id]
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

func (m *memoryHandler) getInfo(id int) (*Todo, bool) {
	v, ok := m.TodoMap[id]
	if ok {
		return v, true
	} else {
		return nil, false
	}
}
func newMemoryHandler() dbHandler {
	m := &memoryHandler{}
	m.TodoMap = make(map[int]*Todo)
	m.Count = 0
	return m
}

func (m *memoryHandler) getTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.TodoMap {
		list = append(list, v)
	}
	return list
}
