# ncloud-sdk-go
For using NAVER CLOUD PLATFORM openAPI in Go

## Getting started

go get 을 사용해 추가합니다
```sh
go get github.com/rightly/ncloud-sdk-go
```

## Compute

## Network

* CDN
  * CDN+
    - Instance list 조회 - CdnPlusInstanceListRequest()
    - Purge 기록 조회
    - Purge 요청
  * Global CDN
    - Instance list 조회 - GlobalCdnInstanceListRequest()
    - Purge 기록 조회 - GlobalCdnPurgeHistoryListRequest()
    - Purge 요청

## AI & Application

- Geolocation
  - 위치정보조회 
    - GeoLocationRequest() // Json response only

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
	"github.com/rightly/ncloud-sdk-go/service/media/vodtranscoder"
	"github.com/rightly/ncloud-sdk-go/key"
	"fmt"
)

// Key file path
// ~/key/setkey.json 의 양식을 참고해 json key file 을 생성
var configPath ="./key/key.json"

func main() {
	// Key Load
	keyConf := key.Load(configPath)

	// Set Credentials, 발급받은 key 값 설정
	// cred := ncloud.MakeCredential(keyConf.Key.ApiKey,keyConf.Key.AccessKey,keyConf.Key.SecretKey)
	// or
	cred := ncloud.MakeCredentialFromKey(keyConf)
	
	// Configuration 인스턴스 생성
	cfg := ncloud.LoadDefaultConfig(cred)
    
	// Product 인스턴스 생성
	svc := vodranscoder.New(cfg)
    
	// Job 생성 요청을 위한 파라미터 인스턴스 생성
	inputs := []vodranscoder.CreateJobInput{
		{
			InputContainerName: "vt-storage",
			InputFilePath:      "/test.mp4",
		},
	}

	outputFiles := []vodranscoder.OutputFile{
		{
			PresetId:       "9bc226df-04c9-11e8-8379-00505685080f",
			OutputFileName: "api-test",
		},
	}
	output := vodranscoder.CreateJobOutput{
		OutputContainerName:    "vt-storage",
		ThumbnailOn:            "true",
		ThumbnailContainerName: "vt-thumb",
		OutputFiles:            outputFiles,
	}

	createJobParam := &vodranscoder.CreateJobParam{
		JobName:jobNmae,
		Inputs: inputs,
		Output:output,
	}

	// 위 생성한 객체를 파라미터로 CreateJobRequest 생성
	req := svc.CreateJobRequest(createJobParam)

	// Request and indented response print
	if resp, err := req.Send(); err == nil {
		fmt.Println(resp)
	}
}
```

