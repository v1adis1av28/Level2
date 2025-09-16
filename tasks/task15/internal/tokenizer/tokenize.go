package tokenizer

import (
	"fmt"
	"strings"
)

func Tokenize(str string) ([]string, error) {
	if len(str) == 0 {
		return []string{}, fmt.Errorf("Empty string")
	}

	tokens := make([]string, 0)
	var builder strings.Builder
	isEscaped := false
	isSingleQuoteOpen, isDoubleQuoteOpen := false, false

	for i := 0; i < len(str); i++ {
		currChar := str[i]

		if isEscaped {
			isEscaped = false
			builder.WriteByte(currChar)
			continue
		}

		if currChar == '\\' {
			if isSingleQuoteOpen {
				if i+1 < len(str) && str[i+1] == '\'' {
					isEscaped = true
					continue
				}
				builder.WriteByte(currChar)
			} else {
				if i+1 < len(str) {
					nextChar := str[i+1]
					if nextChar == '"' || nextChar == '\'' || nextChar == '\\' || nextChar == ' ' || nextChar == '|' {
						isEscaped = true
						continue
					}
				}
				builder.WriteByte(currChar)
			}
			continue
		}

		if currChar == '\'' {
			if isSingleQuoteOpen {
				isSingleQuoteOpen = false
				tokens = append(tokens, builder.String())
				builder.Reset()
			} else if !isDoubleQuoteOpen {
				isSingleQuoteOpen = true
			} else {
				builder.WriteByte(currChar)
			}
			continue
		}

		if currChar == '"' {
			if isDoubleQuoteOpen {
				isDoubleQuoteOpen = false
				tokens = append(tokens, builder.String())
				builder.Reset()
			} else if !isSingleQuoteOpen {
				isDoubleQuoteOpen = true
			} else {
				builder.WriteByte(currChar)
			}
			continue
		}

		if currChar == ' ' && !isSingleQuoteOpen && !isDoubleQuoteOpen {
			if builder.Len() > 0 {
				tokens = append(tokens, builder.String())
				builder.Reset()
			}
			continue
		}
		builder.WriteByte(currChar)
	}

	if builder.Len() > 0 {
		tokens = append(tokens, builder.String())
	}

	if isDoubleQuoteOpen || isSingleQuoteOpen {
		return nil, fmt.Errorf("unclosed quote")
	}

	return tokens, nil
}
