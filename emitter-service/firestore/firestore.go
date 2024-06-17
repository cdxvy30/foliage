package firestore

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cdxvy30/foliage/twse-service/domain"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func CreateClient(ctx context.Context) *firestore.Client {
	// projectID := "saas-platform-lab"
	projectID := "foliage-96964"

	opt := option.WithCredentialsFile("C:/Users/Tintin/Downloads/serviceAccountKey.json")

	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalln("firestore.NewClient: %w", err)
	}

	return client
}

func UpsertRealTimeStocks(ctx context.Context, client *firestore.Client, doc domain.StockData) error {
	_, err := client.Collection("rt_stock").Doc("").Set(ctx, map[string]interface{}{}, firestore.MergeAll)

	if err != nil {
		log.Fatalf("Failed to add document to Firestore: %v", err)
		return err
	}

	return nil
}

func ReadAllStocks(ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	var allData []map[string]interface{}

	iter := client.Collection("stock").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		// Decode and return

		allData = append(allData, doc.Data())
	}
	return allData, nil
}

func ReadStocksById(ctx context.Context, client *firestore.Client, uid string) ([]map[string]interface{}, error) {
	var allData []map[string]interface{}

	iter := client.Collection("stock").Where("uid", "==", uid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		// Decode and return
		allData = append(allData, doc.Data())
	}
	return allData, nil
}

type Stock struct {
	AccountId float64   `json:"accountId"`
	Amount    float64   `json:"amount"`
	CodeName  string    `json:"codeName"`
	CreatedAt time.Time `json:"createdAt"`
	Price     float64   `json:"price"`
	Uid       string    `json:"uid"`
}

func AddStock(ctx context.Context, client *firestore.Client, stock Stock) {
	var _, _, err = client.Collection("stock").Add(ctx, stock)
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
}
