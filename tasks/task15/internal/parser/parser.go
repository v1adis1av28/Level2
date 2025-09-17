package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/v1adis1av28/level2/tasks/task15/internal/handler"
	"github.com/v1adis1av28/level2/tasks/task15/internal/tokenizer"
)

func ParseLine(str string) error {
	isPipeline := false
	if strings.Contains(str, "|") {
		isPipeline = true
	}
	tokens, err := tokenizer.Tokenize(str)
	if errors.Is(err, fmt.Errorf("unclosed quote")) {
		return err
	}
	if len(tokens) == 0 {
		return fmt.Errorf("empty string")
	}
	switch isPipeline {
	case true:
		err := handler.HandlePiplineCommand(tokens)
		if err != nil {
			return err
		}
	case false:
		err := handler.HandleSingleCommand(tokens)
		if err != nil {
			return err
		}
	}
	return nil
}
