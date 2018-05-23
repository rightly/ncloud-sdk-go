package cdn

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
	"reflect"
)

// TODO: [issue]Java date to convert Go time

type ResponseError struct {
	ReturnCode    string `json:"returnCode" xml:"returnCode"`
	ReturnMessage string `json:"returnMessage" xml:"returnMessage"`
}

type ApiGatewayError struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

////
//// CDN Plus
////

type CdnPlusInstanceListResponse struct {
	CdnPlusInstanceList `json:"getCdnPlusInstanceListResponse" xml:"getCdnPlusInstanceListResponse"`
	ResponseError       `json:"responseError" xml:"responseError"`
	ApiGatewayError     `json:"error"`
}

func (r *CdnPlusInstanceListResponse) String() string {
	unmarshal := "json"
	indentedString, err := ncloud.String(r, unmarshal)
	if err == nil {
		return indentedString
	}
	return reflect.TypeOf(r).String() + ".String() is failed"
}

type CdnPlusInstanceList struct {
	RequestId           string            `json:"requestId" xml:"requestId"`
	ReturnCode          string            `json:"returnCode" xml:"returnCode"`
	ReturnMessage       string            `json:"returnMessage" xml:"returnMessage"`
	TotalRows           uint              `json:"totalRows" xml:"totalRows"`
	CdnPlusInstanceList []CdnPlusInstance `json:"cdnPlusInstanceList" xml:"cdnPlusInstanceList"`
}

type CdnPlusInstance struct {
	CdnInstanceNo                 string                 `json:"cdnInstanceNo" xml:"cdnInstanceNo"`
	CdnInstanceStatus             CdnPlusStatus          `json:"cdnInstanceStatus" xml:"cdnInstanceStatus"`
	CdnInstanceOperation          CdnPlusStatus          `json:"cdnInstanceOperation" xml:"cdnInstanceOperation"`
	CdnInstanceStatusName         string                 `json:"cdnInstanceStatusName" xml:"cdnInstanceStatusName"`
	CreateDate                    string                 `json:"createDate" xml:"createDate"`
	LastModifiedDate              string                 `json:"lastModifiedDate" xml:"lastModifiedDate"`
	CdnInstanceDescription        string                 `json:"cdnInstanceDescription" xml:"cdnInstanceDescription"`
	ServiceName                   string                 `json:"serviceName" xml:"serviceName"`
	IsForLiveTranscoder           bool                   `json:"isForLiveTranscoder" xml:"isForLiveTranscoder"`
	LiveTranscoderInstanceNoList  []string               `json:"liveTranscoderInstanceNoList" xml:"liveTranscoderInstanceNoList"`
	IsForImageOptimizer           bool                   `json:"isForImageOptimizer" xml:"isForImageOptimizer"`
	ImageOptimizerInstanceNo      string                 `json:"imageOptimizerInstanceNo" xml:"imageOptimizerInstanceNo"`
	IsAvailablePartialDomainPurge bool                   `json:"isAvailablePartialDomainPurge" xml:"isAvailablePartialDomainPurge"`
	CdnPlusServiceDomainList      []CdnPlusServiceDomain `json:"cdnPlusServiceDomainList" xml:"cdnPlusServiceDomainList"`
	CdnPlusRule                                          `json:"cdnPlusRule" xml:"cdnPlusRule"`
}

type CdnPlusStatus struct {
	Code     string `json:"code" xml:"code"`
	CodeName string `json:"codeName" xml:"codeName"`
}

type CdnPlusServiceDomain struct {
	DomainId              string `json:"domainId" xml:"domainId"`
	ServiceDomainTypeCode string `json:"serviceDomainTypeCode" xml:"serviceDomainTypeCode"`
	ProtocolTypeCode      string `json:"protocolTypeCode" xml:"protocolTypeCode"`
	DefaultDomainName     string `json:"defaultDomainName" xml:"defaultDomainName"`
	UserDomainName        string `json:"userDomainName" xml:"userDomainName"`
}

