## Ncloud Application Product Document

### Product And API List

- GeoLocation
  - 위치정보조회 - GeoLocationRequest()

### Using example

```go
import (
	"github.com/rightly/ncloud-sdk-go/ncloud"
    "github.com/rightly/ncloud-sdk-go/services/media/geolocation"
	"fmt"
)

// Key file path
// ~/internal/key/setkey.json 의 양식을 참고해 json key file 을 생성
var configPath ="./key/key.json"

func main() {
	// Key Load
	keyConf := ncloud.KeyLoad(configPath)

	// Set Credentials, 발급받은 key 값 설정
	// cred := ncloud.MakeCredential(keyConf.Key.ApiKey,keyConf.Key.AccessKey,keyConf.Key.SecretKey)
	// or
	cred := ncloud.MakeCredentialFromKey(keyConf)

	// Configuration 생성 
	cfg := ncloud.LoadDefaultConfig(cred)

	// Create Product Object, Custom http.Client 사용시 파라미터 전달
    svc := geolocation.New(cfg)

	// 요청 및 응답
	resp, err := req.Send()
    
    param := &geolocation.GeolocationRequestParam{
		IP:"127.0.0.1",
		Enc:"utf8",
		Ext:"t",
		ResponseFormatType:"json",
	}
    
	// Request and indented response print
	if resp, err := req.Send(); err == nil {
		fmt.Println(resp)
	}
}
```

