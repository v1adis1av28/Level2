package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/v1adis1av28/level2/tasks/task10/internal/sort"
)

func ReadStdin() []sort.Line {
	strs := []sort.Line{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := parseColumns(strings.ToLower(line))
		strs = append(strs, sort.Line{Columns: columns, BaseStr: line})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading error", err)
		return nil
	}
	return strs
}

func parseColumns(line string) []sort.Column {
	parts := strings.Split(line, "\t")
	columns := make([]sort.Column, len(parts))

	for i, part := range parts {
		columns[i] = sort.Column{Number: i + 1, Line: part}
	}

	return columns
}
