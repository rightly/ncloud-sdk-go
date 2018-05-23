package vodtranscoder

import (
	"github.com/rightly/ncloud-sdk-go/internal"
)

// Actions
const (
	jobAction =  "jobs"
	presetAction = "presets"
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
	*internal.Request
}

// Send marshals and sends the Job create API request.
func (r *CreateJobRequest) Do() (*CreateJobResponse, error) {
	err := r.Request.Do()
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
func (c *VodTranscoder) CreateJobRequest(p *CreateJobRequestParam) CreateJobRequest {
	path := jobAction

	op := &internal.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "POST",
		Path:        path,
		Url:         endpoint + path,
		Param:       p,
	}

	response := &CreateJobResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return CreateJobRequest{Request: req}
}

//JobCreateCancelRequest is Job create cancel request API Operation
type JobCreateCancelRequest struct {
	*internal.Request
}

// Send marshals and sends the Job create cancel API request.
func (r JobCreateCancelRequest) Do() (*JobCreateCancelResponse, error) {
	err := r.Request.Do()
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
func (c *VodTranscoder) JobCreateCancelRequest(jobId string) JobCreateCancelRequest {
	path := jobAction + "/" + jobId + "/cancel"

	op := &internal.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "POST",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &JobCreateCancelResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return JobCreateCancelRequest{Request: req}
}

// JobListRequest is View information of job list Request API Operation
type JobListRequest struct {
	*internal.Request
}

// Send marshals and sends the View information of job list API request.
func (r JobListRequest) Do() (*JobListResponse, error) {
	err := r.Request.Do()
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
func (c *VodTranscoder) JobListRequest() JobListRequest {
	path := jobAction

	op := &internal.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &JobListResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return JobListRequest{Request: req}
}


// JobInfoRequest is View information of specific job request API Operation
type JobInfoRequest struct {
	*internal.Request
}

// Send marshals and sends the View information of specific job API request.
func (r *JobInfoRequest) Do() (*JobInfoResponse, error) {
	err := r.Request.Do()
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
func (c *VodTranscoder) JobInfoRequest(jobId string) JobInfoRequest {
	path := jobAction + "/" + jobId

	op := &internal.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &JobInfoResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return JobInfoRequest{Request: req}
}

/*
 * PresetAPI
 * - Preset Info 조회
 * - Preset List 조회
*/

// PresetListRequest is View information of preset list request API Operation
type PresetListRequest struct {
	*internal.Request
}

// Send marshals and sends the View information of preset list API request.
func (r *PresetListRequest) Do() (*PresetListResponse, error) {
	err := r.Request.Do()
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
func (c *VodTranscoder) PresetListRequest() PresetListRequest {
	path := presetAction

	op := &internal.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &PresetListResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return PresetListRequest{Request: req}
}

// PresetInfoRequest is View information of specific preset request API Operation
type PresetInfoRequest struct {
	*internal.Request
}

// Send marshals and sends the View information of specific preset API request.
func (r *PresetInfoRequest) Do() (*PresetInfoResponse, error) {
	err := r.Request.Do()
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
func (c *VodTranscoder) PresetInfoRequest(presetId string) PresetInfoRequest {
	path := presetAction + "/" + presetId

	op := &internal.Operation{
		Version:     sdkVersion,
		Credentials: "apigw",
		Method:      "GET",
		Path:        path,
		Url:         endpoint + path,
	}

	response := &PresetInfoResponse{}
	handler := c.Handler
	req := c.NewRequest(op, response, handler)

	return PresetInfoRequest{Request: req}
}