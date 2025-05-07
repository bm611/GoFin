package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bm611/gofin/internal/api"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#0047AB")).
			Padding(0, 1)

	cellStyle = lipgloss.NewStyle().
			Padding(0, 1)

	positiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	negativeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000"))
)

// StockTable represents a table for displaying stock data
type StockTable struct {
	stocks       []api.StockQuote
	headers      []string
	columnWidths []int
	focused      bool
	selectedRow  int
}

// NewStockTable creates a new stock table
func NewStockTable() *StockTable {
	headers := []string{
		"Symbol",
		"Name",
		"Price",
		"Change",
		"Change%",
		"Open",
		"High",
		"Low",
		"Volume",
		"Market",
	}

	// Default column widths
	columnWidths := []int{
		10, // Symbol
		25, // Name
		10, // Price
		10, // Change
		10, // Change%
		10, // Open
		10, // High
		10, // Low
		15, // Volume
		10, // Market
	}

	return &StockTable{
		stocks:       []api.StockQuote{},
		headers:      headers,
		columnWidths: columnWidths,
		focused:      true,
		selectedRow:  0,
	}
}

// SetStocks updates the stocks data
func (t *StockTable) SetStocks(stocks []api.StockQuote) {
	t.stocks = stocks
	// Reset selection when data changes
	if t.selectedRow >= len(stocks) {
		t.selectedRow = 0
	}
}

// Focus sets focus on the table
func (t *StockTable) Focus() {
	t.focused = true
}

// Blur removes focus from the table
func (t *StockTable) Blur() {
	t.focused = false
}

// MoveUp moves the selection up
func (t *StockTable) MoveUp() {
	if t.selectedRow > 0 {
		t.selectedRow--
	}
}

// MoveDown moves the selection down
func (t *StockTable) MoveDown() {
	if t.selectedRow < len(t.stocks)-1 {
		t.selectedRow++
	}
}

// SelectedStock returns the currently selected stock
func (t *StockTable) SelectedStock() *api.StockQuote {
	if len(t.stocks) == 0 {
		return nil
	}
	return &t.stocks[t.selectedRow]
}

// RenderHeader renders the table header
func (t *StockTable) RenderHeader() string {
	var header []string
	for i, h := range t.headers {
		width := t.columnWidths[i]
		header = append(header, headerStyle.Width(width).Render(h))
	}
	return strings.Join(header, "")
}

// RenderRow renders a single row
func (t *StockTable) RenderRow(idx int, stock api.StockQuote) string {
	var row []string
	isSelected := t.focused && idx == t.selectedRow

	// Base style
	style := cellStyle
	if isSelected {
		style = style.Copy().Background(lipgloss.Color("#333333"))
	}

	// Style for Symbol column
	row = append(row, style.Width(t.columnWidths[0]).Render(stock.Symbol))

	// Name column
	name := stock.Name
	if len(name) > t.columnWidths[1]-2 {
		name = name[:t.columnWidths[1]-5] + "..."
	}
	row = append(row, style.Width(t.columnWidths[1]).Render(name))

	// Price column
	row = append(row, style.Width(t.columnWidths[2]).Render(FormatCurrency(stock.Close, stock.Currency)))

	// Change column with color
	changeStyle := style.Copy()
	formattedChange := FormatCurrency(stock.Change, stock.Currency)
	if len(stock.Change) > 0 && stock.Change[0] == '-' {
		changeStyle = changeStyle.Inherit(negativeStyle)
	} else if len(stock.Change) > 0 { // Only add '+' if it's a positive change and not zero/empty
		// Add a plus sign for positive changes if not already present
		if len(formattedChange) > 0 && formattedChange[0] != '+' && formattedChange[0] != '-' {
			// Check if the value is actually positive before adding '+'
			// This assumes FormatCurrency doesn't add '+' for positive numbers by default.
			// If stock.Change can be "0" or "0.00", this logic might need adjustment.
			num, err := strconv.ParseFloat(stock.Change, 64)
			if err == nil && num > 0 {
				formattedChange = "+" + formattedChange
			}
		}
		changeStyle = changeStyle.Inherit(positiveStyle)
	} // else, it's likely zero or empty, no specific style or sign needed beyond default
	row = append(row, changeStyle.Width(t.columnWidths[3]).Render(formattedChange))

	// Change% column with color
	changePercentStyle := style.Copy()
	formattedPercentChange := FormatPercentage(stock.PercentChange)
	if len(stock.PercentChange) > 0 && stock.PercentChange[0] == '-' {
		changePercentStyle = changePercentStyle.Inherit(negativeStyle)
	} else if len(stock.PercentChange) > 0 { // Only add '+' if it's a positive change and not zero/empty
		if len(formattedPercentChange) > 0 && formattedPercentChange[0] != '+' && formattedPercentChange[0] != '-' {
			// Check if the value is actually positive before adding '+'
			// This assumes FormatPercentage doesn't add '+' for positive numbers by default.
			num, err := strconv.ParseFloat(stock.PercentChange, 64)
			if err == nil && num > 0 {
				formattedPercentChange = "+" + formattedPercentChange
			}
		}
		changePercentStyle = changePercentStyle.Inherit(positiveStyle)
	} // else, it's likely zero or empty, no specific style or sign needed beyond default
	row = append(row, changePercentStyle.Width(t.columnWidths[4]).Render(formattedPercentChange))

	// Other columns
	row = append(row, style.Width(t.columnWidths[5]).Render(FormatCurrency(stock.Open, stock.Currency)))
	row = append(row, style.Width(t.columnWidths[6]).Render(FormatCurrency(stock.High, stock.Currency)))
	row = append(row, style.Width(t.columnWidths[7]).Render(FormatCurrency(stock.Low, stock.Currency)))
	row = append(row, style.Width(t.columnWidths[8]).Render(FormatLargeNumber(stock.Volume)))

	// Market status
	marketStatus := "Closed"
	if stock.IsMarketOpen {
		marketStatus = "Open"
	}
	row = append(row, style.Width(t.columnWidths[9]).Render(marketStatus))

	return strings.Join(row, "")
}

