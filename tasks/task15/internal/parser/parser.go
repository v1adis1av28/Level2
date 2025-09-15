package parser

import (
	"fmt"
	"strings"
)

func ParseLine(str string) error {
	isPipeline := false
	if strings.Contains(str, "|") {
		isPipeline = true
	}
	tokens := strings.Fields(str)
	if len(tokens) == 0 {
		return fmt.Errorf("Empty string")
	}
	switch isPipeline {
	case true:
		handlePipline(tokens)
	case false:
		handleSingleCommand(tokens)
	}
	return nil
}

func handlePipline(tkns []string) { //здесь обрабатываем пайплайны
	fmt.Println("pipline")
}

func handleSingleCommand(tkns []string) { //сюда так как пока не реалзуем && или || просто реализуем обработку единной команды
	fmt.Println("Command compund")
}
