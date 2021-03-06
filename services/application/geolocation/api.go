package geolocation

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

const (
	geolocation = "geoLocation"
)

type GeoLocationRequest struct {
	*ncloud.Request
}

func (r *GeoLocationRequest) Do() (*GeolocationResponse, error) {
	err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GeolocationResponse), nil
}

func (c *GeoLocation) GeoLocationRequest(p *GeolocationRequestParam) GeoLocationRequest {
	path := geolocation +
		"?ip=" + p.IP + "&enc=" + p.Enc + "&ext=" + p.Ext + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &GeolocationResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return GeoLocationRequest{Request: req}
}