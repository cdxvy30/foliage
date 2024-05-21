package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to OPA service!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting OPA server at port 9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
