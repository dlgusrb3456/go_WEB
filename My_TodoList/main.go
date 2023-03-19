package main

import (
	"go_WEB/My_TodoList/app"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := app.NewRouter()

	n := negroni.Classic() //기본적인 기능을 많이 넣어줌. 파일서버, 로그, recovery 등등..
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
