package cdn

import (
	"../../../ncloud"
	"strconv"
)

const (
	cdnPlusInstanceList = sdkVersion + "getCdnPlusInstanceList"
	gcdnInstanceList = sdkVersion + "getGlobalCdnInstanceList"
	gcdnPurgeHistoryList = sdkVersion + "getGlobalCdnPurgeHistoryList"
)

type CdnPlusInstanceListRequest struct {
	*ncloud.Request
}

func (r *CdnPlusInstanceListRequest) Send() (*CdnPlusInstanceListResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Data.(*CdnPlusInstanceListResponse), nil
}

func (c *Cdn) CdnPlusInstanceListRequest(p *InstanceListRequestParam) CdnPlusInstanceListRequest {
	const opName = "CdnPlusInstanceList"

	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := cdnPlusInstanceList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageNo=" + strconv.Itoa(p.PageNo) + "&pageSize=" + strconv.Itoa(p.PageSize) + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Name:opName,
		Credentials: "apigw",
		Method:"GET",
		Path:path,
		Url:endpoint + path,
	}

	response := &CdnPlusInstanceListResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return CdnPlusInstanceListRequest{Request: req}
}

type GlobalCdnInstanceListRequest struct {
	*ncloud.Request
}

func (r *GlobalCdnInstanceListRequest) Send() (*GlobalCdnInstanceListResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GlobalCdnInstanceListResponse), nil
}

func (c *Cdn) GlobalCdnInstanceListRequest(p *InstanceListRequestParam) GlobalCdnInstanceListRequest {
	const opName = "GlobalCdnInstanceList"

	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := gcdnInstanceList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageNo=" + strconv.Itoa(p.PageNo) + "&pageSize=" + strconv.Itoa(p.PageSize) + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Name:opName,
		Credentials: "apigw",
		Method:"GET",
		Path:path,
		Url:endpoint + path,
	}

	response := &GlobalCdnInstanceListResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return GlobalCdnInstanceListRequest{Request: req}
}

type GlobalCdnPurgeHistoryListRequest struct {
	*ncloud.Request
}

func (r *GlobalCdnPurgeHistoryListRequest) Send() (*GlobalCdnPurgeHistoryListResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Data.(*GlobalCdnPurgeHistoryListResponse), nil
}

func (c *Cdn) GlobalCdnPurgeHistoryListRequest(p *PurgeHistoryListRequestParam) GlobalCdnPurgeHistoryListRequest {
	const opName = "GlobalCdnPurgeHistoryList"

	if p.ResponseFormatType == "" {
		p.ResponseFormatType = "json"
	}

	path := gcdnPurgeHistoryList +
		"?cdnInstanceNo=" + p.CdnInstanceNo + "&pageSize=" + p.PurgeListN + "&responseFormatType=" + p.ResponseFormatType

	op := &ncloud.Operation{
		Name:opName,
		Credentials: "apigw",
		Method:"GET",
		Path:path,
		Url:endpoint + path,
	}

	response := &GlobalCdnPurgeHistoryListResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return GlobalCdnPurgeHistoryListRequest{Request: req}
}