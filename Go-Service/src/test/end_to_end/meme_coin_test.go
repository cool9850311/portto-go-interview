package end_to_end

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"Go-Service/src/main/infrastructure/initializer"
	"Go-Service/src/main/infrastructure/router"
	"Go-Service/src/main/application/DTO"
)

func TestMemeCoinCreate(t *testing.T) {
	// Initialize the logger, config, and MongoDB client
	initializer.InitLog()
	initializer.InitConfig(initializer.Log)
	initializer.InitMongoClient()
	// Clean up any existing test data before running the test
	_, err := initializer.DB.Collection("meme_coin").DeleteMany(
		context.Background(),
		bson.M{"name": "TestCoin"},
	)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}

	// Create a new Gin engine
	r := gin.Default()

	// Setup the router with the initialized database and logger
	router.SetupRouter(r, initializer.DB, initializer.Log)

	// Create a test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Define the request payload
	createRequest := DTO.CreateMemeCoinRequestDTO{
		Name:        "TestCoin",
		Description: "A test meme coin",
	}

	// Marshal the request payload to JSON
	jsonValue, _ := json.Marshal(createRequest)

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", ts.URL+"/api/meme-coin/", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Additional checks can be added here, such as verifying the response body
}
