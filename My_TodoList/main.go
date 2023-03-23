package main

import (
	"fmt"
	"go_WEB/My_TodoList/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.NewRouter("./test.db")
	defer m.Close()

	n := negroni.Classic() //기본적인 기능을 많이 넣어줌. 파일서버, 로그, recovery 등등..
	n.UseHandler(m)

	log.Println("Todo list start")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}
}
