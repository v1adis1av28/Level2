package tokenizer

import (
	"fmt"
	"strings"
)

// //test cases:
// echo "hello world"
// echo 'single quoted'
// echo hello\ world
// echo \"quoted\" -> "quoted"
// echo \| -> |

func Tokenize(str string) ([]string, error) {
	if len(str) == 0 {
		return []string{}, fmt.Errorf("Empty string")
	}
	tokenize := make([]string, 0)
	var builder strings.Builder
	return tokenize, nil
}
