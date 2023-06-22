package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if limit == 0 {
		byteBuffer, err := os.ReadFile(fromPath)
		checkErr(err)

		err = os.WriteFile(toPath, byteBuffer, 0644)
		checkErr(err)

		return err
	}

	fileFrom, err := os.Open(fromPath)
	defer fileFrom.Close()
	checkErr(err)

	fileInfo, err := fileFrom.Stat()
	checkErr(err)
	size := fileInfo.Size()

	if offset > size {
		log.Fatal(ErrOffsetExceedsFileSize)
	}

	if offset > 0 {
		fileFrom.Seek(offset, 0)
	}

	fileTo, err := os.Create(toPath)
	defer fileTo.Close()
	checkErr(err)

	bar := pb.Simple.Start64(limit)
	barReader := bar.NewProxyReader(fileFrom)

	_, err = io.CopyN(fileTo, barReader, limit)
	if err == io.EOF {
		err = nil
	}
	checkErr(err)

	bar.Finish()

	return nil
}
