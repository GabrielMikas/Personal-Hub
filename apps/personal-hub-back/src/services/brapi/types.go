package brapi

type Response struct {
	Results []struct {
		Data struct {
			RegularMarketPrice float64 `json:"regularMarketPrice"`
		} `json:"data"`
	} `json:"results"`
}
