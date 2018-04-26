package ncloud

import (
	"log"
	"net/http"
	"strings"
	"os"
)

//TODO : LOGGER

type Logger struct {
	*log.Logger
}


func WithLogger(l *Logger) Decorator {
	return func(c HttpClient) HttpClient {
		return HandlerFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

func CheckError(e error, level string) {
	l := strings.ToLower(level)

	if e != nil {
		switch l{
		case "fatal" :
			logger := log.New(os.Stdout, level + ": ", log.LstdFlags)
			logger.Fatal(e)
		case "panic" :
			logger := log.New(os.Stdout, level + ": ", log.LstdFlags)
			logger.Panic(e)
		default :
			logger := log.New(os.Stdout, level + ": ", log.LstdFlags)
			logger.Print(e)
		}
	}
}
