package app

import (
	"fmt"
	"go_WEB/My_TodoList/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render
var count int

func redirectToMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:3000/todo.html", http.StatusTemporaryRedirect)
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	list := []*model.Todo{}
	for _, v := range model.TodoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func postTodoList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") //input_box로부터 받아온 내용
	tempTodo := &model.Todo{ID: count, Name: name, Completed: false, CreatedAt: time.Now()}

	model.TodoMap[count] = tempTodo
	count += 1
	rd.JSON(w, http.StatusCreated, tempTodo)
}

type Success struct {
	Success bool `json:"success"`
}

func deleteTodoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, ok := model.TodoMap[id]
	if !ok {
		rd.JSON(w, http.StatusOK, Success{false})
	} else {
		delete(model.TodoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})

	}

}

type Complete struct {
	Success    bool `json:"success"`
	Complition bool `json:"complition"`
}

func completeTodoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
	}
	_, ok := model.TodoMap[id]
	if ok {
		if model.TodoMap[id].Completed {
			model.TodoMap[id].Completed = false
			rd.JSON(w, http.StatusOK, Complete{true, false})
		} else {
			model.TodoMap[id].Completed = true
			rd.JSON(w, http.StatusOK, Complete{true, true})
		}

	} else {
		rd.JSON(w, http.StatusOK, Complete{false, false})
	}
}

func getInfoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
	}
	v, ok := model.TodoMap[id]
	if ok {
		fmt.Println(&v)
		rd.JSON(w, http.StatusOK, &v)
	}
}

func NewRouter() http.Handler {
	model.TodoMap = make(map[int]*model.Todo)
	rd = render.New()

	mux := mux.NewRouter()
	mux.HandleFunc("/", redirectToMain)
	mux.HandleFunc("/TodoList", getTodoList).Methods("GET")
	mux.HandleFunc("/TodoList", postTodoList).Methods("POST")
	mux.HandleFunc("/TodoList/{id:[0-9]+}", deleteTodoList).Methods("DELETE")
	mux.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoList).Methods("GET")
	mux.HandleFunc("/getInfoList/{id:[0-9]+}", getInfoList).Methods("GET")

	return mux
}
