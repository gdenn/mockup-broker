package main

import (
	"log"
	"net/http"

	"github.com/customer-api/server/controller"
)

func main(deployment string) {
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
