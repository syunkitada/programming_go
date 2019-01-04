package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Reader struct {
	offset int64
	file   *os.File
	reader *bufio.Reader
}

func NewReader(path string, maxOffset int64) (*Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReaderSize(file, 1024)

	var offset int64
	if maxOffset > 0 {
		_, err = file.Seek(-1*maxOffset, os.SEEK_END)
		if err != nil {
			return nil, err
		}
		_, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		offset, err = file.Seek(0, os.SEEK_CUR)
		if err != nil {
			return nil, err
		}
	} else {
		offset = 0
	}

	return &Reader{
		offset: offset,
		file:   file,
		reader: reader,
	}, nil
}

func (reader *Reader) Tail() error {
	_, err := reader.file.Seek(reader.offset, 0)
	if err != nil {
		return err
	}

	for {
		line, err := reader.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else {
				fmt.Println("err:", err)
				break
			}
		}
		line = strings.TrimRight(line, "\n")
		fmt.Println(line)
	}

	reader.offset, err = reader.file.Seek(0, os.SEEK_CUR)
	return nil
}

func main() {
	reader, err := NewReader("/tmp/test.log", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		err := reader.Tail()
		fmt.Printf("end, seep 10 secconds, err=%v\n", err)
		time.Sleep(time.Second * 10)
	}

	fmt.Println("end")
}
