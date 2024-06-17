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

const API_URL = "https://mis.twse.com.tw/stock/api/getOddInfo.jsp"

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			fetchData()
		}
	}
}

func fetchData() {
	ctx := context.Background()
	firestore_client := firestore.CreateClient(ctx)
	defer firestore_client.Close()

	filePath := os.Getenv("STOCK_LIST_PATH")
	if filePath == "" {
		filePath = "./data/stock_list.csv" // default path
	}
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV file:", err)
	}

	var allData domain.StockData

	batchSize := 100
	for i := 1; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}

		var stockCodes []string
		for _, record := range records[i:end] {
			// Assuming the stock code is in the first column
			if len(record) > 1 {
				stockCodes = append(stockCodes, fmt.Sprintf("tse_%s.tw", record[1]))
			}
		}

		totalList := strings.Join(stockCodes, "|")

		httpClient := &http.Client{
			Timeout: 10 * time.Second,
		}
		req, err := http.NewRequest("GET", API_URL, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Set("Accepts", "application/json")
		q := req.URL.Query()
		q.Add("ex_ch", totalList)
		req.URL.RawQuery = q.Encode()

		res, err := httpClient.Do(req)
		if err != nil {
			log.Print(err)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Printf("Failed to get TWSE stock data for batch starting at %d, status code: %d", i, res.StatusCode)
			continue
		}

		var batchData domain.StockData
		if err := json.NewDecoder(res.Body).Decode(&batchData); err != nil {
			log.Printf("Failed to decode JSON for batch starting at %d: %v", i, err)
			continue
		}

		for i := range batchData.MsgArray {
			tlongStr := batchData.MsgArray[i].DataUpdatedTime
			if tlongStr == "" {
				continue
			}
			tlongInt, err := strconv.ParseInt(tlongStr, 10, 64)
			if err != nil {
				log.Fatalf("Failed to convert tlong to int64: %v", err)
			}
			seconds := tlongInt / 1000
			nanoseconds := (tlongInt % 1000) * 1000000
			batchData.MsgArray[i].Time = time.Unix(seconds, nanoseconds)
		}

		allData.MsgArray = append(allData.MsgArray, batchData.MsgArray...)
	}

	docUpdatedTime := time.Now().Format(time.RFC3339)
	for i := range allData.MsgArray {
		allData.MsgArray[i].DocUpdatedTime = docUpdatedTime
	}

	const docID = "current_stock_data"
	_, err = firestore_client.Collection("rt_stock").Doc(docID).Set(ctx, &allData)
	if err != nil {
		log.Fatalf("Failed to update document in Firestore: %v", err)
	}

	log.Println("Stock data updated successfully in Firestore")
}
