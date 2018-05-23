package internal

import (
	"net/http"
)

// TODO: Documenting

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type handlerFunc func(*http.Request) (*http.Response, error)

func (f handlerFunc) Do(r *http.Request) (*http.Response, error) {

	return f(r)
}

type Decorator func(httpClient) httpClient

func addHandler(c httpClient, h ...Decorator) httpClient {
	addedHandler := c
	for _, function := range h {
		addedHandler = function(addedHandler)
	}
	return addedHandler
}

type HttpHandler struct {
	HttpClient *http.Client
	*Retryer
	*Logger
}

func NewHttpHandler(c *http.Client, l *Logger ,r *Retryer) *HttpHandler {
	return &HttpHandler{
		HttpClient: c,
		Retryer:    r,
		Logger:     l,
	}
}

func (h *HttpHandler) Do(r *Request) (*http.Response, error) {

	c := addHandler(h.HttpClient,
		WithLogger(h.Logger),
		WithRetryer(h.Retryer),
	)

	resp, err := c.Do(r.HTTPRequest)

	if err != nil {
		return nil, err
	}

	return resp, nil
}