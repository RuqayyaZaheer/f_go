package main

import (
	"fmt"
	"net/http"
)

// add for to see git cmd
func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello world its my first Go web </h1>")

}

func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Its a About page</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server is Starting.....")

	http.ListenAndServe(":3000", nil)
}
