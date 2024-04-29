package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const API_URL = "https://tw.rter.info/capi.php"

func main() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}

	req.Header.Set("Accepts", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Println(res.Status)
	resBody, _ := io.ReadAll(res.Body)
	fmt.Println(string(resBody))
}
