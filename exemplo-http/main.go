package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type User struct {
	Name  string
	Email string
}

func index(rw http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
	}

	temp.ExecuteTemplate(rw, "Index", user)
}

func main() {

	port := ":3000"

	http.HandleFunc("/", index)

	printHost(port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func printHost(port string) {
	host := "http://localhost" + port + "\n"
	log.Writer().Write([]byte(host))
}
