package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cdxvy30/foliage/emitter-service/firestore"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/stocks", getStocks)
	router.GET("/stocks/:id", getStocksById)
	router.POST("/stocks", postStock)

	router.Run("localhost:8000")
}

func getStocks(c *gin.Context) {
	ctx := context.Background()
	client := firestore.CreateClient(ctx)
	defer client.Close()

	jsonData, err := firestore.ReadAllStocks(ctx, client)
	if err != nil {
		log.Fatalf("Failed to read all stocks: %v", err)
	}
	c.IndentedJSON(http.StatusOK, jsonData)
}

func getStocksById(c *gin.Context) {
	ctx := context.Background()
	client := firestore.CreateClient(ctx)
	defer client.Close()

	uid := c.Param("id")
	jsonData, err := firestore.ReadStocksById(ctx, client, uid)
	if err != nil {
		log.Fatalf("Failed to read all stocks: %v", err)
	}
	c.IndentedJSON(http.StatusOK, jsonData)
}

func postStock(c *gin.Context) {
	var stock firestore.Stock
	ctx := context.Background()
	client := firestore.CreateClient(ctx)
	defer client.Close()

	if err := c.BindJSON(&stock); err != nil {
		return
	}
	firestore.AddStock(ctx, client, stock)
	c.IndentedJSON(http.StatusCreated, stock)
}
