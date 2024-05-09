package firestore

import (
	"context"
	"fmt"
	"log"

	"github.com/cdxvy30/foliage/twse-service/domain"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func CreateClient(ctx context.Context) *firestore.Client {
	projectID := "foliage-96964"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
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

func ReadAllStocks(ctx context.Context, client *firestore.Client) error {
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
	}
	return nil
}
