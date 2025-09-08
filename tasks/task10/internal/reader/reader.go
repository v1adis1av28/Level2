package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadStdin() []string {
	strs := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		strs = append(strs, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading error", err)
		return nil
	}

	return strs
}
