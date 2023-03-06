package main

import (
	"go_WEB/WEB_3_RESTful/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
