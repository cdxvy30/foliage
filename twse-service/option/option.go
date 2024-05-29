package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cdxvy30/foliage/twse-service/domain"
	"github.com/cdxvy30/foliage/twse-service/firestore"
)

const API_URL = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp"

func main() {
	ctx := context.Background()
	firestore_client := firestore.CreateClient(ctx)
	defer firestore_client.Close()

	// Open the CSV file
	file, err := os.Open("/Users/cdxvy30/Downloads/stock_list.csv")
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV file:", err)
	}

	// Initialize a large data structure to hold all the data
	var allData domain.StockData

	// Set the batch size
	batchSize := 100 // Adjust based on your URL length restrictions
	for i := 1; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}

		// Collect the stock codes for the current batch
		var stockCodes []string
		for _, record := range records[i:end] {
			if len(record) > 1 { // Assuming the stock code is in the first column
				stockCodes = append(stockCodes, fmt.Sprintf("tse_%s.tw", record[1]))
			}
		}

		// Combine the stock codes into one query parameter
		totalList := strings.Join(stockCodes, "|")

		// Make the HTTP request
		httpClient := &http.Client{Timeout: 10 * time.Second}
		req, err := http.NewRequest("GET", API_URL, nil)
		if err != nil {
			fmt.Println(err)
			continue // Log error and continue with next batch
		}
		req.Header.Set("Accepts", "application/json")
		q := req.URL.Query()
		q.Add("ex_ch", totalList)
		req.URL.RawQuery = q.Encode()

		res, err := httpClient.Do(req)
		if err != nil {
			log.Print(err)
			continue // Log error and continue with next batch
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Printf("Failed to get TWSE stock data for batch starting at %d, status code: %d", i, res.StatusCode)
			continue // Log error and continue with next batch
		}

		var batchData domain.StockData
		if err := json.NewDecoder(res.Body).Decode(&batchData); err != nil {
			log.Printf("Failed to decode JSON for batch starting at %d: %v", i, err)
			continue // Log error and continue with next batch
		}

		for i := range batchData.MsgArray {
			tlongStr := batchData.MsgArray[i].TLONG
			tlongInt, err := strconv.ParseInt(tlongStr, 10, 64)
			if err != nil {
				log.Fatalf("Failed to convert tlong to int64: %v", err)
			}
			seconds := tlongInt / 1000
			nanoseconds := (tlongInt % 1000) * 1000000
			batchData.MsgArray[i].Time = time.Unix(seconds, nanoseconds)
		}

		// Aggregate batch data into allData
		allData.MsgArray = append(allData.MsgArray, batchData.MsgArray...)
	}

	// Write the aggregated data to Firestore
	const docID = "current_stock_data"
	_, err = firestore_client.Collection("rt_stock").Doc(docID).Set(ctx, &allData)
	if err != nil {
		log.Fatalf("Failed to update document in Firestore: %v", err)
	}

	log.Println("Stock data updated successfully in Firestore")
}
