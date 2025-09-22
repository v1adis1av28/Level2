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
	Visited map[string]bool
	Mutex   sync.Mutex
	Queue   chan models.Job
	WG      sync.WaitGroup
}

func StartApp(cfg *models.Config) {
	// 1. Получаем URL от пользователя и создаем директорию
	targetURL, outDir := getUserInput()
	if targetURL == "" {
		log.Fatal("No URL provided")
	}
	cfg.Url = targetURL

	fmt.Printf("Starting download of: %s\n", targetURL)
	fmt.Printf("Output directory: %s\n", outDir)
	fmt.Printf("Max depth: %d, Workers: %d\n", cfg.MaxDepth, cfg.WorkersCount)

	// 2. Создаем приложение
	app := &App{
		Visited: make(map[string]bool),
		Mutex:   sync.Mutex{},
		Queue:   make(chan models.Job, 1000), // Увеличиваем буфер
		WG:      sync.WaitGroup{},
	}

	// 3. Парсим базовый URL
	baseUrl, err := url.Parse(cfg.Url)
	if err != nil {
		log.Fatal("Invalid URL:", err)
	}

	// 4. Запускаем воркеры
	for i := 0; i < cfg.WorkersCount; i++ {
		app.WG.Add(1)
		go downloader.Worker(app, outDir, baseUrl, cfg)
	}

	// 5. Добавляем первую задачу в отдельной горутине
	go func() {
		app.Queue <- models.Job{
			URL:   cfg.Url,
			Depth: cfg.MaxDepth,
		}
		fmt.Printf("Added initial job: %s (depth: %d)\n", cfg.Url, cfg.MaxDepth)
	}()

	// 6. Ждем завершения всех задач
	app.waitForCompletion()
	fmt.Printf("Download completed! Saved successfully to %s\n", outDir)
}

func getUserInput() (string, string) {
	fmt.Println("=== Web Crawler Utility ===")
	fmt.Println("Usage: wget <URL>")
	fmt.Print("Enter command: ")

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("No input provided")
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		log.Fatal("Empty input")
	}

	// Валидируем команду
	if err := cliparser.Parse(line); err != nil {
		log.Fatal("Error parsing command:", err)
	}

	// Извлекаем URL
	tokens := strings.Fields(line)
	if len(tokens) < 2 {
		log.Fatal("Not enough arguments. Usage: wget <URL>")
	}
	url := tokens[1]

	// Создаем директорию
	localDir := filesaver.CreateLocalDir(url)
	return url, localDir
}

func (a *App) waitForCompletion() {
	// Создаем канал для сигнала о завершении
	done := make(chan bool)

	// Запускаем мониторинг завершения
	go func() {
		a.WG.Wait()
		close(a.Queue)
		done <- true
	}()

	// Ждем сигнала о завершении
	<-done
}

// AddJob безопасно добавляет задачу в очередь
func (a *App) AddJob(job models.Job) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

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
