package parser

import (
	"fmt"
	"net/url"
	"strings"
)

func Parse(str string) error {
	tokens := strings.Fields(str)

	if len(tokens) < 2 {
		return fmt.Errorf("not enough arguments. Usage: wget <URL>")
	}

	if strings.ToLower(tokens[0]) != "wget" {
		return fmt.Errorf("wrong command name. Use 'wget'")
	}

	if !isUrlValid(tokens[1]) {
		return fmt.Errorf("invalid URL: %s", tokens[1])
	}

	return nil
}

func isUrlValid(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" {
		return false
	}

	// Проверяем поддерживаемые схемы
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}
