package app

import (
	"github.com/gorilla/pat"
)

func NewRouter() *pat.Router {
	mux := pat.New()

	return mux
}
