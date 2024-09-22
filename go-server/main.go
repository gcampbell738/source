package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Server running on port 8080...\n")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil{
        fmt.Fprintf(w,"Parsing form error: %v", err)
        return
    }

    fmt.Fprintf(w,"Form Posted")
    name := r.FormValue("Name")
    email := r.FormValue("Email")

    fmt.Fprintf(w,"Name: %v\nEmail:%v",name,email)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "Invalid Path",http.StatusNotFound)
        return
    }

    if r.Method != http.MethodGet {
        http.Error(w,"Not a get method",http.StatusNotFound)
    }
    
    fmt.Fprintf(w, "Hello, World!")

}