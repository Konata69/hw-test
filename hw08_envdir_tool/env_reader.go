package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envMap := make(Environment)
	for _, entry := range dirEntry {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return envMap, err
		}

		path := dir + "/" + info.Name()
		file, err := os.Open(path)
		if err != nil {
			return envMap, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		valueString := scanner.Text()
		valueString = string(bytes.TrimRight(bytes.ReplaceAll([]byte(valueString), []byte("\x00"), []byte("\n")), " "))

		needRemove := info.Size() == 0

		name := info.Name()
		name = strings.ReplaceAll(name, "=", "")
		value := EnvValue{
			Value:      valueString,
			NeedRemove: needRemove,
		}

		envMap[name] = value
	}

	return envMap, nil
}
