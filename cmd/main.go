package main

import (
	"log"
	"net/http"

	"github.com/KaT0819/sample-go-gorilla-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhsot:8000", r))
}
