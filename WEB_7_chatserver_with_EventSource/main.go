package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Println("postMessageHandler ", msg, name)
	sendMessage(name, msg)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	sendMessage("", fmt.Sprintf("add user: %s", username))
}

func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sendMessage("", fmt.Sprintf("left user: %s", username))
}

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var msgCh chan Message

func sendMessage(name, msg string) {
	// send message to all client
	msgCh <- Message{name, msg}
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func main() {
	//Eventsource란 server-sent events에 대한 웹 콘텐츠 인터페이스이다.
	// => 서버가 보내는 정보만 여러 클라이언트가 받음. (서버 -> 클라 의 단방향임) (push 알림, event 알림 등에 사용됨)
	msgCh = make(chan Message)

	es := eventsource.New(nil, nil)
	defer es.Close()

	go processMsgCh(es)

	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es)
	mux.Post("/users", addUserHandler)
	mux.Delete("/users", leftUserHandler)

	n := negroni.Classic() //기본적인 기능을 많이 넣어줌. 파일서버, 로그, recovery 등등..
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)

}
