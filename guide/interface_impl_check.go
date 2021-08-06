package main

import "net/http"

type Handler struct {
	//...
}

// interface的合理性验证
var _ http.Handler = (*Handler)(nil)

func (h Handler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}
