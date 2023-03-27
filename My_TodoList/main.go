package main

import (
	"fmt"
	"go_WEB/My_TodoList/app"
	"log"
	"net/http"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "?sr06468sr"
	DB_NAME     = "todo_1"
)

func main() {
	//m := app.NewRouter("./test.db")
	//DATABASE_URL := app.GetDBUrl()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println(dbinfo)
	m := app.NewRouter(dbinfo)

	defer m.Close()

	log.Println("Todo list start")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}
}
