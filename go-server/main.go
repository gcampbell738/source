package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/formHandler", formHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Server running on port 8080...\n")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func formHandler(w http.ResponseWriter, r *http.Request) {
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}