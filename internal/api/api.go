package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchStockPrice() {
	apiKey := os.Getenv("TWELVE_DATA_API_KEY")
	symbol := "AAPL"   // Example stock symbol
	interval := "1day" // Example interval (e.g., 1min, 5min, 1h, 1day)

	// Construct the API URL
	// Example using the Time Series endpoint. Refer to Twelve Data docs for others.
	apiURL := fmt.Sprintf("https://api.twelvedata.com/time_series?symbol=%s&interval=%s&apikey=%s",
		symbol, interval, apiKey)

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error: Received non-200 status code %d. Response: %s\n", resp.StatusCode, string(bodyBytes))
		return
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal the JSON response into our struct
	var timeSeries TimeSeriesData
	err = json.Unmarshal(bodyBytes, &timeSeries)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		// Optionally print the raw body to see what was received
		// fmt.Println("Raw response body:", string(bodyBytes))
		return
	}

	// Check the status from the API response
	if timeSeries.Status != "ok" {
		fmt.Printf("API returned status: %s\n", timeSeries.Status)
		// You might want to check for specific error messages from the API here
		return
	}

	// Print some data (e.g., the first data point if available)
	fmt.Printf("Successfully fetched data for: %s\n", timeSeries.Meta.Symbol)
	if len(timeSeries.Values) > 0 {
		firstDataPoint := timeSeries.Values[0]
		fmt.Printf("First data point (%s):\n", firstDataPoint.Datetime) // Print the date string directly
		fmt.Printf("  Open: %s\n", firstDataPoint.Open)
		fmt.Printf("  High: %s\n", firstDataPoint.High)
		fmt.Printf("  Low: %s\n", firstDataPoint.Low)
		fmt.Printf("  Close: %s\n", firstDataPoint.Close)
		fmt.Printf("  Volume: %s\n", firstDataPoint.Volume)
	} else {
		fmt.Println("No time series values returned.")
	}
}
