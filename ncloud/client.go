package ncloud

// TODO: Documenting

// Client is for using each service API operation
type Client struct {
	Config
	Credentials *Credentials
	Handler     *HttpHandler
}

func NewClient(cfg *Config) *Client {

	handler := NewHttpHandler(cfg.Client, cfg.Logger, cfg.Retryer)

	svc := &Client{
		Credentials: cfg.Credentials,
		Handler: handler,
	}

	return svc
}

func (c *Client) NewRequest(operation *Operation, response interface{}, handler *HttpHandler) *Request {
	method := operation.Method
	path := operation.Version + operation.Path

	switch operation.Credentials {
	case "apigw":
		c.Credentials.setAuthParams(method, path)
	case "oauth":
		c.Credentials.setOauthClient(c.Handler.HttpClient)
	default:
		panic("Create Credentials fail")
	}

	return New(operation, c.Credentials, response, handler)
}