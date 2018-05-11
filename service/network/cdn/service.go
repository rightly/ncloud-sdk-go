package cdn

import "github.com/rightly/ncloud-sdk-go/ncloud"

type Cdn struct {
	*ncloud.Client
}

// CDN OpenAPI End Point
const endpoint = "https://ncloud.apigw.ntruss.com"

// Used for custom client initialization logic
var initClient func(*Cdn)

// Used for custom request initialization logic
var initRequest func(*Cdn, *ncloud.Request)

func New(cfg *ncloud.Config) *Cdn {

	svc := &Cdn{
		Client: ncloud.NewClient(cfg),
	}

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a CDN operation and runs any
// custom request initialization.
func (c *Cdn) newRequest(operation *ncloud.Operation, response interface{}, client *ncloud.HttpHandler) *ncloud.Request  {
	req := c.NewRequest(operation, response, client)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}