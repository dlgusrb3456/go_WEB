package app

import (
	"fmt"
	"go_WEB/My_TodoList/model"
	"go_WEB/WEB_UUID"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render = render.New()
var SESSION_KEY string = WEB_UUID.GetUUID()

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte(SESSION_KEY))

type AppHandler struct {
	http.Handler //embedded : http.Handler를 AppHandler가 포함하고 있다. has-a 관계임
	dbHandler    model.DBHandler
}

func getSessionID(r *http.Request) string {
	session, errs := store.Get(r, "session")
	if errs != nil {
		return ""
	}

	val := session.Values["id"]
	if val == nil {
		return ""
	} else {
		return val.(string) //https://stackoverflow.com/questions/27137521/how-to-convert-interface-to-string
		//type assertion이 필요함
	}
}

func (a *AppHandler) redirectToMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:3000/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoList(w http.ResponseWriter, r *http.Request) {
	// list := []*model.Todo{}
	// for _, v := range model.TodoMap {
	// 	list = append(list, v)
	// }
	list := a.dbHandler.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) postTodoList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") //input_box로부터 받아온 내용
	tempTodo := a.dbHandler.AddTodo(name)
	// tempTodo := &model.Todo{ID: count, Name: name, Completed: false, CreatedAt: time.Now()}

	// model.TodoMap[count] = tempTodo
	// count += 1
	rd.JSON(w, http.StatusCreated, tempTodo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) deleteTodoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.dbHandler.DeleteTodo(id)

	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})

	}
	// _, ok := model.TodoMap[id]
	//  if !ok {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// } else {
	// 	delete(model.TodoMap, id)
	// 	rd.JSON(w, http.StatusOK, Success{true})

	// }

}

type Complete struct {
	Success    bool `json:"success"`
	Complition bool `json:"complition"`
}

func (a *AppHandler) completeTodoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
	}
	ok_value := a.dbHandler.CompleteTodo(id)
	if ok_value == 1 { //true to false
		rd.JSON(w, http.StatusOK, Complete{true, false})
	} else if ok_value == 2 { //false to true
		rd.JSON(w, http.StatusOK, Complete{true, true})
	} else { // no id in map
		rd.JSON(w, http.StatusOK, Complete{false, false})
	}

	// _, ok := model.TodoMap[id]
	// if ok {
	// 	if model.TodoMap[id].Completed {
	// 		model.TodoMap[id].Completed = false
	// 		rd.JSON(w, http.StatusOK, Complete{true, false})
	// 	} else {
	// 		model.TodoMap[id].Completed = true
	// 		rd.JSON(w, http.StatusOK, Complete{true, true})
	// 	}

	// } else {
	// 	rd.JSON(w, http.StatusOK, Complete{false, false})
	// }
}

func (a *AppHandler) getInfoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
	}
	v, ok := a.dbHandler.GetInfo(id)
	if ok { //존재하면
		fmt.Println(&v)
		rd.JSON(w, http.StatusOK, &v)
	}
}

func (a *AppHandler) Close() {
	a.dbHandler.CloseDB()
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//if user request url is signin.html, then next()
	if strings.Contains(r.URL.Path, "/signin.html") || strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	// if user already signed in
	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
	// if user non already signed in
	// redirect "signin.html"
}

func NewRouter(filepath string) *AppHandler { //main으로 AppHandler를 넘김
	// model.TodoMap = make(map[int]*model.Todo)
	fmt.Println(ClientID_google)
	fmt.Println(ClientPW_google)
	r := mux.NewRouter()

	//n := negroni.Classic() //기본적인 기능을 많이 넣어줌. 파일서버, 로그, recovery 등등..
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.HandlerFunc(CheckSignin), negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)

	a := &AppHandler{}
	a.Handler = n
	a.dbHandler = model.NewDBHandler(filepath)

	fmt.Println(SESSION_KEY)
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)

	r.HandleFunc("/", a.redirectToMain)
	r.HandleFunc("/TodoList", a.getTodoList).Methods("GET")
	r.HandleFunc("/TodoList", a.postTodoList).Methods("POST")
	r.HandleFunc("/TodoList/{id:[0-9]+}", a.deleteTodoList).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoList).Methods("GET")
	r.HandleFunc("/getInfoList/{id:[0-9]+}", a.getInfoList).Methods("GET")

	return a
}
