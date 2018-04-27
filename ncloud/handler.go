package ncloud

import (
	"net/http"
)

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

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HandlerFunc func(*http.Request) (*http.Response, error)

func (f HandlerFunc)Do(r *http.Request) (*http.Response, error) {

	return f(r)
}

type Decorator func(HttpClient) HttpClient

func AddHandler(c HttpClient, h ...Decorator) HttpClient {
	addedHandler := c
	for _, function := range h {
		addedHandler = function(addedHandler)
	}
	return addedHandler
}

func (h *HttpHandler)Run(r *Request) (*http.Response, error) {

	c := AddHandler(h.HttpClient,
		WithLogger(h.Logger),
		WithRetryer(h.Retryer),
	)

	resp, err := c.Do(r.HTTPRequest)

	if err != nil {
		return nil, err
	}

	return resp, nil
}