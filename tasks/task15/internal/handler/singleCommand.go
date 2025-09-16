package handler

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

var builInCommands []string = []string{"cd", "pwd", "echo", "kill", "ps"}

func HandleSingleCommand(tkns []string) error {
	if !slices.Contains(builInCommands, tkns[0]) {
		return fmt.Errorf("unknown command")
	}

	switch tkns[0] {
	case "cd":
		if len(tkns) < 2 {
			return fmt.Errorf("cd missing arguments")
		}
		_, err := os.Stat(tkns[1])
		if os.IsNotExist(err) {
			return fmt.Errorf("directory not exist")
		}
		return os.Chdir(tkns[1])
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(tkns[1:], " "))
		return nil
	case "ps":
		proccesList, err := ps.Processes()
		if err != nil {
			return err
		}
		for _, process := range proccesList {
			fmt.Println("Process id: ----- process name:", process.Pid(), process.Executable())
		}
		return nil
	case "kill":
		if len(tkns) < 2 {
			return fmt.Errorf("missing argument")
		}

		pid, err := strconv.Atoi(tkns[1])
		if err != nil {
			return fmt.Errorf("wrong argument on pid argument get: %v", err)
		}
		if pid < 0 {
			return fmt.Errorf("pid value cannot be negative")
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return fmt.Errorf("process doesn`t exists")
		}
		err = proc.Kill()
		if err != nil {
			return err
		}
	}
	return nil
}
