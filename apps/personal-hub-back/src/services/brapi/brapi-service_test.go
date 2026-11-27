package brapi

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Setenv("BRAPI_KEY", "test-token")

	c := NewClient()
	if c.BaseURL != DefaultBaseURL || c.Token != "test-token" || c.HTTPClient != http.DefaultClient {
		t.Fatalf("unexpected client: %+v", c)
	}
}

func TestGetQuote(t *testing.T) {
	tests := []struct {
		name      string
		status    int
		body      string
		wantPrice float64
		wantErr   string
	}{
		{name: "success", status: http.StatusOK, body: `{"results":[{"data":{"regularMarketPrice":35.42}}]}`, wantPrice: 35.42},
		{name: "unexpected status", status: http.StatusUnauthorized, body: `{}`, wantErr: "unexpected status 401"},
		{name: "invalid json", status: http.StatusOK, body: `{`, wantErr: "decoding response"},
		{name: "no results", status: http.StatusOK, body: `{"results":[]}`, wantErr: ErrNoResults.Error()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/stocks/quote" || r.URL.Query().Get("symbols") != "TAEE11" {
					t.Errorf("unexpected request %s", r.URL)
				}
				if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
					t.Errorf("Authorization = %q", got)
				}
				w.WriteHeader(tt.status)
				w.Write([]byte(tt.body))
			}))
			defer server.Close()

			c := &Client{BaseURL: server.URL, Token: "test-token", HTTPClient: server.Client()}
			price, err := c.GetQuote("TAEE11")

			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("err = %v, want it to contain %q", err, tt.wantErr)
				}
				return
			}
			if err != nil || price != tt.wantPrice {
				t.Fatalf("GetQuote = %v, %v; want %v, nil", price, err, tt.wantPrice)
			}
		})
	}
}

func TestGetQuoteRequestErrors(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		wantErr string
	}{
		{name: "invalid url", baseURL: "http://[::1", wantErr: "building request"},
		{name: "unreachable server", baseURL: "http://127.0.0.1:1", wantErr: "requesting quote"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{BaseURL: tt.baseURL, HTTPClient: http.DefaultClient}
			_, err := c.GetQuote("TAEE11")
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want it to contain %q", err, tt.wantErr)
			}
			if errors.Is(err, ErrNoResults) {
				t.Fatal("request errors should not be ErrNoResults")
			}
		})
	}
}
