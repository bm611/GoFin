package components

import (
	"github.com/bm611/gofin/internal/api"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// StockViewModel is the main model for the stock view application
type StockViewModel struct {
	table         *StockTable
	stocks        []api.StockQuote
	width         int
	height        int
	ready         bool
	loading       bool
	showHelp      bool
	footerMessage string
}

// NewStockViewModel creates a new instance of the stock view model
func NewStockViewModel() StockViewModel {
	return StockViewModel{
		table:         NewStockTable(),
		stocks:        []api.StockQuote{},
		ready:         false,
		loading:       false,
		showHelp:      true,
		footerMessage: "Loading...",
	}
}

// SetStocks updates the stock data
func (m *StockViewModel) SetStocks(stocks []api.StockQuote) {
	m.stocks = stocks
	m.table.SetStocks(stocks)
	m.loading = false

	if len(stocks) > 0 {
		m.footerMessage = "Stock data loaded"
	} else {
		m.footerMessage = "No stock data available"
	}
}

// Init initializes the model
func (m StockViewModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and user input
func (m StockViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.table.MoveUp()
		case "down", "j":
			m.table.MoveDown()
		case "h":
			m.showHelp = !m.showHelp
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if !m.ready {
			m.ready = true
		}

		// Adjust table column widths to fit screen
		m.table.ResizeColumns(m.width)
	}

	return m, cmd
}

// View renders the UI
func (m StockViewModel) View() string {
	if !m.ready {
		return "Initializing..."
	}

	// Create a container style for the entire view
	containerStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height)

	// Header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#0047AB")).
		Width(m.width).
		Align(lipgloss.Center).
		Render("GoFin Stock Viewer")

	// Main content area (table)
	tableContent := m.table.Render()

	// Footer with help text
	footerStyle := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Left)

	helpText := "↑/↓: Navigate • h: Toggle Help • q: Quit"

	if m.loading {
		helpText = "Loading stock data..."
	} else if !m.showHelp {
		helpText = m.footerMessage
	}

	footer := footerStyle.Render(helpText)

	// Combine all elements
	return containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			tableContent,
			footer,
		),
	)
}

// Message types for TEA
// type stocksMsg []api.StockQuote // No longer needed if refresh is removed and initial load is direct
