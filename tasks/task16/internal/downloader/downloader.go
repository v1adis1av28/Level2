package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/v1adis1av28/level2/tasks/task16/internal/app"
	"github.com/v1adis1av28/level2/tasks/task16/internal/models"
	"github.com/v1adis1av28/level2/tasks/task16/internal/parser"
	"github.com/v1adis1av28/level2/tasks/task16/internal/saver"
)

func Worker(app *app.App, outDir string, baseUrl *url.URL, cfg *models.Config) {
	defer app.WG.Done()

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

	log.Printf("Worker started, waiting for jobs...")

	for job := range app.Queue {
		processJob(app, client, job, outDir, baseUrl, cfg)
	}

	log.Printf("Worker finished")
}

func processJob(app *app.App, client *http.Client, job models.Job, outDir string, baseUrl *url.URL, cfg *models.Config) {
	if job.Depth < 0 {
		return
	}

	log.Printf("Processing: %s (depth: %d)", job.URL, job.Depth)

	// Проверяем URL
	targetURL, err := url.Parse(job.URL)
	if err != nil {
		log.Printf("Invalid URL %s: %v", job.URL, err)
		return
	}

	// Проверяем домен
	if !sameDomain(targetURL, baseUrl) {
		log.Printf("Skipping external domain: %s", job.URL)
		return
	}

	// Скачиваем контент
	content, contentType, err := downloadContent(client, job.URL)
	if err != nil {
		log.Printf("Download error for %s: %v", job.URL, err)
		return
	}

	// Сохраняем файл
	filePath, err := saveContent(content, job.URL, outDir, baseUrl)
	if err != nil {
		log.Printf("Save error for %s: %v", job.URL, err)
		return
	}

	log.Printf("Saved: %s -> %s", job.URL, filePath)

	// Если это HTML, парсим ссылки
	if strings.Contains(contentType, "text/html") && job.Depth > 0 {
		extractAndQueueLinks(app, content, job)
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

func saveContent(content []byte, urlStr, outDir string, baseUrl *url.URL) (string, error) {
	targetURL, _ := url.Parse(urlStr)
	localPath := urlToLocalPath(targetURL, baseUrl)
	fullPath := filepath.Join(outDir, localPath)

	// Создаем директорию если нужно
	if err := saver.EnsureDir(fullPath); err != nil {
		return "", err
	}

	if err := saver.SaveFile(fullPath, content); err != nil {
		return "", err
	}

	return fullPath, nil
}

func extractAndQueueLinks(app *app.App, content []byte, job models.Job) {
	links, assets, err := parser.ExtractLinksAndAssets(string(content), job.URL)
	if err != nil {
		log.Printf("Parse error for %s: %v", job.URL, err)
		return
	}

	// Добавляем все найденные ссылки и ресурсы
	for _, link := range append(links, assets...) {
		app.AddJob(models.Job{
			URL:   link,
			Depth: job.Depth - 1,
		})
	}
}

func sameDomain(u1, u2 *url.URL) bool {
	return u1.Host == u2.Host
}

func urlToLocalPath(u, base *url.URL) string {
	path := u.Path
	if path == "" || path == "/" {
		return "index.html"
	}

	// Убираем начальный слэш
	if path[0] == '/' {
		path = path[1:]
	}

	// Если путь заканчивается на /, добавляем index.html
	if len(path) > 0 && path[len(path)-1] == '/' {
		path += "index.html"
	}

	// Если нет расширения, добавляем .html
	if !strings.Contains(filepath.Base(path), ".") {
		// Проверяем, есть ли query parameters
		if u.RawQuery != "" {
			path += ".html"
		} else {
			path += "/index.html"
		}
	}

	return path
}
