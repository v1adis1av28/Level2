package app

import (
	"bufio"
	"fmt"
	"os"

	filesaver "github.com/v1adis1av28/level2/tasks/task16/internal/fileSaver"
	cliparser "github.com/v1adis1av28/level2/tasks/task16/internal/parser/cliParser"
)

func StartApp() {
	fmt.Println("Start of wget utility!")
	fmt.Println("Write wget command \"URL of resourse you want to download\"")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		err := cliparser.Parse(line)
		if err != nil {
			fmt.Println("error on parsing cli:", err.Error())
			os.Exit(1)
		}
		localDir, err := filesaver.CreateLocalDir()
		if err != nil {
			fmt.Println("error on creating local directory : ", err.Error())
			os.Exit(1)
		}
		fmt.Println(localDir)
	}
}
