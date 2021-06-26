package routes

import (
	"github.com/KaT0819/sample-go-gorilla-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("UPDATE")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
