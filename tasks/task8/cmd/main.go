package main

import (
	"fmt"
	"os"

	"github.com/v1adis1av28/Level2/tasks/task8/time"
)

// В главной функции программы, будем выполнять единственный вызов функции GetCurrentTime() для получения точного времени с помощью ntp
func main() {
	time, err := time.GetCurrentTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
