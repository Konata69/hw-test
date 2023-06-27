package main

import (
	"os"
	"reflect"
	"testing"
)

const tmpFolder = "./tmp"

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func makeFile(name string, body string) {
	file, err := os.Create(tmpFolder + "/" + name)
	checkErr(err)
	defer file.Close()

	_, err = file.WriteString(body)
	checkErr(err)
}

func setUp() {
	err := os.Mkdir(tmpFolder, os.ModePerm)
	checkErr(err)

	makeFile("BAR", "bar")
	makeFile("EMPTY", "")
	makeFile("SPACE", "space  ")
	makeFile("NULL", "null\x00")
	makeFile("EQUAL=", "equal")
}

func tearDown() {
	err := os.RemoveAll(tmpFolder)
	checkErr(err)
}

func TestReadDir(t *testing.T) {
	expected := Environment{
		"BAR": {
			Value:      "bar",
			NeedRemove: false,
		},
		"EMPTY": {
			Value:      "",
			NeedRemove: true,
		},
		"SPACE": {
			Value:      "space",
			NeedRemove: false,
		},
		"NULL": {
			Value:      "null\n",
			NeedRemove: false,
		},
		"EQUAL": {
			Value:      "equal",
			NeedRemove: false,
		},
	}

	setUp()
	defer tearDown()
	result, err := ReadDir(tmpFolder)
	checkErr(err)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result:\n %v\n\n expected:\n %v", result, expected)
	}
}
