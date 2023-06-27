package main

import (
	"os"
	"os/exec"
)

func prepareCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
}

func setEnv(env Environment) (returnCode int) {
	for key, value := range env {
		err := os.Unsetenv(key)
		if err != nil {
			return 1
		}

		if value.NeedRemove {
			continue
		}

		err = os.Setenv(key, value.Value)
		if err != nil {
			return 1
		}
	}

	return 0
}

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here.
	command := exec.Command(cmd[0], cmd[1:]...)

	returnCode = setEnv(env)
	if returnCode != 0 {
		return returnCode
	}
	prepareCommand(command)

	err := command.Run()
	if err != nil {
		return 1
	}

	return 0
}
