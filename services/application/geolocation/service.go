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

// GeoLocation OpenAPI endpoint
const endpoint = "https://ncloud.apigw.ntruss.com" + sdkVersion

func New(cfg *ncloud.Config) *GeoLocation {

	svc := &GeoLocation{
		Client: ncloud.NewClient(cfg),
	}

	return svc
}