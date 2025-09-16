package handler

import (
	"fmt"
	"slices"
)

var builInCommands []string = []string{"cd", "pwd", "echo", "kill", "ps"}

func HandleSingleCommand(tkns []string) error {
	if !slices.Contains(builInCommands, tkns[0]) {
		return fmt.Errorf("unknown command")
	}
	return nil
}