// Render renders the entire table
func (t *StockTable) Render() string {
	if len(t.stocks) == 0 {
		return "No stock data available. Please fetch data first."
	}

	var sb strings.Builder

	// Render header
	sb.WriteString(t.RenderHeader())
	sb.WriteString("\n")

	// Render separator
	width := 0
	for _, w := range t.columnWidths {
		width += w
	}
	sb.WriteString(strings.Repeat("─", width))
	sb.WriteString("\n")

	// Render rows
	for i, stock := range t.stocks {
		sb.WriteString(t.RenderRow(i, stock))
		sb.WriteString("\n")
	}

	return sb.String()
}

// ResizeColumns adjusts column widths to fit the available width
func (t *StockTable) ResizeColumns(availWidth int) {
	// Calculate total width
	totalWidth := 0
	for _, w := range t.columnWidths {
		totalWidth += w
	}

	// If we have extra space or need to shrink
	if totalWidth != availWidth {
		// Calculate scaling factor
		scale := float64(availWidth) / float64(totalWidth)

		// Apply scaling to all columns while preserving minimums
		newWidths := make([]int, len(t.columnWidths))
		remainingWidth := availWidth

		for i, width := range t.columnWidths {
			// Ensure minimum width of 5 for each column
			newWidth := max(5, int(float64(width)*scale))
			newWidths[i] = newWidth
			remainingWidth -= newWidth
		}

		// Distribute any remaining width (or deficit) to the name column which is most flexible
		if remainingWidth != 0 {
			nameColumnIdx := 1 // Index of the "Name" column
			newWidths[nameColumnIdx] += remainingWidth
			// Ensure minimum width
			if newWidths[nameColumnIdx] < 5 {
				newWidths[nameColumnIdx] = 5
			}
		}

		t.columnWidths = newWidths
	}
}

// Summary returns a summary of the selected stock
func (t *StockTable) Summary() string {
	if len(t.stocks) == 0 {
		return "No stock selected"
	}

	stock := t.stocks[t.selectedRow]

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s (%s)\n", stock.Name, stock.Symbol))
	sb.WriteString(fmt.Sprintf("Exchange: %s | Currency: %s | Last Updated: %s\\n",
		stock.Exchange, stock.Currency, stock.Datetime))

	sb.WriteString("\\n52-Week Range:\\n")
	sb.WriteString(fmt.Sprintf("Low: %s | High: %s | Range: %s\\n",
		FormatCurrency(stock.FiftyTwoWeek.Low, stock.Currency), 
		FormatCurrency(stock.FiftyTwoWeek.High, stock.Currency), 
		stock.FiftyTwoWeek.Range))

	return sb.String()
}
