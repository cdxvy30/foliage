package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	redis "github.com/redis/go-redis/v9"
)

// listenDocument listens to a single document.
func listenDocument(ctx context.Context, w io.Writer, projectID, collection string) error {
	// projectID := "project-id"
	// Сontext with timeout stops listening to changes.
	ctx = context.Background()
	// defer cancel()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("firestore.NewClient: %w", err)
	}
	defer client.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	it := client.Collection(collection).Doc("current_stock_data").Snapshots(ctx)
	for {
		snap, err := it.Next()
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Snapshots.Next: %w", err)
		}
		if !snap.Exists() {
			fmt.Fprintf(w, "Document no longer exists\n")
			return nil
		}
		// fmt.Fprintf(w, "Received document snapshot: %v\n", snap.Data())
		fmt.Fprintf(w, "Something changed!\n")

		data := snap.Data()
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Failed to marshal data: %v", err)
		}

		// Publish to Redis
		err = redisClient.Publish(ctx, "stockUpdates", jsonData).Err()
		if err != nil {
			log.Fatalf("Failed to publish to Redis: %v", err)
		}
	}
}

func main() {
	ctx := context.Background()

	listenDocument(ctx, os.Stdout, "foliage-96964", "rt_stock")
}
