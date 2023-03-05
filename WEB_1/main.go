package main

import (
	"go_WEB/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
	/*
		HandleFunc : 어떤 주소로의 Request가 들어왔을때 어떻게 Handling 할것인지에 대한 내용이다.
		"/" 경로로 요청이 들어오면 func(w http.ResponseWriter, r *http.Request)에 정의된 함수를 실행할 것이다.
		func(w http.ResponseWriter, r *http.Request)이거는 정해진 형식임
		w: Response를 Write 해주는 인자
		r: 사용자가 요청한 Request 정보를 지닌 인자
	*/

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hello world") //Fprint : w에 해당하는 곳에 적어라
	// })

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hello bar") //Fprint : w에 해당하는 곳에 적어라
	// })

	// http.HandleFunc("/bar2", barHandler) //함수로 빼서 넣어줘도 동일하게 동작함

	/*
		위의 HandleFunc은 핸들러를 func() 형태로 직접 등록할 때 사용함
		아래의 Handle은 핸들러를 인스턴스 형태로 등록
	*/
	//http.Handle("/foo", &fooHandler{}) //func Handle(pattern string, handler Handler)

	// type Handler interface {
	// 	ServeHTTP(ResponseWriter, *Request)
	// } => fooHandler가 인터페이스라는건데... 구조체 아닌가..?
	// 저 인터페이스에 대입 가능하게 method를 구현한 구조체 라고 받아들이면 될듯하다.
	/*
		ListenAndServe 실행시 webserver가 구동이 되고 Request를 기다리는 상태가 됨
		Request가 왔을때 이미 사전에 등록된 handler에 대한 요청의 경우 handler대로 처리함
		없는 핸들러 요청시 / 경로의 화면을 보여줌. 이 설정을 바꿀수 있을듯?
	*/

	//http.ListenAndServe(":3000", nil)

}
