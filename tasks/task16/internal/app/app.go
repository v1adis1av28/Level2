package app

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/v1adis1av28/level2/tasks/task16/internal/downloader"
	filesaver "github.com/v1adis1av28/level2/tasks/task16/internal/fileSaver"
	cliparser "github.com/v1adis1av28/level2/tasks/task16/internal/parser/cliParser"
)

type Config struct {
	MaxDepth   int
	MaxWorkers int
	Timeout    time.Duration
	UserAgent  string
}

type Crawler struct {
	config  Config
	visited *sync.Map
	queue   chan *Task
	client  *http.Client
	baseURL *url.URL
	baseDir string
}

type Task struct {
	URL   string
	Depth int
}

func NewCrawler(cfg Config, uri string) *Crawler {
	baseUrl, err := url.Parse(uri)
	if err != nil {
		fmt.Println("Error on creating crawler on parsing base uri")
		return nil
	}
	return &Crawler{
		config:  cfg,
		visited: &sync.Map{},
		queue:   make(chan *Task),
		client:  downloader.CreateHTTPClient(cfg.Timeout),
		baseDir: "./data",
		baseURL: baseUrl,
	}
}

func (cr *Crawler) StartApp() {
	fmt.Println("Start of wget utility!")
	fmt.Println("Write wget command \"URL of resourse you want to download\"")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		err := cliparser.Parse(line)
		if err != nil {
			fmt.Println("error on parsing cli:", err.Error())
			os.Exit(1)
		}
		localDir, err := filesaver.CreateLocalDir()
		if err != nil {
			fmt.Println("error on creating local directory : ", err.Error())
			os.Exit(1)
		}
		fmt.Println(localDir)
	}
	//скачивание страниц
	testErr := cr.Download("https://dzen.ru")
	if testErr != nil {
		fmt.Println("Error in download")
		os.Exit(1)
	}
}

func (cr *Crawler) Download(url string) error {
	resp, err := cr.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error on executing get request")
	}
	f, err := os.Create(cr.baseDir + "someName.html")
	fmt.Println(f.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
