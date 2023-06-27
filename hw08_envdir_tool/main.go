package main

import (
	"errors"
	"fmt"
	"os"
)

func returnErrorCode(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func main() {
	// Place your code here.
	args := os.Args
	if len(args) < 2 {
		returnErrorCode(errors.New("invalid args count"))
	}

	path := args[1]
	command := args[2:]

	env, err := ReadDir(path)
	if err != nil {
		returnErrorCode(err)
	}

	os.Exit(RunCmd(command, env))
}
