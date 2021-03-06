package ncloud

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type Request struct {
	Operation    *Operation
	ExpireTime   time.Duration
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	HTTPHandler  *HttpHandler
	Retryable    *bool
	Delay        time.Time
	MaxRetries   int
	Body         []byte
	Error        error
	Data         interface{}
}

type Operation struct {
	Version     string
	Credentials string
	Method      string
	Path        string
	Url         string
	Param       interface{}
}

func New(operation *Operation, credentials *Credentials, response interface{}, client *HttpHandler) *Request {
	// method 미지정시 Default : GET
	method := operation.Method
	url := operation.Url
	byteBody := &bytes.Buffer{}

	if operation.Param != nil {
		jsonBody, _ := json.Marshal(operation.Param)
		byteBody = bytes.NewBuffer(jsonBody)
	}

	httpReq, err := http.NewRequest(method, url, byteBody)

	// api gateway 인증이 필요한 API 일 경우 request에 인증헤더 추가
	if operation.Credentials == "apigw" {
		httpReq.Header["x-ncp-apigw-timestamp"] = []string{credentials.Timestamp}
		httpReq.Header["x-ncp-apigw-api-key"] = []string{credentials.ApiKey}
		httpReq.Header["x-ncp-iam-access-key"] = []string{credentials.AccessKey}
		httpReq.Header["x-ncp-apigw-signature-v1"] = []string{credentials.Signature}
		httpReq.Header["Content-Type"] = []string{"application/json"}
	}

	req := &Request{
		Operation:   operation,
		HTTPRequest: httpReq,
		Error:       err,
		Data:        response,
		HTTPHandler: client,
	}

	return req
}

//Send is Request 정책에 대해 설정
func (r *Request) Do() (err error) {
	r.HTTPResponse, err = r.HTTPHandler.Do(r)
	if r.HTTPResponse != nil {
		defer r.HTTPResponse.Body.Close()
	}
	if err != nil {
		return
	}

	r.Body, err = ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return
	}

	size := len(r.Body)
	if size > 0 {
		err = Unmarshal(r)
		if err != nil {
			return
		}
	}

	return nil
}
