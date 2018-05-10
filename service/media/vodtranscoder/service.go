package vodtranscoder

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

// VODTranscoder is provides the API Operation methods for making requests to
// Ncloud Media Product VOD Transcoder.

const serviceName = "ncloud-media-vodtranscoder"

// VodTranscoder
type VodTranscoder struct {
	*ncloud.Client
}

// VOD Transcoder OpenAPI End Point
const endpoint = "https://vodtranscoder.apigw.ntruss.com"

// Used for custom client initialization logic
var initClient func(*VodTranscoder)

// Used for custom request initialization logic
var initRequest func(*VodTranscoder, *ncloud.Request)

func New(cfg *ncloud.Config) *VodTranscoder {

	svc := &VodTranscoder{
		Client: ncloud.NewClient(cfg),
	}

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a VodTranscoder operation and runs any
// custom request initialization.
func (c *VodTranscoder) newRequest(operation *ncloud.Operation, response interface{}, client *ncloud.HttpHandler) *ncloud.Request  {
	req := c.NewRequest(operation, response, client)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}