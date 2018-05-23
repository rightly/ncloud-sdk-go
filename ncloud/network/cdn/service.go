package cdn

import 	"github.com/rightly/ncloud-sdk-go/internal"

type Cdn struct {
	*internal.Client
}

// CDN OpenAPI Endpoint
const endpoint = "https://ncloud.apigw.ntruss.com" + sdkVersion

func New(cfg *internal.Config) *Cdn {

	svc := &Cdn{
		Client: internal.NewClient(cfg),
	}

	return svc
}