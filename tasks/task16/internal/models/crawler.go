package models

import (
	"net/http"
	"net/url"
	"sync"
)

type Crawler struct {
	baseURL       *url.URL
	visitedURLs   sync.Map
	downloadQueue chan string
	maxDepth      int
	client        *http.Client
	baseDir       string
}
