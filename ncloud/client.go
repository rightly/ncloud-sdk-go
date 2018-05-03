package ncloud

import (
	"log"
	"os"
)

// TODO: Documenting

// Client is for using each service API operation
type Client struct {
	Config
	Credential *Credentials
	Handler *HttpHandler
}

func NewClient(cfg *Config) *Client {
	//method := operation.HTTPMethod

	handler := NewHttpHandler(cfg.Client, cfg.Logger, cfg.Retryer)

	svc := &Client{
		Credential: cfg.Credentials,
		Handler: handler,
	}

	return svc
}

func (c *Client)NewRequest(operation *Operation, response interface{}, handler *HttpHandler) *Request {
	method := operation.Method
	path := operation.Path

	switch operation.Credential {
	case "apigw":
		c.Credential.BuildAuthParams(method, path)
	case "oauth":
		c.Credential.BuildOauthClient(c.Handler.HttpClient)
	default:
		log.Printf("[%s.newRequest().NewRequest()] create credential fail, operation.Credential is not valid value", operation.Name)
		os.Exit(1)
	}

	return New(operation, c.Credential, response, handler)
}
