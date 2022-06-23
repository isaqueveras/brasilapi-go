package brasilapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
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

// httpRequest is a http request that can be performed by an httpClient
type httpRequest struct {
	// inner internal http request
	inner *http.Request
}

// httpErr models data from a error
type httpErr struct {
	// url the requested url
	url url.URL

	// statusCode the http status code
	// returned by the remote server
	statusCode int

	// cause an internal error that
	// caused the request to fail
	cause error

	// clientName name of the client that
	// attempted to perform the request
	clientName string
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

// errorMessage return error message
func (e *httpErr) Error() string {
	return fmt.Sprintf("Request failed to [%v] '%v' with status code %v: %+v", http.MethodGet, e.url.String(), e.statusCode, e.cause)
}

// unwrap returns the cause of the error
func (e *httpErr) unwrap() error {
	return e.cause
}

// get creates a new GET request to be sent using a client
func (h *httpClient) get(uri string, data interface{}, e error) (err error) {
	url, err := h.url.Parse(path.Join(h.url.Path, uri))
	if err != nil {
		return err
	}

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, url.String(), nil); err != nil {
		return err
	}

	var resp *http.Response
	if resp, err = h.client.Do(req); err != nil {
		return &httpErr{
			url:        *req.URL,
			cause:      err,
			statusCode: -1,
			clientName: h.name,
		}
	}

	defer func() {
		deferErr := resp.Body.Close()
		if deferErr != nil {
			log.Println(err)
		}
	}()

	switch resp.StatusCode {
	case http.StatusNoContent:
	case http.StatusOK:
		switch value := data.(type) {
		case *[]byte:
			if value != nil {
				if *value, err = ioutil.ReadAll(resp.Body); err != nil {
					return
				}
				err = resp.Body.Close()
			} else {
				err = errors.New("Invalid data type, try *[]byte")
			}
		default:
			err = json.NewDecoder(resp.Body).Decode(data)
		}
	default:
		if err = json.NewDecoder(resp.Body).Decode(e); err == nil {
			err = e
		}
	}

	if err != nil {
		return &httpErr{
			url:        *req.URL,
			statusCode: resp.StatusCode,
			cause:      err,
			clientName: h.name,
		}
	}

	return
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

// withContext constructs a copy of the request using a context
func (req *httpRequest) withContext(ctx context.Context) *httpRequest {
	return &httpRequest{inner: req.inner.WithContext(ctx)}
}

// withTimeout returns a shallow copy of the request with a cancellation timeout mechanism
func (req *httpRequest) withTimeout(seconds int64) (*httpRequest, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(req.inner.Context(), time.Duration(seconds)*time.Second)
	return req.withContext(ctx), cancel
}
