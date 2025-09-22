package app

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/v1adis1av28/level2/tasks/task16/internal/downloader"
	"github.com/v1adis1av28/level2/tasks/task16/internal/filesaver"
	"github.com/v1adis1av28/level2/tasks/task16/internal/models"
	cliparser "github.com/v1adis1av28/level2/tasks/task16/internal/parser"
)

type App struct {
	Visited     map[string]bool
	Mutex       sync.Mutex
	Queue       chan models.Job
	WG          sync.WaitGroup
	IsCompleted bool
}

func StartApp(cfg *models.Config) {
	targetURL, outDir := getUserInput()
	if targetURL == "" {
		log.Fatal("No URL provided")
	}
	cfg.Url = targetURL
	app := &App{
		Visited: make(map[string]bool),
		Queue:   make(chan models.Job, 1000),
	}

	baseUrl, err := url.Parse(cfg.Url)
	if err != nil {
		log.Fatal("Invalid URL:", err)
	}

	for i := 0; i < cfg.WorkersCount; i++ {
		app.WG.Add(1)
		go downloader.Worker(app, outDir, baseUrl, cfg, &app.WG, app.Queue) // Pass the queue
	}

	app.Queue <- models.Job{
		URL:   cfg.Url,
		Depth: cfg.MaxDepth,
	}

	app.waitForCompletion()
}

func getUserInput() (string, string) {
	fmt.Println("=== Web Crawler Utility ===")
	fmt.Println("Usage: wget <URL>")
	fmt.Print("Enter command: ")

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("No input")
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		log.Fatal("Empty input")
	}

	if err := cliparser.Parse(line); err != nil {
		log.Fatal("Error parsing command:", err)
	}

	tokens := strings.Fields(line)
	if len(tokens) < 2 {
		log.Fatal("Not enough arguments. Usage: wget <URL>")
	}
	url := tokens[1]

	localDir := filesaver.CreateLocalDir(url)
	return url, localDir
}

func (a *App) waitForCompletion() {
	go func() {
		a.WG.Wait()
		close(a.Queue)
		a.IsCompleted = true
	}()

	a.WG.Wait()
}

func (a *App) AddJob(job models.Job) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	if a.IsCompleted {
		return
	}

	if !a.Visited[job.URL] && job.Depth >= 0 {
		a.Visited[job.URL] = true
		select {
		case a.Queue <- job:
			fmt.Printf("Added job: %s (depth: %d)\n", job.URL, job.Depth)
		default:
			log.Printf("Queue full, skipping: %s", job.URL)
		}
	}
}
