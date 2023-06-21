package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
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

	fileInfo, err := fileFrom.Stat()
	checkErr(err)
	size := fileInfo.Size()

	if offset > size {
		log.Fatal(err)
	}

	if offset > 0 {
		fileFrom.Seek(offset, 0)
	}

	fileTo, err := os.Create(to)
	checkErr(err)
	defer fileTo.Close()

	bar := pb.Simple.Start64(limit)
	barReader := bar.NewProxyReader(fileFrom)

	_, err = io.CopyN(fileTo, barReader, limit)
	if err == io.EOF {
		err = nil
	}
	checkErr(err)

	bar.Finish()
}
