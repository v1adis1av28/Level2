package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/v1adis1av28/level2/tasks/task16/internal/filesaver"
	"github.com/v1adis1av28/level2/tasks/task16/internal/models"
	"github.com/v1adis1av28/level2/tasks/task16/internal/parser"
)

type JobQueue interface {
	AddJob(job models.Job)
}

func Worker(q JobQueue, outDir string, baseUrl *url.URL, cfg *models.Config, wg *sync.WaitGroup, queue chan models.Job) {
	defer wg.Done()

	client := &http.Client{
		Timeout: cfg.Timeout,
		Transport: &http.Transport{
			MaxIdleConns:       100,
			MaxConnsPerHost:    10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: false,
			ForceAttemptHTTP2:  true,
		},
	}

	for job := range queue {
		processJob(q, client, job, outDir, baseUrl)
	}

	log.Printf("Worker finished")
}

func processJob(q JobQueue, client *http.Client, job models.Job, outDir string, baseUrl *url.URL) {
	if job.Depth < 0 {
		return
	}
	targetURL, err := url.Parse(job.URL)
	if err != nil {
		log.Printf("Invalid URL %s: %v", job.URL, err)
		return
	}

	if !sameDomain(targetURL, baseUrl) {
		log.Printf("Skipping external domain: %s", job.URL)
		return
	}

	content, contentType, err := downloadContent(client, job.URL)
	if err != nil {
		log.Printf("Download error for %s: %v", job.URL, err)
		return
	}

	filePath, err := saveContent(content, job.URL, outDir)
	if err != nil {
		log.Printf("Save error for %s: %v", job.URL, err)
		return
	}

	log.Printf("Saved: %s -> %s", job.URL, filePath)

	if strings.Contains(contentType, "text/html") && job.Depth > 0 {
		extractAndQueueLinks(q, content, job)
	}
}

func downloadContent(client *http.Client, url string) ([]byte, string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return content, resp.Header.Get("Content-Type"), nil
}

func saveContent(content []byte, urlStr, outDir string) (string, error) {
	targetURL, _ := url.Parse(urlStr)
	localPath := urlToLocalPath(targetURL)
	fullPath := filepath.Join(outDir, localPath)

	if err := filesaver.EnsureDir(fullPath); err != nil {
		return "", err
	}

	if err := os.WriteFile(fullPath, content, 0644); err != nil {
		return "", err
	}

	return fullPath, nil
}

func extractAndQueueLinks(q JobQueue, content []byte, job models.Job) {
	links, assets, err := parser.ExtractLinksAndAssets(string(content), job.URL)
	if err != nil {
		log.Printf("Parse error for %s: %v", job.URL, err)
		return
	}

	for _, link := range append(links, assets...) {
		q.AddJob(models.Job{
			URL:   link,
			Depth: job.Depth - 1,
		})
	}
}

func sameDomain(u1, u2 *url.URL) bool {
	return u1.Host == u2.Host
}

func urlToLocalPath(u *url.URL) string {
	path := u.Path
	if path == "" || path == "/" {
		return "index.html"
	}

	if path[0] == '/' {
		path = path[1:]
	}

	if len(path) > 0 && path[len(path)-1] == '/' {
		path += "index.html"
	}

	if !strings.Contains(filepath.Base(path), ".") {
		if u.RawQuery != "" {
			path += ".html"
		} else {
			path += "/index.html"
		}
	}

	return path
}
