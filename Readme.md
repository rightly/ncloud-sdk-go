# ncloud-sdk-go
For using Ncloud openAPI in Go

## Getting started

go get 을 사용해 추가합니다
```sh
go get github.com/rightly/ncloud-sdk-go
```

## Compute

## Network

## AI & Application

- Geolocation

## Media

- VOD Transcoder
    - Job Operation
        - Job 생성 - CreateJobRequest()
        - Job 생성 취소 - JobCreateCancelRequest()
        - Job 목록 조회 - JobListRequest()
        - Job 정보 조회 - JobInfoRequest()

    - Preset Operation
        - Preset 목록 조회 - PresetListRequest()
        - Preset 정보 조회 - PresetInfoRequest()

- Live Transcoder

### Using example
```go
import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
	"github.com/rightly/ncloud-sdk-go/service/media/VODTranscoder"
	"github.com/rightly/ncloud-sdk-go/key"
	"fmt"
)

// Key file path
var configPath ="./key/key.json"

func main() {
	// Key Load
	keyConf := key.Load(configPath)

	// Set Credentials, 발급받은 key 값 설정
	cred := ncloud.SetCredential(keyConf.Key.ApiKey,keyConf.Key.AccessKey,keyConf.Key.SecretKey)

	// Configuration 생성 
	cfg := ncloud.LoadDefaultConfig(cred)

	// Create Product Object, Custom http.Client 사용시 파라미터 전달
	svc := VODTranscoder.New(cfg)

	// Job 생성 요청을 위한 파라미터 객체 생성
	inputs := []VODTranscoder.CreateJobInput{
		{
			InputContainerName: "vt-storage",
			InputFilePath:      "/test.mp4",
		},
	}

	outputFiles := []VODTranscoder.OutputFile{
		{
			PresetId:       "9bc226df-04c9-11e8-8379-00505685080f",
			OutputFileName: "api-test",
		},
	}
	output := VODTranscoder.CreateJobOutput{
		OutputContainerName:    "vt-storage",
		ThumbnailOn:            "true",
		ThumbnailContainerName: "vt-thumb",
		OutputFiles:            outputFiles,
	}

	createJobParam := &VODTranscoder.CreateJobParam{
		JobName:jobNmae,
		Inputs: inputs,
		Output:output,
	}

	// 위 생성한 객체를 파라미터로 CreateJobRequest 생성
	req := svc.CreateJobRequest(createJobParam)

	// 요청 및 응답
	resp, err := req.Send()

	if err == nil {
		fmt.Println(resp.String())
	}
}
```
