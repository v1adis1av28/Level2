package filesaver

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dataDir string = "./data"

func CreateLocalDir(url string) string {
	addStr := SanitizeHost(url)
	outputDir := filepath.Join(dataDir, addStr)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
		return ""
	}
	return outputDir
}

func SanitizeHost(host string) string {
	if i := strings.Index(host, ":"); i != -1 {
		host = host[:i]
	}
	return strings.ReplaceAll(host, ".", "_")
}
