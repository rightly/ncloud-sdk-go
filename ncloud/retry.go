package ncloud

import (
	"time"
	"net/http"
	"sort"
)

// TODO: Documenting, Retryer 구현

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
					NoRetryCode(resp.StatusCode) == true ||
					rty.ShouldRetry == false {
					break
				}
				time.Sleep(rty.Delay)
			}
			return resp, err
		})
	}
}

func NoRetryCode(statusCode int) bool {
	noRetryCode := sort.IntSlice{200, 400, 401, 403, 404, 500, 503}

	i := sort.Search(len(noRetryCode), func(i int) bool {
		return noRetryCode[i] >= statusCode
	})

	if i < len(noRetryCode) && noRetryCode[i] == statusCode {
		return true
	} else {
		return false
	}
}