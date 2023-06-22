package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	err := Copy("testdata/input.txt", "out.txt", 0, 0)
	if err != nil {
		return
	}

	fileFrom, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		return
	}
	fileTo, err := os.ReadFile("out.txt")
	if err != nil {
		return
	}

	require.Equal(t, fileFrom, fileTo)
}

func TestCopyLimit10(t *testing.T) {
	limitBytes := int64(10)
	err := Copy("testdata/input.txt", "out.txt", 0, limitBytes)
	if err != nil {
		return
	}

	fileFrom, err := os.Open("testdata/input.txt")
	if err != nil {
		return
	}
	fileFromBytes := make([]byte, limitBytes)
	_, err = io.ReadFull(fileFrom, fileFromBytes)
	if err != nil {
		return
	}

	fileTo, err := os.ReadFile("out.txt")
	if err != nil {
		return
	}

	require.Equal(t, fileFromBytes, fileTo)
}

func TestCopyLimit1000(t *testing.T) {
	limitBytes := int64(1000)
	err := Copy("testdata/input.txt", "out.txt", 0, limitBytes)
	if err != nil {
		return
	}

	fileFrom, err := os.Open("testdata/input.txt")
	if err != nil {
		return
	}
	fileFromBytes := make([]byte, limitBytes)
	_, err = io.ReadFull(fileFrom, fileFromBytes)
	if err != nil {
		return
	}

	fileTo, err := os.ReadFile("out.txt")
	if err != nil {
		return
	}

	require.Equal(t, fileFromBytes, fileTo)
}

func TestCopyLimit10000(t *testing.T) {
	limitBytes := int64(10000)
	err := Copy("testdata/input.txt", "out.txt", 0, limitBytes)
	if err != nil {
		return
	}

	fileFrom, err := os.Open("testdata/input.txt")
	if err != nil {
		return
	}
	fileFromBytes := make([]byte, limitBytes)
	_, err = io.ReadFull(fileFrom, fileFromBytes)
	if err != nil {
		return
	}

	fileTo, err := os.ReadFile("out.txt")
	if err != nil {
		return
	}

	if len(fileFromBytes) > len(fileTo) {
		fileFromBytes = fileFromBytes[:len(fileTo)]
	}

	require.Equal(t, fileFromBytes, fileTo)
}

func TestCopyLimit1000Offset100(t *testing.T) {
	limitBytes := int64(1000)
	offsetBytes := int64(100)
	err := Copy("testdata/input.txt", "out.txt", offsetBytes, limitBytes)
	if err != nil {
		return
	}

	fileFrom, err := os.Open("testdata/input.txt")
	if err != nil {
		return
	}
	_, err = fileFrom.Seek(offsetBytes, 0)
	if err != nil {
		return
	}
	fileFromBytes := make([]byte, limitBytes)
	_, err = io.ReadFull(fileFrom, fileFromBytes)
	if err != nil {
		return
	}

	fileTo, err := os.ReadFile("out.txt")
	if err != nil {
		return
	}

	require.Equal(t, fileFromBytes, fileTo)
}

func TestCopyLimit1000Offset6000(t *testing.T) {
	limitBytes := int64(1000)
	offsetBytes := int64(6000)
	err := Copy("testdata/input.txt", "out.txt", offsetBytes, limitBytes)
	if err != nil {
		return
	}

	fileFrom, err := os.Open("testdata/input.txt")
	if err != nil {
		return
	}
	_, err = fileFrom.Seek(offsetBytes, 0)
	if err != nil {
		return
	}
	fileFromBytes := make([]byte, limitBytes)
	_, err = io.ReadFull(fileFrom, fileFromBytes)
	if err != nil {
		return
	}

	fileTo, err := os.ReadFile("out.txt")
	if err != nil {
		return
	}

	if len(fileFromBytes) > len(fileTo) {
		fileFromBytes = fileFromBytes[:len(fileTo)]
	}

	require.Equal(t, fileFromBytes, fileTo)
}