type CdnPlusRule struct {
	ProtocolTypeCode                  string   `json:"protocolTypeCode" xml:"protocolTypeCode"`
	ServiceDomainTypeCode             string   `json:"serviceDomainTypeCode" xml:"serviceDomainTypeCode"`
	OriginUrl                         string   `json:"originUrl" xml:"originUrl"`
	OriginPath                        string   `json:"originPath" xml:"originPath"`
	OriginHttpPort                    int      `json:"originHttpPort" xml:"originHttpPort"`
	OriginHttpsPort                   int      `json:"originHttpsPort" xml:"originHttpsPort"`
	ForwardHostHeaderTypeCode         string   `json:"forwardHostHeaderTypeCode" xml:"forwardHostHeaderTypeCode"`
	ForwardHostHeader                 string   `json:"forwardHostHeader" xml:"forwardHostHeader"`
	CacheKeyHostNameTypeCode          string   `json:"cacheKeyHostNameTypeCode" xml:"cacheKeyHostNameTypeCode"`
	IsGzipCompressionUse              bool     `json:"isGzipCompressionUse" xml:"isGzipCompressionUse"`
	CachingOptionTypeCode             string   `json:"cachingOptionTypeCode" xml:"cachingOptionTypeCode"`
	IsErrorContentsResponseUse        bool     `json:"isErrorContentsResponseUse" xml:"isErrorContentsResponseUse"`
	CachingTtlTime                    int      `json:"cachingTtlTime" xml:"cachingTtlTime"`
	IsQueryStringIgnoreUse            bool     `json:"isQueryStringIgnoreUse" xml:"isQueryStringIgnoreUse"`
	IsRemoveVaryHeaderUse             bool     `json:"isRemoveVaryHeaderUse" xml:"isRemoveVaryHeaderUse"`
	IsLargeFileOptimizationUse        bool     `json:"isLargeFileOptimizationUse" xml:"isLargeFileOptimizationUse"`
	GzipResponseTypeCode              string   `json:"gzipResponseTypeCode" xml:"gzipResponseTypeCode"`
	IsReferrerDomainUse               bool     `json:"isReferrerDomainUse" xml:"isReferrerDomainUse"`
	ReferrerDomainList                []string `json:"referrerDomainList" xml:"referrerDomainList"`
	IsReferrerDomainRestrictUse       bool     `json:"isReferrerDomainRestrictUse" xml:"isReferrerDomainRestrictUse"`
	IsSecureTokenUse                  bool     `json:"isSecureTokenUse" xml:"isSecureTokenUse"`
	SecureTokenPassword               string   `json:"secureTokenPassword" xml:"secureTokenPassword"`
	IsReissueSecureTokenPassword      bool     `json:"isReissueSecureTokenPassword" xml:"isReissueSecureTokenPassword"`
	CertificateName                   string   `json:"certificateName" xml:"certificateName"`
	IsAccessLogUse                    bool     `json:"isAccessLogUse" xml:"isAccessLogUse"`
	AccessLogFileStorageContainerName string   `json:"accessLogFileStorageContainerName" xml:"accessLogFileStorageContainerName"`
}

type CdnPlusPurgeHistoryListResponse struct {
	CdnPlusPurgeHistory
	ResponseError   `json:"responseError" xml:"responseError"`
	ApiGatewayError `json:"error"`
}

func (r *CdnPlusPurgeHistoryListResponse) String() string {
	unmarshal := "json"
	indentedString, err := ncloud.String(r, unmarshal)
	if err == nil {
		return indentedString
	}

	return reflect.TypeOf(r).String() + ".String() is failed"
}

type CdnPlusPurgeHistory struct {
	RequestId               string                `json:"requestId" xml:"requestId"`
	ReturnCode              string                `json:"returnCode" xml:"returnCode"`
	ReturnMessage           string                `json:"returnMessage" xml:"returnMessage"`
	TotalRows               uint                  `json:"totalRows" xml:"totalRows"`
	CdnPlusPurgeHistoryList []struct {
		CdnInstanceNo            string                 `json:"cdnInstanceNo" xml:"cdnInstanceNo"`
		PurgeId                  string                 `json:"purgeId" xml:"purgeId"`
		IsWholePurge             bool                   `json:"isWholePurge" xml:"isWholePurge"`
		IsWholeDomain            bool                   `json:"isWholeDomain" xml:"isWholeDomain"`
		CdnPlusServiceDomainList []CdnPlusServiceDomain `json:"cdnPlusServiceDomainList" xml:"cdnPlusServiceDomainList"`
		TargetDirectoryName      string                 `json:"targetDirectoryName" xml:"targetDirectoryName"`
		TargetFileList           []string               `json:"targetFileList" xml:"targetFileList"`
		RequestDate              string                 `json:"requestDate" xml:"requestDate"`
		PurgeStatusName          string                 `json:"purgeStatusName" xml:"purgeStatusName"`
	}`json:"cdnPlusPurgeHistoryList" xml:"cdnPlusPurgeHistoryList"`
}

