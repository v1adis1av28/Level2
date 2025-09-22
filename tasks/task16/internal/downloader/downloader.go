package downloader

import (
	"net/http"
	"time"
)

func CreateHTTPClient(duration time.Duration) *http.Client {
	return &http.Client{
		Timeout: duration,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxConnsPerHost:     10,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}
