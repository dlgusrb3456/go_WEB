package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue(("msg"))
	name := r.FormValue("name")
	log.Println("postMessageHandler ", msg, name)
}
func main() {
	//Eventsource란 server-sent events에 대한 웹 콘텐츠 인터페이스이다.
	// => 서버가 보내는 정보만 여러 클라이언트가 받음. (서버 -> 클라 의 단방향임) (push 알림, event 알림 등에 사용됨)

	mux := pat.New()
	mux.Post("/messages", postMessageHandler)

	n := negroni.Classic() //기본적인 기능을 많이 넣어줌. 파일서버, 로그, recovery 등등..
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)

}
