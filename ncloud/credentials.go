package ncloud

import (
	"strconv"
	"time"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"github.com/rightly/ncloud-sdk-go/internal/oauth"
	"github.com/rightly/ncloud-sdk-go/internal/key"
)

// TODO: Documenting

type Credentials struct {
	Method    string
	Timestamp string
	ApiKey    string
	AccessKey string
	SecretKey string
	Signature string
	config    *oauth.Config
	token     *oauth.Token
}

// 초기 api, access, secret key 값을 받아 Credential Struct에 채워준다.
func MakeCredential(apiKey string, accessKey string, secretKey string) *Credentials {
	return &Credentials{
		ApiKey:apiKey,
		AccessKey:accessKey,
		SecretKey:secretKey,
	}
}

func MakeCredentialFromKey(c key.KeyConfig) *Credentials{
	return &Credentials{
		ApiKey:    c.Key.ApiKey,
		AccessKey: c.Key.AccessKey,
		SecretKey: c.Key.SecretKey,
	}
}

func (c *Credentials) setOauthClient(client *http.Client) {
	c.config = oauth.NewConfig(c.AccessKey, c.SecretKey)
	c.token = oauth.NewToken("", "")
	client = c.config.Client(oauth.NoContext, c.token)
}

func (c *Credentials) setAuthParams(method, apiUrl string) {

	c.Method = method
	// millisecond timestamp
	c.Timestamp = strconv.FormatInt(int64(time.Now().UnixNano()) / int64(time.Millisecond), 10)
	c.Signature = c.makeSignature(apiUrl)
}

func (c *Credentials) makeSignature(apiUrl string) string {
	var message string

	message = fmt.Sprint(message, c.Method + " ") // Request Method
	message = fmt.Sprint(message, apiUrl + "\n" ) // Uri
	message = fmt.Sprint(message, c.Timestamp + "\n") // Time Stamp
	//message = fmt.Sprint(message, c.ApiKey + "\n") // Api Key
	message = fmt.Sprint(message, c.AccessKey)

	key := []byte(c.SecretKey)	// Secret Key

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func KeyLoad(path string) key.KeyConfig {
	return key.Load(path)
}