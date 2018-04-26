package geolocation

import "encoding/xml"

//JSON Response를 위한 struct
type JsonResponse struct {
	ReturnCode    int          `json:"returnCode"`
	ReturnMessage string       `json:"returnMessage"`
	RequestId     string       `json:"requestId"`
	GeoLocation   JsonLocation `json:"geolocation"`
}

type ResponseError struct {
	ResponseError Status `json:"responseError"`
}

type Status struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
}

type JsonLocation struct {
	Country string  `json:"country"`
	Code    string  `json:"code"`
	R1      string  `json:"r1"`
	R2      string  `json:"r2"`
	R3      string  `json:"r3"`
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Net     string  `json:"net"`
}

//XML Response 를 위한 struct
type XmlResponse struct {
	XMLName       xml.Name    `xml:"getLocationResponse"`
	ReturnCode    int         `xml:"returnCode"`
	ReturnMessage string      `xml:"returnMessage"`
	RequestId     string      `xml:"requestId"`
	GeoLocation   XmlLocation `xml:"geoLocation"`
}

type XmlResponseError struct {
	XMLName       xml.Name `xml:"responseError"`
	ReturnCode    string   `xml:"returnCode"`
	ReturnMessage string   `xml:"returnMessage"`
}

type XmlLocation struct {
	Country string  `xml:"country"`
	Code    string  `xml:"code"`
	R1      string  `xml:"r1"`
	R2      string  `xml:"r2"`
	R3      string  `xml:"r3"`
	Lat     float64 `xml:"lat"`
	Long    float64 `xml:"long"`
	Net     string  `xml:"net"`
}

