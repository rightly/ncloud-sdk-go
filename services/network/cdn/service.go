package cdn

import 	"github.com/rightly/ncloud-sdk-go/ncloud"

type Cdn struct {
	*ncloud.Client
}

// CDN OpenAPI Endpoint
const endpoint = "https://ncloud.apigw.ntruss.com" + sdkVersion

func New(cfg *ncloud.Config) *Cdn {

	svc := &Cdn{
		Client: ncloud.NewClient(cfg),
	}

	return svc
}