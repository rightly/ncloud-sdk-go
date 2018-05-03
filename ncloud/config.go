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

	defaultRetryer := &Retryer{
		Delay:       30 * time.Second,
		MaxRetries:  3,
		ShouldRetry: true,
	}

	defaultLogLevel := INFO

	defaultLogger := &Logger{
		Logger:NewLogger("ncloud_api.log", "", log.Ldate|log.Ltime|log.Lshortfile),
	}

	defaultLogger.SetLogLevel(defaultLogLevel)

	return &Config{
		Client:      defaultClient,
		LogLevel:    defaultLogger.LogLevel(),
		Retryer:     defaultRetryer,
		Logger:      defaultLogger,
		Credentials: cred,
	}
}

func LoadConfig(cred *Credentials, args... interface{}) *Config {
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

	defaultLogLevel := INFO

	defaultLogger := &Logger{
		Logger:NewLogger("ncloud_api.log", "", log.Ldate|log.Ltime|log.Lshortfile),
	}

	defaultLogger.SetLogLevel(defaultLogLevel)

	for _, v := range args {
		switch v.(type) {
		case *http.Client:
			defaultClient = v.(*http.Client)
		case *Logger:
			defaultLogger = v.(*Logger)
		case *Retryer:
			defaultRetryer = v.(*Retryer)
		case LogLevel:
			defaultLogger.SetLogLevel(v.(LogLevel))
		}
	}

	return &Config {
		Client:      defaultClient,
		LogLevel:    defaultLogger.LogLevel(),
		Retryer:     defaultRetryer,
		Logger:      defaultLogger,
		Credentials: cred,
	}
}