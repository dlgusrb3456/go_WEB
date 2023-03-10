package main

import (
	"go_WEB/WEB_4_2_Decorator_WebHandler/decoHandler"
	"log"
	"net/http"
	"time"
	"web4/decoHandler"
	"web4/myapp"
)

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h1 := decoHandler.NewDecoHandler(mux, logger)
	h2 := decoHandler.NewDecoHandler(h1, logger2)
	return h2 //h2 수행시 logger2 -> logger1 -> 요청 url 순으로 실행됨
}

// Decorator 해보기
func logger(w http.ResponseWriter, r *http.Request, h http.Handler) { //Decoragorfunc 타입 구조체를 여기서 구현해서 사용. logger수행 => 기존 url 실행
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, r) //인자로 받은 핸들러 호출 (chaning), 이것이 요청들어온 url임
	log.Print("[LOGGER1] Completed time: ", time.Since(start).Milliseconds())

}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) { //Decoragorfunc 타입 구조체를 여기서 구현해서 사용. logger수행 => 기존 url 실행
	start := time.Now()
	log.Print("[LOGGER2] Started")
	h.ServeHTTP(w, r) //인자로 받은 핸들러 호출 (chaning), 이것이 요청들어온 url임
	log.Print("[LOGGER2] Completed time: ", time.Since(start).Milliseconds())

}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World")
// }

func main() {

	http.ListenAndServe(":3000", NewHandler())
}
