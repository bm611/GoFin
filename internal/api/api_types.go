package api

type FiftyTwoWeekStats struct {
	Low               string `json:"low"`
	High              string `json:"high"`
	LowChange         string `json:"low_change"`
	HighChange        string `json:"high_change"`
	LowChangePercent  string `json:"low_change_percent"`
	HighChangePercent string `json:"high_change_percent"`
	Range             string `json:"range"`
}

type StockQuote struct {
	Symbol                string            `json:"symbol"`
	Name                  string            `json:"name"`
	Exchange              string            `json:"exchange"`
	MicCode               string            `json:"mic_code"`
	Currency              string            `json:"currency"`
	Datetime              string            `json:"datetime"`
	Timestamp             int64             `json:"timestamp"`
	Open                  string            `json:"open"`
	High                  string            `json:"high"`
	Low                   string            `json:"low"`
	Close                 string            `json:"close"`
	Volume                string            `json:"volume"`
	PreviousClose         string            `json:"previous_close"`
	Change                string            `json:"change"`
	PercentChange         string            `json:"percent_change"`
	AverageVolume         string            `json:"average_volume"`
	Rolling1dChange       string            `json:"rolling_1d_change"`
	Rolling7dChange       string            `json:"rolling_7d_change"`
	RollingPeriodChange   string            `json:"rolling_period_change"`
	IsMarketOpen          bool              `json:"is_market_open"`
	FiftyTwoWeek          FiftyTwoWeekStats `json:"fifty_two_week"`
	ExtendedChange        string            `json:"extended_change"`
	ExtendedPercentChange string            `json:"extended_percent_change"`
	ExtendedPrice         string            `json:"extended_price"`
	ExtendedTimestamp     int64             `json:"extended_timestamp"`
	LastQuoteAt           int64             `json:"last_quote_at"`
}
