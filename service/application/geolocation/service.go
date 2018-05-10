package geolocation

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

// GeoLocation is provides the API Operation methods for making requests to
// Ncloud Application Product Geo Location.

const serviceName = "ncloud-application-geolocation"

// GeoLocation
type GeoLocation struct {
	*ncloud.Client
}

// GeoLocation OpenAPI End Point
const endpoint = "https://ncloud.apigw.ntruss.com"

// Used for custom client initialization logic
var initClient func(*GeoLocation)

// Used for custom request initialization logic
var initRequest func(*GeoLocation, *ncloud.Request)

func New(cfg *ncloud.Config) *GeoLocation {

	svc := &GeoLocation{
		Client: ncloud.NewClient(cfg),
	}

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a GeoLocation operation and runs any
// custom request initialization.
func (c *GeoLocation) newRequest(operation *ncloud.Operation, response interface{}, client *ncloud.HttpHandler) *ncloud.Request  {
	req := c.NewRequest(operation, response, client)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}