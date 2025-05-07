package components

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatCurrency formats a numerical string as a currency value
func FormatCurrency(value string, currency string) string {
	// Convert string to float
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value // Return original if parsing fails
	}

	// Default currency symbol if not specified
	symbol := "$"
	
	// Common currency symbols
	switch currency {
	case "USD":
		symbol = "$"
	case "EUR":
		symbol = "€"
	case "GBP":
		symbol = "£"
	case "JPY":
		symbol = "¥"
	}
	
	// Format with 2 decimal places for most currencies
	decimals := 2
	if currency == "JPY" {
		decimals = 0
	}
	
	formattedValue := fmt.Sprintf("%s%."+strconv.Itoa(decimals)+"f", symbol, num)
	
	return formattedValue
}

// FormatPercentage formats a percentage string
func FormatPercentage(value string) string {
	// Check if the value already has a % symbol
	if strings.Contains(value, "%") {
		return value
	}
	
	// Convert string to float
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value // Return original if parsing fails
	}
	
	return fmt.Sprintf("%.2f%%", num)
}

// FormatLargeNumber formats large numbers with K, M, B suffixes
func FormatLargeNumber(value string) string {
	// Convert string to float
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value // Return original if parsing fails
	}
	
	absNum := num
	if absNum < 0 {
		absNum = -absNum
	}
	
	var suffix string
	var divisor float64 = 1
	
	switch {
	case absNum >= 1_000_000_000:
		suffix = "B"
		divisor = 1_000_000_000
	case absNum >= 1_000_000:
		suffix = "M"
		divisor = 1_000_000
	case absNum >= 1_000:
		suffix = "K"
		divisor = 1_000
	default:
		// No suffix for small numbers
		return fmt.Sprintf("%.2f", num)
	}
	
	return fmt.Sprintf("%.2f%s", num/divisor, suffix)
}