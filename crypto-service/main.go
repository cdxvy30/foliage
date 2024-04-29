package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"crypto-service/balance"
)

var SANDBOX_API = "https://sandbox-api.coinmarketcap.com"
var PRO_API = "https://pro-api.coinmarketcap.com"
var TEST_API_KEY = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
var MY_API_KEY = "f3bdc643-1c1c-4e19-a717-fdc5436f6c55"

func main() {
	fmt.Println("Hello, World!")

	// client := &http.Client{}
	req, err := http.NewRequest("GET", PRO_API+`/v1/cryptocurrency/listings/latest`, nil)
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}

	// q := url.Values{}
	// q.Add("symbol", "BTC,ETH,USDT")
	// q.Add("start", "1")
	// q.Add("limit", "5000")
	// q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	// req.Header.Set("Accept-Encoding", "deflate, gzip")
	req.Header.Set("X-CMC_PRO_API_KEY", MY_API_KEY)
	// req.URL.RawQuery = q.Encode()

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending request to server")
	// 	log.Print(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(res.Status)
	// resBody, _ := io.ReadAll(res.Body)
	// fmt.Println(string(resBody))

	// Call account balance func
	balance.GetWalletBalance("0x3fe42272e6d8859764e752454755632ae510e0c2")
}
