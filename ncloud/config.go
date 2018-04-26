package ncloud

import (
	"net/http"
	"time"
	"log"
	"os"
)

type Config struct {
	Client *http.Client
	*Retryer
	*Logger
	*Credentials
}

func LoadDefaultConfig(cred *Credentials) *Config {
	if cred == nil {
		return nil
	}

	keepAliveTimeout := 600 * time.Second
	timeout := 2 * time.Second
	defaultTransport := &http.Transport{
		MaxIdleConns:        1024,
		TLSHandshakeTimeout: keepAliveTimeout,
	}
	defaultClient := &http.Client{
		Timeout:   timeout,
		Transport: defaultTransport,
	}

	defaultRetryer := &Retryer{
		Delay:       30 * time.Second,
		MaxRetries:  3,
		ShouldRetry: true,
	}

	defaultLogger := &Logger{
		Logger: log.New(os.Stdout, "[Client]: ", log.LstdFlags),
	}

	return &Config{
		Client:      defaultClient,
		Retryer:     defaultRetryer,
		Logger:      defaultLogger,
		Credentials: cred,
	}
}