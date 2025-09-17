package reader

import (
	"bufio"
	"fmt"
	"os"

	"github.com/v1adis1av28/level2/tasks/task15/internal/parser"
)

func ReadLines() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return fmt.Errorf("Error on parsing command line err")
		}
		err := parser.ParseLine(line)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