////
//// Global CDN
////

type GlobalCdnInstanceListResponse struct {
	GlobalCdnInstanceList `json:"getGlobalCdnInstanceListResponse" xml:"getGlobalCdnInstanceListResponse"`
	ResponseError         `json:"responseError" xml:"responseError"`
	ApiGatewayError       `json:"error"`
}

func (r *GlobalCdnInstanceListResponse) String() string {
	unmarshal := "json"
	indentedString, err := ncloud.String(r, unmarshal)
	if err == nil {
		return indentedString
	}
	return reflect.TypeOf(r).String() + ".String() is failed"
}

type GlobalCdnInstanceList struct {
	RequestId             string              `json:"requestId" xml:"requestId"`
	ReturnCode            string              `json:"returnCode" xml:"returnCode"`
	ReturnMessage         string              `json:"returnMessage" xml:"returnMessage"`
	TotalRows             uint                `json:"totalRows" xml:"totalRows"`
	GlobalCdnInstanceList []GlobalCdnInstance `json:"globalCdnInstanceList" xml:"globalCdnInstanceList"`
}

type GlobalCdnInstance struct {
	CdnInstanceNo                 string                   `json:"cdnInstanceNo" xml:"cdnInstanceNo"`
	CdnInstanceStatusName         string                   `json:"cdnInstanceStatusName" xml:"cdnInstanceStatusName"`
	CreateDate                    string                   `json:"createDate" xml:"createDate"`
	LastModifiedDate              string                   `json:"lastModifiedDate" xml:"lastModifiedDate"`
	CdnInstanceDescription        string                   `json:"cdnInstanceDescription" xml:"cdnInstanceDescription"`
	ServiceName                   string                   `json:"serviceName" xml:"serviceName"`
	IsAvailablePartialDomainPurge bool                     `json:"isAvailablePartialDomainPurge" xml:"isAvailablePartialDomainPurge"`
	GlobalCdnServiceDomainList    []GlobalCdnServiceDomain `json:"globalCdnServiceDomainList" xml:"globalCdnServiceDomainList"`
	GlobalCdnRule                                          `json:"globalCdnRule" xml:"globalCdnRule"`
}

type GlobalCdnServiceDomain struct {
	ServiceDomainTypeCode string `json:"serviceDomainTypeCode" xml:"serviceDomainTypeCode"`
	ProtocolTypeCode      string `json:"protocolTypeCode" xml:"protocolTypeCode"`
	DefaultDomainName     string `json:"defaultDomainName" xml:"defaultDomainName"`
	UserDomainName        string `json:"userDomainName" xml:"userDomainName"`
}

