package cliparser

import (
	"fmt"
	"net/url"
	"strings"
)

// Пока не принимаем глубину рекурсии и поэтому будет только два элемента команада + ссылка
func Parse(str string) error {

	tokens := strings.Fields(str)

	if strings.Compare(tokens[0], "wget") != 0 {
		return fmt.Errorf("wrong command name")
	}

	//пропарсить второй аргумент проверить корректность ссылки
	if !isUrlValid(tokens[1]) {
		return fmt.Errorf("invalid url")
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
	return true
}
