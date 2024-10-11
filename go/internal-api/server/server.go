package server

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer() *http.Server {
	router := newRouter()

	addr := "localhost:8088"
	srv := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(router, &http2.Server{}),
	}
	return srv
}