type GlobalCdnRule struct {
	ProtocolTypeCode                  string   `json:"protocolTypeCode" xml:"protocolTypeCode"`
	ServiceDomainTypeCode             string   `json:"serviceDomainTypeCode" xml:"serviceDomainTypeCode"`
	OriginUrl                         string   `json:"originUrl" xml:"originUrl"`
	OriginPath                        string   `json:"originPath" xml:"originPath"`
	OriginHttpPort                    int      `json:"originHttpPort" xml:"originHttpPort"`
	OriginHttpsPort                   int      `json:"originHttpsPort" xml:"originHttpsPort"`
	ForwardHostHeaderTypeCode         string   `json:"forwardHostHeaderTypeCode" xml:"forwardHostHeaderTypeCode"`
	ForwardHostHeader                 string   `json:"forwardHostHeader" xml:"forwardHostHeader"`
	CacheKeyHostNameTypeCode          string   `json:"cacheKeyHostNameTypeCode" xml:"cacheKeyHostNameTypeCode"`
	IsGzipCompressionUse              bool     `json:"isGzipCompressionUse" xml:"isGzipCompressionUse"`
	CachingOptionTypeCode             string   `json:"cachingOptionTypeCode" xml:"cachingOptionTypeCode"`
	IsErrorContentsResponseUse        bool     `json:"isErrorContentsResponseUse" xml:"isErrorContentsResponseUse"`
	CachingTtlTime                    int      `json:"cachingTtlTime" xml:"cachingTtlTime"`
	IsQueryStringIgnoreUse            bool     `json:"isQueryStringIgnoreUse" xml:"isQueryStringIgnoreUse"`
	IsRemoveVaryHeaderUse             bool     `json:"isRemoveVaryHeaderUse" xml:"isRemoveVaryHeaderUse"`
	IsLargeFileOptimizationUse        bool     `json:"isLargeFileOptimizationUse" xml:"isLargeFileOptimizationUse"`
	GzipResponseTypeCode              string   `json:"gzipResponseTypeCode" xml:"gzipResponseTypeCode"`
	IsReferrerDomainUse               bool     `json:"isReferrerDomainUse" xml:"isReferrerDomainUse"`
	ReferrerDomainList                []string `json:"referrerDomainList" xml:"referrerDomainList"`
	IsReferrerDomainRestrictUse       bool     `json:"isReferrerDomainRestrictUse" xml:"isReferrerDomainRestrictUse"`
	IsSecureTokenUse                  bool     `json:"isSecureTokenUse" xml:"isSecureTokenUse"`
	SecureTokenPassword               string   `json:"secureTokenPassword" xml:"secureTokenPassword"`
	IsReissueSecureTokenPassword      bool     `json:"isReissueSecureTokenPassword" xml:"isReissueSecureTokenPassword"`
	IsAccessLogUse                    bool     `json:"isAccessLogUse" xml:"isAccessLogUse"`
	AccessLogFileStorageContainerName string   `json:"accessLogFileStorageContainerName" xml:"accessLogFileStorageContainerName"`
}

type GlobalCdnPurgeHistoryListResponse struct {
	GlobalCdnPurgeHistory //`json:"getGlobalCdnPurgeHistoryList" xml:"getGlobalCdnPurgeHistoryList"`
	ResponseError   `json:"responseError" xml:"responseError"`
	ApiGatewayError `json:"error"`
}

func (r *GlobalCdnPurgeHistoryListResponse) String() string {
	unmarshal := "json"
	indentedString, err := ncloud.String(r, unmarshal)
	if err == nil {
		return indentedString
	}
	return reflect.TypeOf(r).String() + ".String() is failed"
}

type GlobalCdnPurgeHistory struct {
	RequestId                 string                  `json:"requestId" xml:"requestId"`
	ReturnCode                string                  `json:"returnCode" xml:"returnCode"`
	ReturnMessage             string                  `json:"returnMessage" xml:"returnMessage"`
	TotalRows                 uint                    `json:"totalRows" xml:"totalRows"`
	GlobalCdnPurgeHistoryList []struct {
		CdnInstanceNo              string                   `json:"cdnInstanceNo" xml:"cdnInstanceNo"`
		PurgeId                    string                   `json:"purgeId" xml:"purgeId"`
		IsWholePurge               bool                     `json:"isWholePurge" xml:"isWholePurge"`
		IsWholeDomain              bool                     `json:"isWholeDomain" xml:"isWholeDomain"`
		GlobalCdnServiceDomainList []GlobalCdnServiceDomain `json:"cdnPlusServiceDomainList" xml:"cdnPlusServiceDomainList"`
		TargetDirectoryName        string                   `json:"targetDirectoryName" xml:"targetDirectoryName"`
		TargetFileList             []string                 `json:"targetFileList" xml:"targetFileList"`
		RequestDate                string                   `json:"requestDate" xml:"requestDate"`
		PurgeStatusName            string                   `json:"purgeStatusName" xml:"purgeStatusName"`
	}
}

///
/// Request struct
///

type InstanceListRequestParam struct {
	CdnInstanceNo      string `json:"cdnInstanceNo"`
	PageNo             int    `json:"pageNo"`
	PageSize           int    `json:"pageSize"`
	ResponseFormatType string `json:"responseFormatType"`
}

type PurgeHistoryListRequestParam struct {
	CdnInstanceNo string `json:"cdnInstanceNo"`
	PurgeListN    string `json:"purgeList.N"`
	ResponseFormatType string `json:"responseFormatType"`
}