package main

import (
	"fmt"
	"net/http"

	"github.com/cdxvy30/foliage/opa-service/handler"
)

func main() {
	// Register handlers
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/gen-jwt", handler.GenJWTHandler)

	fmt.Println("Starting OPA server at port 9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
