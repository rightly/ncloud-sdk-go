## Ncloud Network Product Document

### Product And API List

- CDN+
  - Instance list 조회 - CdnPlusInstanceListRequest()
  - Purge 기록 조회
  - Purge 요청
- Global CDN
  - Instance list 조회 - GlobalCdnInstanceListRequest()
  - Purge 기록 조회 - GlobalCdnPurgeHistoryListRequest()
  - Purge 요청

### Using example

```go
import (
	"github.com/rightly/ncloud-sdk-go/internal"	
	"github.com/rightly/ncloud-sdk-go/internal/key"
    "github.com/rightly/ncloud-sdk-go/ncloud/media/network"
	"fmt"
)

// Key file path
// ~/internal/key/setkey.json 의 양식을 참고해 json key file 을 생성
var configPath ="./key/key.json"

func main() {
	// Key Load
	keyConf := key.Load(configPath)

	// Set Credentials, 발급받은 key 값 설정
	internal
	// or
	cred := ncloud.MakeCredentialFromKey(keyConf)

	// Configuration 생성 
	cfg := ncloud.LoadDefaultConfig(cred)

	// Create Product Object, Custom http.Client 사용시 파라미터 전달
    svc := cdn.New(cfg)

	// 요청 및 응답
	resp, err := req.Send()
    
    // CDN instance list 조회에 필요한 파라미터 인스턴스 생성
    param := &cdn.InstanceListRequestParam{
        CdnInstanceNo:"0",
		PageNo:1,
		PageSize:1,
		ResponseFormatType:"json",
    }
    
	// Request and indented response print
	if resp, err := req.Send(); err == nil {
		fmt.Println(resp)
	}
}
```

