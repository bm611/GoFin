package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchStockPrice(symbols []string) []StockQuote {
	var stockQuotes []StockQuote
	apiKey := os.Getenv("TWELVE_DATA_API_KEY")

	for _, symbol := range symbols {
		var stockPrice StockQuote
		apiUrl := fmt.Sprintf("https://api.twelvedata.com/quote?symbol=%s&apikey=%s", symbol, apiKey)

		// make api request
		resp, err := http.Get(apiUrl)
		if err != nil {
			fmt.Printf("Error making request for %s: %v\n", symbol, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", symbol, err)
			continue
		}
		if err := json.Unmarshal(body, &stockPrice); err != nil {
			fmt.Printf("Error unmarshaling data for %s: %v\n", symbol, err)
			continue
		}

		stockQuotes = append(stockQuotes, stockPrice)
	}

	return stockQuotes
}
