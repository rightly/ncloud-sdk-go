package geolocation

import (
	"github.com/rightly/ncloud-sdk-go/internal"
)

// GeoLocation is provides the API Operation methods for making requests to
// Ncloud Application Product Geo Location.

const serviceName = "ncloud-application-geolocation"

// GeoLocation
type GeoLocation struct {
	*internal.Client
}

// GeoLocation OpenAPI endpoint
const endpoint = "https://ncloud.apigw.ntruss.com" + sdkVersion

func New(cfg *internal.Config) *GeoLocation {

	svc := &GeoLocation{
		Client: internal.NewClient(cfg),
	}

	return svc
}