package ncloud

import (
	"net/http"
	"time"
	"log"
)

// TODO: Documenting

type Config struct {
	Client *http.Client
	LogLevel
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

	defaultLogLevel := INFO

	defaultRetryer := &Retryer{
		Delay:       30 * time.Second,
		MaxRetries:  3,
		ShouldRetry: true,
	}

	defaultLogger := &Logger{
		Logger:NewLogger("ncloudapi.log", "", log.Ldate|log.Ltime|log.Lshortfile),
		logLevel:defaultLogLevel,
	}

	return &Config{
		Client:      defaultClient,
		LogLevel:    defaultLogLevel,
		Retryer:     defaultRetryer,
		Logger:      defaultLogger,
		Credentials: cred,
	}
}

func LoadConfig(cred *Credentials, logLevel int, r *Retryer, l *Logger) *Config {
	if cred == nil {
		return nil
	}


	return &Config {

	}
}