package main

import (
	"Tapi/internal"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	port := internal.GenPort()
	r := mux.NewRouter()

	r.HandleFunc("/users", internal.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", internal.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", internal.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", internal.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", internal.DeleteUserHandler).Methods("DELETE")
	r.Handle("/", r)

	http.ListenAndServe(port, r)
}
