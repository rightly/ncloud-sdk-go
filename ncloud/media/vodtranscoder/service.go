package vodtranscoder

import (
	"github.com/rightly/ncloud-sdk-go/internal"
)

// VODTranscoder is provides the API Operation methods for making requests to
// Ncloud Media Product VOD Transcoder.

const serviceName = "internal-media-vodtranscoder"

// VodTranscoder
type VodTranscoder struct {
	*internal.Client
}

// VOD Transcoder OpenAPI End Point
const endpoint = "https://vodtranscoder.apigw.ntruss.com" + sdkVersion


func New(cfg *internal.Config) *VodTranscoder {

	svc := &VodTranscoder{
		Client: internal.NewClient(cfg),
	}

	return svc
}