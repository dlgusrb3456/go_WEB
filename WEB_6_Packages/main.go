package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Lee", Email: "dlgusrb@jjj"}

	// w.Header().Add("Content-type", "application/json") //data의 content-type 추가
	// w.WriteHeader(http.StatusOK)                       //data에 status code 설정
	// data, _ := json.Marshal(user)                      //data json화

	// fmt.Fprint(w, string(data))

	rd.JSON(w, http.StatusOK, user) //위 네줄을 이 한줄로 치환 가능함.

}

func addUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user.CreatedAt = time.Now() //Decoding된 data의 created_at 필드를 현재 시간으로 변경해줌.
	rd.JSON(w, http.StatusOK, user)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	rd.JSON(w, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "Lackm")
	user := User{Name: "Lee", Email: "asdf@asdf.com"}
	rd.HTML(w, http.StatusOK, "body", user) //render 패키지가 template의 확장자는 빼고 등록함. .tmpl 없어도 됨

}

func main() {
	rd = render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl"},
		Layout:     "hello"})

	// 확장자가 tmpl이 아닌 html인 파일도 읽어들여라.
	// rd = render.New(render.Options{
	//	Directory: "template" //폴더도 templates가 아닌 template 폴더에서 읽어라
	// 	Extensions: []string{".html",".tmpl"},
	// })
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserInfoHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic() //기본적인 기능을 많이 넣어줌. 파일서버, 로그, recovery 등등..
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
