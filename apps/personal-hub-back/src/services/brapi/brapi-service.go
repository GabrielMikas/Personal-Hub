package brapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// DefaultBaseURL is the brapi API root used by NewClient.
const DefaultBaseURL = "https://brapi.dev/api/v2"

// ErrNoResults is returned when brapi answers without any quote.
var ErrNoResults = errors.New("brapi: no results")

// Client talks to the brapi API. BaseURL and HTTPClient are configurable so
// tests can point it to a fake server.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// NewClient returns a client for the real API, authenticated with the
// BRAPI_KEY environment variable.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		Token:      os.Getenv("BRAPI_KEY"),
		HTTPClient: http.DefaultClient,
	}
}

// GetQuote returns the current market price of the given ticker (e.g. TAEE11).
func (c *Client) GetQuote(symbol string) (float64, error) {
	endpoint := c.BaseURL + "/stocks/quote?symbols=" + url.QueryEscape(symbol)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return 0, fmt.Errorf("brapi: building request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("brapi: requesting quote: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("brapi: unexpected status %d", res.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("brapi: decoding response: %w", err)
	}
	if len(response.Results) == 0 {
		return 0, ErrNoResults
	}
	return response.Results[0].Data.RegularMarketPrice, nil
}
