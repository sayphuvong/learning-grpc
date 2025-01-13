package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 5188...")
	err := http.ListenAndServe(":5188", nil)
	if err != nil {
		panic(err)
	}
}
