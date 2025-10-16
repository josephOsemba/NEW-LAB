package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	baseURL := "http://localhost:8080/api/v1"

	// Test health endpoint
	fmt.Println("Testing Health Endpoint...")
	if err := testEndpoint(baseURL + "/health"); err != nil {
		log.Fatal("Health check failed:", err)
	}

	// Test universities endpoint
	fmt.Println("\nTesting Universities Endpoint...")
	if err := testEndpoint(baseURL + "/universities"); err != nil {
		log.Fatal("Universities endpoint failed:", err)
	}

	// Test labs endpoint
	fmt.Println("\nTesting Labs Endpoint...")
	if err := testEndpoint(baseURL + "/labs?university_id=1"); err != nil {
		log.Fatal("Labs endpoint failed:", err)
	}

	// Test experiments endpoint
	fmt.Println("\nTesting Experiments Endpoint...")
	if err := testEndpoint(baseURL + "/experiments?lab_id=1"); err != nil {
		log.Fatal("Experiments endpoint failed:", err)
	}

	fmt.Println("\nAll tests passed! Backend is working correctly.")
}

func testEndpoint(url string) error {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	fmt.Printf(" %s - Status: %d\n", url, resp.StatusCode)
	return nil
}
