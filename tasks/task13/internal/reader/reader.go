package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/v1adis1av28/level2/tasks/task13/internal/cut"
)

func ReadLine(delimetr string) []cut.Line {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(delimetr)
	strs := make([]cut.Line, 0)
	for scanner.Scan() {
		text := scanner.Text()
		line := cut.Line{Text: text}
		if strings.Contains(text, delimetr) {
			line.IsSeparated = true
		} else {
			line.IsSeparated = false
		}
		strs = append(strs, line)
	}
	return strs
}
