package filesaver

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateLocalDir(url string) string {
	// Создаем основную директорию data если не существует
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Fatalf("Failed to create base directory: %v", err)
	}

	// Создаем поддиректорию на основе домена
	domain := sanitizeDomain(url)
	outputDir := filepath.Join("./data", domain)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	log.Printf("Created output directory: %s", outputDir)
	return outputDir
}

func sanitizeDomain(urlStr string) string {
	// Убираем протокол и путь, оставляем только домен
	domain := urlStr
	if strings.Contains(domain, "://") {
		domain = strings.Split(domain, "://")[1]
	}
	if strings.Contains(domain, "/") {
		domain = strings.Split(domain, "/")[0]
	}

	// Заменяем недопустимые символы
	domain = strings.ReplaceAll(domain, ":", "_")
	domain = strings.ReplaceAll(domain, ".", "_")
	domain = strings.ReplaceAll(domain, "-", "_")

	return domain
}

// EnsureDir создает директорию для файла если не существует
func EnsureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, 0755)
}
