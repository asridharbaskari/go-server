package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

// entry point for Go executable
func main() {
	fileServer := http.FileServer(http.Dir("./static")) // creates an empty fileServer checked out at the static directory
	// automatically checks out the index.html file
	http.Handle("/", fileServer)          // starts handling root route
	http.HandleFunc("/form", formHandler) // registers handler function
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server @ port 8080\n")

	// Create the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
