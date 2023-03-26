module go_WEB/My_TodoList

go 1.20

require go_WEB/My_TodoList/app v0.0.0-00010101000000-000000000000

require (
	cloud.google.com/go/compute/metadata v0.2.0 // indirect
	github.com/dlgusrb3456/get_UUID v0.0.0-20230326055157-b26045e76c9c // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/unrolled/render v1.6.0 // indirect
	github.com/urfave/negroni v1.0.0 // indirect
	go_WEB/My_TodoList/model v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/oauth2 v0.6.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace go_WEB/My_TodoList/app => ./app

replace go_WEB/My_TodoList/model => ./model

replace go_WEB/WEB_UUID => ../WEB_UUID
