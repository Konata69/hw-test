package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	if limit == 0 {
		byteBuffer, err := os.ReadFile(from)
		checkErr(err)

		err = os.WriteFile(to, byteBuffer, 0644)
		checkErr(err)

		return
	}

	fileFrom, err := os.Open(from)
	checkErr(err)
	defer fileFrom.Close()

	fileTo, err := os.Create(to)
	checkErr(err)
	defer fileTo.Close()

	_, err = io.CopyN(fileTo, fileFrom, limit)
	if err == io.EOF {
		err = nil
	}
	checkErr(err)
}
