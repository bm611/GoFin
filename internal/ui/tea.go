package ui

import (
	"fmt"

	"github.com/bm611/gofin/internal/api"
	"github.com/bm611/gofin/internal/ui/components"
	tea "github.com/charmbracelet/bubbletea"
)

// StartStockViewer launches the stock data viewer
func StartStockViewer(stocks []api.StockQuote) {
	model := components.NewStockViewModel()
	model.SetStocks(stocks)
	
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),       // Use the full screen
		tea.WithMouseCellMotion(), // Enable mouse support
	)
	
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running stock viewer: %v", err)
	}
}

// StartTea launches the main menu (legacy function)
func StartTea() {
	// Launch the stock viewer with some default data
	symbols := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "META"}
	stocks := api.FetchStockPrice(symbols)
	StartStockViewer(stocks)
}
