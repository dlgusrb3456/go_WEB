module My_TodoList/app

go 1.20

require (
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.2
	github.com/unrolled/render v1.6.0
	go_WEB/My_TodoList/model v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go_WEB/My_TodoList/model => ../model
