package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Route not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello Raisa!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error %v", err)
	}

	name := r.FormValue("name")
	lastname := r.FormValue("lastname")

	fmt.Fprintf(w, "Your name: %s\n", name)
	fmt.Fprintf(w, "Your lastname: %s\n", lastname)
}

func main() {
	fileServer := http.FileServer(http.Dir("views"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting at 8080")
	http.ListenAndServe(":8080", nil)
}
