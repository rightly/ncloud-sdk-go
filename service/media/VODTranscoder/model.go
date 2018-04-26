package VODTranscoder

import (
	"time"
	"encoding/json"
)

// Job operation common properties
// Metadata is Video metadata properties
type Metadata struct {
	FileName string  `json:"fileName"`
	FileSize uint64  `json:"fileSize"`
	Duration float32 `json:"duration"`
	Profile  Profile `json:"profile"`
}

// Profile is Video profile properties
type Profile struct {
	VideoCodec        string `json:"videoCodec"`
	VideoBitrate      string `json:"videoBitrate"`
	Profile           string `json:"profile"`
	Width             string `json:"width"`
	Height            string `json:"height"`
	Level             string `json:"level"`
	Framerate         string `json:"framerate"`
	KeyframeInterval  string `json:"keyframeInterval"`
	AudioCodec        string `json:"audioCodec"`
	AudioBitrate      string `json:"audioBitrate"`
	AudioSamplingRate string `json:"audioSamplingRate"`
	AudioChannel      int    `json:"audioChannel"`
	ContainerFormat   string `json:"containerFormat"`
}

// OutputFile is Job create operation
type OutputFile struct {
	PresetId       string	`json:"presetId"`
	OutputFileName string	`json:"outputFileName"`
}

type Jobs struct {
	JobId         string    `json:"jobId"`
	CreateTime    time.Time `json:"createTime"`
	SubmittedTime time.Time `json:"submittedTime"`
	StartedTime   time.Time `json:"startedTime"`
	FinishedTime  time.Time `json:"finishedTime"`
	Status        string    `json:"status"`
	JobErrorCode  string    `json:"jobErrorCode"`
	JobName       string    `json:"jobName"`
	Inputs        []Input   `json:"inputs"`
	Output        Output    `json:"output"`
}

type Input struct {
	Metadata           string `json:"metadata"`
	InputContainerName string `json:"inputContainerName"`
	InputFilePath      string `json:"inputFilePath"`
}

type Output struct {
	Metadata                         `json:"metadata"`
	OutputContainerName string       `json:"outputContainerName"`
	OutputFilePath      string       `json:"outputFilePath"`
	OutputFiles         []OutputFile `json:"outputFiles"`
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

// Job request operations
// CreateJobResponse Object
type CreateJobResponse struct {
	ResponseStatus string `json:"responseStatus"`
	Error                 `json:"error"`
}

func (r *CreateJobResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "Marshal Fail"
	}
	return string(b[:])
}

type CreateJobParam struct {
	JobName string           `json:"jobName"`
	Inputs  []CreateJobInput `json:"inputs"`
	Output  CreateJobOutput  `json:"output"`
}

type CreateJobInput struct {
	InputContainerName string `json:"inputContainerName"`
	InputFilePath      string `json:"inputFilePath"`
}

type CreateJobOutput struct {
	OutputContainerName    string       `json:"outputContainerName"`
	ThumbnailOn            string       `json:"thumbnailOn"`
	ThumbnailContainerName string       `json:"thumbnailContainerName"`
	OutputFiles            []OutputFile `json:"outputFiles"`
}

// JobCreateCancelResponse Object
type JobCreateCancelResponse struct {
	ResponseStatus string `json:"responseStatus"`
	Error                 `json:"error"`
}

func (r *JobCreateCancelResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "Marshal Fail"
	}
	return string(b[:])
}

//JobListResponse Object
type JobListResponse struct {
	Jobs []Jobs `json:"jobs"`
	Error       `json:"error"`
}

func (r *JobListResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "Marshal Fail"
	}
	return string(b[:])
}

//JobInfoResponse Object
type JobInfoResponse struct {
	Jobs           []Jobs `json:"jobs"`
	Error                 `json:"error"`
	ResponseStatus string `json:"responseStatus"`
}

func (r *JobInfoResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "Marshal Fail"
	}
	return string(b[:])
}

// Preset request operations
//
type PresetOverviewList struct {
	PresetId       string        `json:"presetId"`
	PresetGroup    string        `json:"presetGroup"`
	PresetType     string        `json:"presetType"`
	PresetCostType string        `json:"presetCostType"`
	PresetName     string        `json:"presetName"`
	Profile        PresetProfile `json:"profile"`
}

//
type PresetProfile struct {
	VideoCodec        string `json:"videoCodec"`
	VideoBitrata      string `json:"videoBitrata"`
	Profile           string `json:"profile"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
	Level             string `json:"level"`
	Framerate         string `json:"framerate"`
	KeyFrameInterval  int    `json:"keyFrameInterval"`
	AudioCodec        string `json:"audioCodec"`
	AudioSamplingRate string `json:"audioSamplingRate"`
	AudioChannel      int    `json:"audioChannel"`
}

//PresetListResponse Object
type PresetListResponse struct {
	PresetOverviewList []PresetOverviewList `json:"presetOverviewList"`
	ResponseStatus     string               `json:"responseStatus"`
	Error                                   `json:"error"`
}

func (r *PresetListResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "Marshal Fail"
	}
	return string(b[:])
}

//
type PresetInfoResponse struct {
	PresetOverviewList []PresetOverviewList `json:"presetOverviewList"`
	ResponseStatus     string               `json:"responseStatus"`
	Error                                   `json:"error"`
}

func (r *PresetInfoResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "Marshal Fail"
	}
	return string(b[:])
}