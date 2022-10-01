package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name   string
	Email  string
	City   string
	Active bool
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	city := r.FormValue("city")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "E-mail = %s\n", email)
	fmt.Fprintf(w, "City = %s\n", city)

	fmt.Fprintf(w, "JSON Marshalled struct:\n")
	u, err := json.Marshal(User{Name: name, Email: email, City: city, Active: true})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(u))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
