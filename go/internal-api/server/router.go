package server

import (
	"net/http"
)

func newRouter() *http.ServeMux {

	router := http.NewServeMux()

	return router
}