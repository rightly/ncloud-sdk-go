package geolocation

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
	"reflect"
	"encoding/xml"
)

//JSON Response를 위한 struct
type GeolocationResponse struct {
	XMLName       xml.Name      `xml:"GeolocationResponse"`
	ReturnCode    int           `json:"returnCode" xml:"returnCode"`
	RequestId     string        `json:"requestId" xml:"requestId"`
	GeoLocation   Location      `json:"geolocation" xml:"geolocation"`
	Error         ErrorResponse `json:"responseError" xml:"responseError"`
}

type ErrorResponse struct {
	XMLName       xml.Name `xml:"responseError"`
	ReturnCode    string   `json:"returnCode" xml:"returnCode"`
	ReturnMessage string   `json:"returnMessage" xml:"returnMessage"`
}

type Location struct {
	XMLName xml.Name `xml:"geolocation"`
	Country string   `json:"country" xml:"country"`
	Code    string   `json:"code" xml:"code"`
	R1      string   `json:"r1" xml:"r1"`
	R2      string   `json:"r2" xml:"r2"`
	R3      string   `json:"r3" xml:"r3"`
	Lat     float64  `json:"lat" xml:"lat"`
	Long    float64  `json:"long" xml:"long"`
	Net     string   `json:"net" xml:"net"`
}

type GeolocationParam struct {
	IP                 string
	Enc                string
	Ext                string
	ResponseFormatType string
}

func (r *GeolocationResponse) String() string {
	unmarshal := "json"
	indentedString, err := ncloud.String(r, unmarshal)
	if err == nil {
		return indentedString
	}
	return reflect.TypeOf(r).String() + ".String() is failed"
}