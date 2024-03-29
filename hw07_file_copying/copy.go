package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrCopyToSameFile        = errors.New("copy to same file")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == toPath {
		return ErrCopyToSameFile
	}
	if limit == 0 {
		byteBuffer, err := os.ReadFile(fromPath)
		if err != nil {
			return err
		}

		err = os.WriteFile(toPath, byteBuffer, 0o600)
		if err != nil {
			return err
		}

		return err
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	fileInfo, err := fileFrom.Stat()
	if err != nil {
		return err
	}
	size := fileInfo.Size()

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		fileFrom.Seek(offset, 0)
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileTo.Close()

	bar := pb.Simple.Start64(limit)
	defer bar.Finish()
	barReader := bar.NewProxyReader(fileFrom)

	_, err = io.CopyN(fileTo, barReader, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}
