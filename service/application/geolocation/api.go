package geolocation

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

const (
	geolocation = SDKVersion + "geoLocation"
)

type GeoLocationRequest struct {
	*ncloud.Request
}

func (r *GeoLocationRequest)Send() (*GeolocationResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GeolocationResponse), nil
}

func (c *GeoLocation) GeoLocationRequest(p *GeolocationParam) GeoLocationRequest {
	const opName = "GeoLocation"

	path := geolocation +
		"?ip=" + p.IP + "&enc=" + p.Enc + "&ext=" + p.Ext + "&responseFormatType=" + p.ResponseFormatType

	/*path := geolocation +
		"?ip=" + p.IP + "&enc=utf8&ext=t" + "&responseFormatType=" + p.ResponseFormatType*/

	op := &ncloud.Operation{
		Name:opName,
		Credential: "apigw",
		Method:"GET",
		Path:path,
		Url:END_POINT + path,
	}

	response := &GeolocationResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return GeoLocationRequest{Request: req}

}