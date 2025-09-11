package reader

import (
	"bufio"
	"os"
	"strings"

	"github.com/v1adis1av28/level2/tasks/task13/internal/cut"
)

var DefaltDelimetr string = " "

func ReadLine(delimetr string) []cut.Line {
	scanner := bufio.NewScanner(os.Stdin)
	strs := make([]cut.Line, 0)
	isDefaultDelimetr := true
	if delimetr != DefaltDelimetr {
		isDefaultDelimetr = false
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, delimetr) {
			arr := make([]string, 0)
			ln := cut.Line{FullString: line, CollumnsArr: arr, IsSeparated: true}
			if isDefaultDelimetr {
				arr = strings.Split(line, DefaltDelimetr)
			} else {
				arr = strings.Split(line, delimetr)
			}
			ln.CollumnsArr = arr
			strs = append(strs, ln)
		} else {
			strs = append(strs, cut.Line{FullString: line, CollumnsArr: []string{line}, IsSeparated: false})
		}
	}
	return strs
}
