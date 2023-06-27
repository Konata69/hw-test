package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	expectedEnvName := "BAR"
	expectedEnvValue := "bar"
	expectedReturnCode := 0

	returnCode := RunCmd(
		[]string{"echo"},
		map[string]EnvValue{
			"BAR": {
				Value:      "bar",
				NeedRemove: false,
			},
		})

	os.LookupEnv(expectedEnvName)
	require.Equal(t, os.Getenv(expectedEnvName), expectedEnvValue)

	if returnCode != expectedReturnCode {
		t.Errorf("ReturnCode\n result: %v\n expected: %v", returnCode, expectedReturnCode)
	}
}
