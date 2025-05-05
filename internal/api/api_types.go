package api

type TimeSeriesData struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Interval         string `json:"interval"`
		Currency         string `json:"currency"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Exchange         string `json:"exchange"`
		MicCode          string `json:"mic_code"`
		Type             string `json:"type"`
	} `json:"meta"`
	Values []struct {
		Datetime string `json:"datetime"`
		Open     string `json:"open"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Close    string `json:"close"`
		Volume   string `json:"volume"`
	} `json:"values"`
	Status string `json:"status"`
}
