// main.go
package main

import (
	"fmt"
	"os"

	"github.com/v1adis1av28/level2/tasks/task10/internal/config"
	"github.com/v1adis1av28/level2/tasks/task10/internal/reader"
	"github.com/v1adis1av28/level2/tasks/task10/internal/sort"
)

func main() {
	conf := config.ParseConfig()
	if conf == nil {
		os.Exit(1)
	}

	lines := reader.ReadStdin()
	if lines == nil {
		os.Exit(1)
	}

	fmt.Println("Config:", conf.Reverse, conf.Numeric, conf.KeyColumn, conf.Unique)
	fmt.Println("Lines read:", len(lines))

	sorted := sort.SortLines(lines, conf.KeyColumn, conf.Numeric, conf.Reverse, conf.Unique)

	for _, line := range sorted {
		fmt.Println(line)
	}
}
