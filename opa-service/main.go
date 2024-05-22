package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/cdxvy30/foliage/opa-service/domain"
)

var tmpSecretKey = []byte("this-is-tmp-key")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to OPA service!")
}

func genJWTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody domain.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	claims := domain.Claims{
		UID:  requestBody.UID,
		Role: "bronze",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			Issuer:    "opa-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tmpSecretKey)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func main() {
	// Register handlers
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/gen-jwt", genJWTHandler)

	fmt.Println("Starting OPA server at port 9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
