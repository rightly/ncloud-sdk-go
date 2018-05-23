package cdn

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
	"strconv"
)

const (
	cdnPlusInstanceList = "getCdnPlusInstanceList"
	cdnPlusPurgeHistoryList = "getCdnPlusPurgeHistoryList"
	gcdnInstanceList = "getGlobalCdnInstanceList"
	gcdnPurgeHistoryList = "getGlobalCdnPurgeHistoryList"
)

type CdnPlusInstanceListRequest struct {
	*ncloud.Request
}

func (r *CdnPlusInstanceListRequest) Do() (*CdnPlusInstanceListResponse, error) {
	err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	return r.Data.(*CdnPlusInstanceListResponse), nil
}

func (c *Cdn) CdnPlusInstanceListRequest(p *InstanceListRequestParam) CdnPlusInstanceListRequest {
	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := cdnPlusInstanceList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageNo=" + strconv.Itoa(p.PageNo) + "&pageSize=" + strconv.Itoa(p.PageSize) + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &CdnPlusInstanceListResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return CdnPlusInstanceListRequest{Request: req}
}

type CdnPlusPurgeHistoryListRequest struct {
	*ncloud.Request
}

func (r *CdnPlusPurgeHistoryListRequest) Do() (*CdnPlusPurgeHistoryListResponse, error) {
	err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	return r.Data.(*CdnPlusPurgeHistoryListResponse), nil
}

func (c *Cdn) CdnPlusPurgeHistoryListRequest(p *PurgeHistoryListRequestParam) CdnPlusPurgeHistoryListRequest {
	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := cdnPlusPurgeHistoryList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageSize=" + p.PurgeListN + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &CdnPlusPurgeHistoryListResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return CdnPlusPurgeHistoryListRequest{Request:req}
}

type GlobalCdnInstanceListRequest struct {
	*ncloud.Request
}

func (r *GlobalCdnInstanceListRequest) Do() (*GlobalCdnInstanceListResponse, error) {
	err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GlobalCdnInstanceListResponse), nil
}

func (c *Cdn) GlobalCdnInstanceListRequest(p *InstanceListRequestParam) GlobalCdnInstanceListRequest {
	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := gcdnInstanceList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageNo=" + strconv.Itoa(p.PageNo) + "&pageSize=" + strconv.Itoa(p.PageSize) + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &GlobalCdnInstanceListResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return GlobalCdnInstanceListRequest{Request: req}
}

type GlobalCdnPurgeHistoryListRequest struct {
	*ncloud.Request
}

func (r *GlobalCdnPurgeHistoryListRequest) Do() (*GlobalCdnPurgeHistoryListResponse, error) {
	err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GlobalCdnPurgeHistoryListResponse), nil
}

func (c *Cdn) GlobalCdnPurgeHistoryListRequest(p *PurgeHistoryListRequestParam) GlobalCdnPurgeHistoryListRequest {
	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := gcdnPurgeHistoryList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageSize=" + p.PurgeListN + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &GlobalCdnPurgeHistoryListResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return GlobalCdnPurgeHistoryListRequest{Request: req}
}