package controller

import (
	"github.com/gorilla/mux"
  "github.com/customer-api/server/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			ConsoleLog(HttpBaseAuth(Handler(route.HandlerFunc)))
	}
	return router
}
