package geolocation

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

const (
	geolocation = sdkVersion + "geoLocation"
)

type GeoLocationRequest struct {
	*ncloud.Request
}

func (r *GeoLocationRequest) Send() (*GeolocationResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GeolocationResponse), nil
}

func (c *GeoLocation) GeoLocationRequest(p *GeolocationRequestParam) GeoLocationRequest {
	const opName = "GeoLocation"

	path := geolocation +
		"?ip=" + p.IP + "&enc=" + p.Enc + "&ext=" + p.Ext + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Name:opName,
		Credentials: "apigw",
		Method:"GET",
		Path:path,
		Url:endpoint + path,
	}

	response := &GeolocationResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return GeoLocationRequest{Request: req}

}