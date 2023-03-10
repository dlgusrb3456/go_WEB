package decoHandler

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn DecoratorFunc //logger
	h  http.Handler  //mux
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h) //url 요청시 수행되는 함수. logger가 수행됨.
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{fn: fn, h: h}
}
