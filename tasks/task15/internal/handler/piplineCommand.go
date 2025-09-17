package handler

import (
	"fmt"
	"os"
	"os/exec"
)

func HandlePiplineCommand(tkns []string) error {
	commands := splitPiplines(tkns)
	if len(commands) == 0 {
		return fmt.Errorf("no commands in pipeline")
	}

	cmds := make([]*exec.Cmd, len(commands))
	for i, cmdTokens := range commands {
		if len(cmdTokens) == 0 {
			return fmt.Errorf("empty command in pipeline")
		}

		if isBuiltInCommand(cmdTokens[0]) {
			return fmt.Errorf("builtin command '%s' cannot be used in pipeline", cmdTokens[0])
		}

		cmds[i] = exec.Command(cmdTokens[0], cmdTokens[1:]...)
	}

	for i := 0; i < len(cmds)-1; i++ {
		stdoutPipe, err := cmds[i].StdoutPipe()
		if err != nil {
			return fmt.Errorf("creating stdout pipe: %w", err)
		}
		cmds[i+1].Stdin = stdoutPipe
	}

	cmds[0].Stdin = os.Stdin
	cmds[len(cmds)-1].Stdout = os.Stdout
	cmds[len(cmds)-1].Stderr = os.Stderr

	for i, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			for j := 0; j < i; j++ {
				if cmds[j].Process != nil {
					cmds[j].Process.Kill()
				}
			}
			return fmt.Errorf("starting command %d: %w", i, err)
		}
	}

	var lastError error
	for i, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			if i == len(cmds)-1 {
				lastError = err
			}
		}
	}

	return lastError

}

func splitPiplines(tkns []string) [][]string {
	var commands [][]string
	var currentCommand []string

	for _, token := range tkns {
		if token == "|" {
			if len(currentCommand) > 0 {
				commands = append(commands, currentCommand)
				currentCommand = []string{}
			}
		} else {
			currentCommand = append(currentCommand, token)
		}
	}

	if len(currentCommand) > 0 {
		commands = append(commands, currentCommand)
	}

	return commands
}
