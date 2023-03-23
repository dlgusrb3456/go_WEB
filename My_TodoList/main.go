package main

import (
	"fmt"
	"go_WEB/My_TodoList/app"
	"log"
	"net/http"
)

func main() {
	m := app.NewRouter("./test.db")
	defer m.Close()

	log.Println("Todo list start")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}
}
