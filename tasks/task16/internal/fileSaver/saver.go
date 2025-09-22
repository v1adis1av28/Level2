package filesaver

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateLocalDir(url string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	baseDir := filepath.Join(homeDir, "web_crawler_data")

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Printf("Warning: Could not create %s, trying current directory: %v", baseDir, err)
		baseDir = "./web_crawler_data"
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	domain := sanitizeDomain(url)
	outputDir := filepath.Join(baseDir, domain)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	log.Printf("Created output directory: %s", outputDir)
	return outputDir
}

func sanitizeDomain(urlStr string) string {
	domain := urlStr
	if strings.Contains(domain, "://") {
		domain = strings.Split(domain, "://")[1]
	}
	if strings.Contains(domain, "/") {
		domain = strings.Split(domain, "/")[0]
	}

	domain = strings.ReplaceAll(domain, ":", "_")
	domain = strings.ReplaceAll(domain, ".", "_")
	domain = strings.ReplaceAll(domain, "-", "_")

	return domain
}

func EnsureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, 0755)
}

func SaveFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}
