package ncloud

//TODO : Credential 외의 다른 변수들 활용.

//Client is
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

	c.Credential.BuildAuthParams(method, path)

	return New(operation, c.Credential, response, handler)
}
