module go_WEB/My_TodoList

go 1.20

require (
	github.com/urfave/negroni v1.0.0
	go_WEB/My_TodoList/app v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/pat v1.0.1 // indirect
)

replace go_WEB/My_TodoList/app => ./app
