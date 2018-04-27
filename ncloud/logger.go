package ncloud

import (
	"log"
	"net/http"
)

type Logger struct {
	*log.Logger
}

func WithLogger(l *Logger) Decorator {
	return func(c HttpClient) HttpClient {
		return HandlerFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s : %s : %s %s", r.RemoteAddr, r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

