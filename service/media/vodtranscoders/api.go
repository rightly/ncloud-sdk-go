package vodtranscoders

import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
)

// Actions
const (
	jobAction =  SDKVersion + "jobs"
	presetAction = SDKVersion + "presets"
)

/*
 * JobAPI
 * - Job 생성
 * - Job 생성 취소
 * - Job Info 조회
 * - Job List 조회
*/

// CreateJobRequest is Job create request API Operation
type CreateJobRequest struct {
	*ncloud.Request
}

// Send marshals and sends the Job create API request.
func (r *CreateJobRequest)Send() (*CreateJobResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if r.Data.(*CreateJobResponse).ResponseStatus == "" {
		r.Data.(*CreateJobResponse).ResponseStatus = r.HTTPResponse.Status
	}

	return r.Data.(*CreateJobResponse), nil
}

// CreateJobRequest returns a Job create request Operation for
// Ncloud VOD Transcoder.
//
// 	// Example
// 	// CreateJobParam 객체 생성
// 	jobNmae := "api-test"
//	inputs := []vodtranscoder.CreateJobInput{
//		{
//			InputContainerName: "vt-storage",
//			InputFilePath:      "/test.mp4",
//		},
//	}
//
//	outputFiles := []vodtranscoder.OutputFile{
//		{
//			PresetId:       "9bc226df-04c9-11e8-8379-00505685080f",
//			OutputFileName: "api-test",
//		},
//	}
//	output := vodtranscoder.CreateJobOutput{
//		OutputContainerName:    "vt-storage",
//		ThumbnailOn:            "true",
//		ThumbnailContainerName: "vt-thumb",
//		OutputFiles:            outputFiles,
//	}
//
//	createJobParam := &vodtranscoder.CreateJobParam{
//		JobName:jobNmae,
//		Inputs: inputs,
//		Output:output,
//	}
//
//  // CreateJobRequest 생성
//  req := svc.CreateJobRequest(createJobParam)
//
//	// API 요청
//	resp, err := req.Send()
//
// 	//resp.String() -> Unmarshal string data for JobCreateResponse
//	if err == nil {
//		fmt.Println(resp.String())
//	}
//
func (c *VodTranscoder)CreateJobRequest(p *CreateJobParam) CreateJobRequest {

	const opName = "CreateJobRequest"

	op := &ncloud.Operation{
		Name:opName,
		Credential: "apigw",
		Method:"POST",
		Path:jobAction,
		Url:END_POINT + jobAction,
		Param:p,
	}

	response := &CreateJobResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return CreateJobRequest{Request: req}
}

//JobCreateCancelRequest is Job create cancel request API Operation
type JobCreateCancelRequest struct {
	*ncloud.Request
}

// Send marshals and sends the Job create cancel API request.
func (r JobCreateCancelRequest)Send() (*JobCreateCancelResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if r.Data.(*JobCreateCancelResponse).ResponseStatus == "" {
		r.Data.(*JobCreateCancelResponse).ResponseStatus = r.HTTPResponse.Status
	}
	return r.Data.(*JobCreateCancelResponse), nil
}

// JobCreateCancelRequest returns a Job create cancel request Operation for
// Ncloud VOD Transcoder.
//
// 	// Example
//	// string 형 JobId 를 파라미터로 JobCreateCancelRequest 생성
//	cancelJob := svc.JobCreateCancelRequest("jbdexlcozdlhjka0zlof8qp4q2fg1h7t")
//
// 	// API 요청
//	resp, err := cancelJob.Send()
//
//	//resp.String() -> Unmarshal string data for JobCreateCancelResponse
//	if err == nil {
//		fmt.Println(resp.String())
//	}
//
func (c *VodTranscoder)JobCreateCancelRequest(jobId string) JobCreateCancelRequest {
	const opName = "JobCreateCancelRequest"
	jobCancelAction := jobAction + "/" + jobId + "/cancel"

	op := &ncloud.Operation{
		Name:opName,
		Credential: "apigw",
		Method:"POST",
		Path:jobCancelAction,
		Url:END_POINT + jobCancelAction,
	}

	response := &JobCreateCancelResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return JobCreateCancelRequest{Request: req}
}

// JobListRequest is View information of job list Request API Operation
type JobListRequest struct {
	*ncloud.Request
}

// Send marshals and sends the View information of job list API request.
func (r JobListRequest)Send() (*JobListResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Data.(*JobListResponse), nil
}

