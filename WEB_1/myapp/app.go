package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
	main에서 작성한 코드 다른 폴더에서 패키지화. (main에서는 얘를 import해서 사용함)
	=> 이렇게 코드를 분리해야 testing이 간편해짐
*/

type User struct { // Json으로 받을 구조체
	FirstName string    `json:"first_name"` // json에서의 convention과 Go에서의 convention을 맞추기 위해 json에서는 해당 property를 저렇게 할 것이다. 라고 명시해줌. annotation
	LastName  string    `json:"last_name"`  // 이렇게 하면 Decode와 marshal 과정에서 해당 값으로 key값을 변경해줌
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct {
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world") //Fprint : w에 해당하는 곳에 적어라
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Request에 있을 Json을 받아서 읽어보자. (User 구조체 형식으로 진행할 것임)
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) //r.Body 는 Reader임, io.ReadCloser
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)                      //어떤 interface를 받아서 Json 형태로 인코딩 해주는 marshal
	w.Header().Add("content-type", "application/json") //Response Header에 이 data 유형이 json이라는 것을 알림. => data가 예쁘게 나옴
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) { // http://localhost:3000/bar2?name=LEE
	name := r.URL.Query().Get("name") //URL에서 쿼리문 읽어들이고 name에 해당하는 값 가져오기
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() // 동적 라우팅이 가능하게 해줌. 아래의 예시들은 모두 정적 라우팅임

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar") //Fprint : w에 해당하는 곳에 적어라
	})

	mux.HandleFunc("/bar2", barHandler) //함수로 빼서 넣어줘도 동일하게 동작함
	mux.Handle("/foo", &fooHandler{})
	return mux
}
