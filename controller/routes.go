package controller

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes =  {
	return &Routes{
		Route{
			"Catalog",
			"GET",
			"/v2/catalog",
			GetCatalog,
		}
	}
}