// JobListRequest returns a View information of job list request Operation for
// Ncloud VOD Transcoder.
//
// 	// Example
// 	// JobListRequest 생성
// 	req := svc.JobListRequest()
//
//	// API 요청
// 	resp, err := req.Send()
//
//	// resp.String() -> Unmarshal string data for JobListResponse
//	if err == nil {
//		fmt.Println(resp.String())
//	}
//
func (c *VodTranscoder)JobListRequest() JobListRequest {

	const opName = "JobListRequest"

	op := &ncloud.Operation{
		Name:opName,
		Credential: "apigw",
		Method:"GET",
		Path:jobAction,
		Url:END_POINT + jobAction,
	}

	response := &JobListResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return JobListRequest{Request: req}
}


// JobInfoRequest is View information of specific job request API Operation
type JobInfoRequest struct {
	*ncloud.Request
}

// Send marshals and sends the View information of specific job API request.
func (r *JobInfoRequest)Send() (*JobInfoResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if r.Data.(*JobInfoResponse).ResponseStatus == "" {
		r.Data.(*JobInfoResponse).ResponseStatus = r.HTTPResponse.Status
	}
	return r.Data.(*JobInfoResponse), nil
}

// JobInfoRequest returns a View information of specific job request Operation for
// Ncloud VOD Transcoder.
//
// 	// Example
//	// JobInfoRequest 생성
// 	req := svc.JobInfoRequest("jobId")
//
//	// API 요청
// 	resp, err := req.Send()
//
//	// resp.String() -> Unmarshal string data for JobInfoResponse
//	if err == nil {
//		fmt.Println(resp.String())
//	}
//
func (c *VodTranscoder)JobInfoRequest(jobId string) JobInfoRequest {

	const opName = "JobInfoRequest"
	url := jobAction + "/" + jobId

	op := &ncloud.Operation{
		Name:   opName,
		Credential: "apigw",
		Method: "GET",
		Path:   url,
		Url:    END_POINT + url,
	}

	response := &JobInfoResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return JobInfoRequest{Request: req}
}

/*
 * PresetAPI
 * - Preset Info 조회
 * - Preset List 조회
*/

// PresetListRequest is View information of preset list request API Operation
type PresetListRequest struct {
	*ncloud.Request
}

// Send marshals and sends the View information of preset list API request.
func (r *PresetListRequest)Send() (*PresetListResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if r.Data.(*PresetListResponse).ResponseStatus == "" {
		r.Data.(*PresetListResponse).ResponseStatus = r.HTTPResponse.Status
	}
	return r.Data.(*PresetListResponse), nil
}

// JobListRequest returns a View information of preset list request Operation for
// Ncloud VOD Transcoder.
//
// 	// Example
// 	// JobListRequest 생성
// 	req := svc.PresetListRequest()
//
//	// API 요청
// 	resp, err := req.Send()
//
//	// resp.String() -> Unmarshal string data for PresetListResponse
//	if err == nil {
//		fmt.Println(resp.String())
//	}
//
func (c *VodTranscoder)PresetListRequest() PresetListRequest {
	const opName= "PresetListRequest"

	op := &ncloud.Operation{
		Name:   opName,
		Credential: "apigw",
		Method: "GET",
		Path:   presetAction,
		Url:    END_POINT + presetAction,
	}

	response := &PresetListResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return PresetListRequest{Request: req}
}

// PresetInfoRequest is View information of specific preset request API Operation
type PresetInfoRequest struct {
	*ncloud.Request
}

// Send marshals and sends the View information of specific preset API request.
func (r *PresetInfoRequest)Send() (*PresetInfoResponse, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}
	if r.Data.(*PresetInfoResponse).ResponseStatus == "" {
		r.Data.(*PresetInfoResponse).ResponseStatus = r.HTTPResponse.Status
	}
	return r.Data.(*PresetInfoResponse), nil
}

// JobListRequest returns a View information of specific preset request Operation for
// Ncloud VOD Transcoder.
//
// 	// Example
// 	// string 형 presetId 를 파라미터로 PresetInfoRequest 생성
// 	req := svc.PresetInfoRequest("presetId")
//
//	// API 요청
// 	resp, err := req.Send()
//
//	// resp.String() -> Unmarshal string data for PresetInfoResponse
//	if err == nil {
//		fmt.Println(resp.String())
//	}
//
func (c *VodTranscoder)PresetInfoRequest(presetId string) PresetInfoRequest {
	const opName= "PresetInfoRequest"
	url := presetAction + "/" + presetId

	op := &ncloud.Operation{
		Name:   opName,
		Credential: "apigw",
		Method: "GET",
		Path:   url,
		Url:    END_POINT + url,
	}

	response := &PresetInfoResponse{}
	handler := c.Handler
	req := c.newRequest(op, response, handler)

	return PresetInfoRequest{Request: req}
}