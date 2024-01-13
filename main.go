package main

import (
	"fmt"
	"log"
	"net/http"
)

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
