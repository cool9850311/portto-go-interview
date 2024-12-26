package end_to_end

import (
	"Go-Service/src/main/application/DTO"
	"Go-Service/src/main/infrastructure/initializer"
	"Go-Service/src/main/infrastructure/router"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func setupTestServer() *httptest.Server {
	// Initialize the logger, config, and MongoDB client
	initializer.InitLog()
	initializer.InitConfig(initializer.Log)
	initializer.InitMongoClient()

	// Create a new Gin engine
	r := gin.Default()

	// Setup the router with the initialized database and logger
	router.SetupRouter(r, initializer.DB, initializer.Log)

	// Create a test server
	return httptest.NewServer(r)
}

func cleanupTestData() {
	// Clean up any existing test data before running the test
	_, err := initializer.DB.Collection("meme_coin").DeleteMany(
		context.Background(),
		bson.M{"name": "TestCoin"},
	)
	if err != nil {
		panic("Failed to clean up test data")
	}
}

func createRequest(method, url string, body interface{}) (*http.Request, error) {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func checkResponse(t *testing.T, resp *http.Response, expectedStatus int) {
	if resp.StatusCode != expectedStatus {
		t.Errorf("Expected status code %d, got %d", expectedStatus, resp.StatusCode)
	}
}

func TestMemeCoinCreate(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()
	cleanupTestData()
	createReqDTO := DTO.CreateMemeCoinRequestDTO{
		Name:        "TestCoin",
		Description: "A test meme coin",
	}

	req, err := createRequest("POST", ts.URL+"/api/meme-coin/", createReqDTO)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	checkResponse(t, resp, http.StatusCreated)
}

func TestMemeCoinGetByID(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()
	cleanupTestData()

	// First create a meme coin to get
	createReqDTO := DTO.CreateMemeCoinRequestDTO{
		Name:        "TestCoin",
		Description: "A test meme coin",
	}

	// Create the coin
	createReq, _ := createRequest("POST", ts.URL+"/api/meme-coin/", createReqDTO)
	createResp, err := http.DefaultClient.Do(createReq)
	if err != nil {
		t.Fatalf("Could not create test coin: %v", err)
	}
	defer createResp.Body.Close()

	// Get the ID from the response
	var createResult map[string]string
	json.NewDecoder(createResp.Body).Decode(&createResult)
	coinID := createResult["id"]

	// Test getting the coin
	getReqDTO := DTO.GetMemeCoinRequestDTO{
		ID: coinID,
	}

	req, err := createRequest("GET", ts.URL+"/api/meme-coin/", getReqDTO)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	checkResponse(t, resp, http.StatusOK)

	var response DTO.GetMemeCoinResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	if response.Name != "TestCoin" {
		t.Errorf("Expected name TestCoin, got %s", response.Name)
	}
}

func TestMemeCoinUpdate(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()
	cleanupTestData()

	// First create a meme coin to update
	createReqDTO := DTO.CreateMemeCoinRequestDTO{
		Name:        "TestCoin",
		Description: "A test meme coin",
	}

	createReq, _ := createRequest("POST", ts.URL+"/api/meme-coin/", createReqDTO)
	createResp, err := http.DefaultClient.Do(createReq)
	if err != nil {
		t.Fatalf("Could not create test coin: %v", err)
	}
	defer createResp.Body.Close()

	var createResult map[string]string
	json.NewDecoder(createResp.Body).Decode(&createResult)
	coinID := createResult["id"]

	// Update the coin
	updateReqDTO := DTO.UpdateMemeCoinRequestDTO{
		ID:          coinID,
		Description: "Updated description",
	}

	req, err := createRequest("PUT", ts.URL+"/api/meme-coin/", updateReqDTO)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	checkResponse(t, resp, http.StatusOK)
}

func TestMemeCoinDelete(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()
	cleanupTestData()

	// First create a meme coin to delete
	createReqDTO := DTO.CreateMemeCoinRequestDTO{
		Name:        "TestCoin",
		Description: "A test meme coin",
	}

	createReq, _ := createRequest("POST", ts.URL+"/api/meme-coin/", createReqDTO)
	createResp, err := http.DefaultClient.Do(createReq)
	if err != nil {
		t.Fatalf("Could not create test coin: %v", err)
	}
	defer createResp.Body.Close()

	var createResult map[string]string
	json.NewDecoder(createResp.Body).Decode(&createResult)
	coinID := createResult["id"]

	// Delete the coin
	deleteReqDTO := DTO.DeleteMemeCoinRequestDTO{
		ID: coinID,
	}

	req, err := createRequest("DELETE", ts.URL+"/api/meme-coin/", deleteReqDTO)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	checkResponse(t, resp, http.StatusOK)
}

func TestMemeCoinPoke(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()
	cleanupTestData()

	// First create a meme coin to poke
	createReqDTO := DTO.CreateMemeCoinRequestDTO{
		Name:        "TestCoin",
		Description: "A test meme coin",
	}

	createReq, _ := createRequest("POST", ts.URL+"/api/meme-coin/", createReqDTO)
	createResp, err := http.DefaultClient.Do(createReq)
	if err != nil {
		t.Fatalf("Could not create test coin: %v", err)
	}
	defer createResp.Body.Close()

	var createResult map[string]string
	json.NewDecoder(createResp.Body).Decode(&createResult)
	coinID := createResult["id"]

	// Poke the coin
	pokeReqDTO := DTO.PokeMemeCoinRequestDTO{
		ID: coinID,
	}

	req, err := createRequest("POST", ts.URL+"/api/meme-coin/poke", pokeReqDTO)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	checkResponse(t, resp, http.StatusOK)

	// Verify the popularity score increased
	getReqDTO := DTO.GetMemeCoinRequestDTO{
		ID: coinID,
	}

	getReq, _ := createRequest("GET", ts.URL+"/api/meme-coin/", getReqDTO)
	getResp, err := http.DefaultClient.Do(getReq)
	if err != nil {
		t.Fatalf("Could not get updated coin: %v", err)
	}
	defer getResp.Body.Close()

	var response DTO.GetMemeCoinResponseDTO
	if err := json.NewDecoder(getResp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	if response.PopularityScore != 1 {
		t.Errorf("Expected popularity score 1, got %d", response.PopularityScore)
	}
}
