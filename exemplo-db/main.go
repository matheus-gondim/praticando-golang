package main

import (
	"exemplo-db/controllers"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/user", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", controllers.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", controllers.FindById).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods(http.MethodPut)

	fmt.Println("http:/localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
