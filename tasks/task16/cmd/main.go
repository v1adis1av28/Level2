package main

import (
	"github.com/v1adis1av28/level2/tasks/task16/internal/app"
)

func main() {
	cfg := app.Config{MaxDepth: 2, MaxWorkers: 10, Timeout: 10, UserAgent: " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.5 Safari/605.1.15"}
	crawler := app.NewCrawler(cfg, "https://dzen.ru")
	crawler.StartApp()
}
