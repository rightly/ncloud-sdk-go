package ncloud

import (
	"time"
	"net/http"
	"sort"
)

type Retryer struct {
	Delay       time.Duration
	MaxRetries  int
	ShouldRetry bool
}

func WithRetryer(rty *Retryer) Decorator {
	noRetryCode := sort.IntSlice{200, 400, 401, 403, 404, 500, 503}
	return func(c HttpClient) HttpClient {
		return HandlerFunc(func(r *http.Request) (resp *http.Response, err error) {
			for i := 0; i <= rty.MaxRetries; i++ {
				if resp, err = c.Do(r);
					sort.SearchInts(noRetryCode, resp.StatusCode) == 0 ||
					rty.ShouldRetry == false {

					break
				}
				time.Sleep(rty.Delay)
			}
			return resp, err
		})
	}
}