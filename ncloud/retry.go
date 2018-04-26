package ncloud

import (
	"time"
	"net/http"
)

type Retryer struct {
	Delay       time.Duration
	MaxRetries  int
	ShouldRetry bool
}

func WithRetryer(rty *Retryer) Decorator {
	return func(c HttpClient) HttpClient {
		return HandlerFunc(func(r *http.Request) (resp *http.Response, err error) {
			for i := 0; i <= rty.MaxRetries; i++ {
				if resp, err = c.Do(r);
				resp.StatusCode == 200 ||
				resp.StatusCode == 404 ||
				resp.StatusCode == 403 ||
				resp.StatusCode == 401 ||
				resp.StatusCode == 500 ||
				rty.ShouldRetry == false {
					break
				}
				time.Sleep(rty.Delay)
			}
			return resp, err
		})
	}
}