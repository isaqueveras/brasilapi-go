package brasilapi

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

// httpClient models data from a request
type httpClient struct {
	// client internal http client
	client *http.Client
	// url address for this request
	url *url.URL
	// name defines a human-readable name for the request
	name string
}

// newHttpClient creates a new httpClient
func newHttpClient(baseUrl string) *httpClient {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return &httpClient{
		url: u,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// withName set the readable name for this request
func (h *httpClient) withName(name string) *httpClient {
	h.name = name
	return h
}

// withTimeout sets a maximum timeout for outgoing request
func (h *httpClient) withTimeout(seconds int) *httpClient {
	h.client.Timeout = time.Duration(seconds) * time.Second
	return h
}
