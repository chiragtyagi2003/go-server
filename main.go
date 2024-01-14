package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// check the path
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// by default the method of request is Get
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// print to the screen
	fmt.Fprintf(w, "Hello world")

}

func formHanlder(w http.ResponseWriter, r *http.Request) {

	// parse the form data
	err := r.ParseForm()
	if err != nil {
		// log the error
		log.Fatal(err)
		fmt.Fprintf(w, "ParseForm() err : %v", err)
	}

	// print to the screen
	fmt.Fprintf(w, "POST request successful")
	fmt.Fprintf(w, "Hello %s", r.FormValue("name"))

}

func main() {

	// fileserver
	// we are telling golang to look for index.html in static dir
	fileserver := http.FileServer(http.Dir("./static"))

	// making routes
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHanlder)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000 \n")
	// this creates the server
	if err := http.ListenAndServe(":8080", nil); err != nil {

		// log the error
		log.Fatal(err)
	}

}
