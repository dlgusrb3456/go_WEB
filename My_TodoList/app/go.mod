module My_TodoList/app

go 1.20

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/stretchr/testify v1.8.2
	github.com/unrolled/render v1.6.0
	github.com/urfave/negroni v1.0.0
	go_WEB/My_TodoList/model v0.0.0-00010101000000-000000000000
	go_WEB/WEB_UUID v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.6.0
)

require (
	cloud.google.com/go/compute/metadata v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go_WEB/WEB_UUID => ../../WEB_UUID

replace go_WEB/My_TodoList/model => ../model
