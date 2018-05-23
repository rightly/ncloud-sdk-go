package vodtranscoder

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

// VODTranscoder is provides the API Operation methods for making requests to
// Ncloud Media Product VOD Transcoder.

const serviceName = "internal-media-vodtranscoder"

// VodTranscoder
type VodTranscoder struct {
	*ncloud.Client
}

// VOD Transcoder OpenAPI End Point
const endpoint = "https://vodtranscoder.apigw.ntruss.com" + sdkVersion


func New(cfg *ncloud.Config) *VodTranscoder {

	svc := &VodTranscoder{
		Client: ncloud.NewClient(cfg),
	}

	return svc
}