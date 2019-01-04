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
	path      string
	maxOffset int64
	offset    int64
	file      *os.File
	reader    *bufio.Reader
}

func NewReader(path string, maxOffset int64) (*Reader, error) {
	reader := &Reader{
		path:      path,
		maxOffset: maxOffset,
	}
	err := reader.ReOpen()
	return reader, err
}

func (reader *Reader) ReOpen() error {
	var err error
	reader.file, err = os.Open(reader.path)
	if err != nil {
		return err
	}

	reader.reader = bufio.NewReaderSize(reader.file, 1024)

	if reader.maxOffset > 0 {
		var endOffset int64
		endOffset, err = reader.file.Seek(0, os.SEEK_END)
		if err != nil {
			return err
		}
		if reader.maxOffset > endOffset {
			reader.offset = 0
		} else {
			_, err = reader.file.Seek(-1*reader.maxOffset, os.SEEK_END)
			if err != nil {
				return err
			}
			_, err = reader.reader.ReadString('\n')
			if err != nil {
				return err
			}
			reader.offset, err = reader.file.Seek(0, os.SEEK_CUR)
			if err != nil {
				return err
			}
		}
	} else {
		reader.offset = 0
	}

	return nil
}

func (reader *Reader) Tail() error {
	var err error

	// logrotateなどでファイル移動された場合に対応できないため、ファイルは都度開く
	reader.file, err = os.Open(reader.path)
	defer reader.file.Close()
	if err != nil {
		return err
	}

	reader.reader = bufio.NewReaderSize(reader.file, 1024)

	endOffset, err := reader.file.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}

	// offsetとファイル末尾の位置を確認してファイルがローテされたかをチェックする
	// しかし、ファイルがローテされ、かつTailの周期よりも早く既存のofffsetよりもログが膨大になった場合検知できない
	// 実用的にはログのテキストには時間が記録されているので、パースしてそれも見てローテされたかを判定するなど必要
	if reader.offset > endOffset {
		return fmt.Errorf("Invalid offset")
	}

	_, err = reader.file.Seek(reader.offset, 0)
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
				return err
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
		if err != nil {
			fmt.Printf("err: %v\n", err)
			// ログローテなどで、ファイルが切り詰めらりたり、
			if err = reader.ReOpen(); err != nil {
				fmt.Printf("err: %v\n", err)
			} else {
				fmt.Printf("ReOpen")
			}
		}
		fmt.Printf("end, seep 10 secconds, err=%v\n", err)
		time.Sleep(time.Second * 10)
	}

	fmt.Println("end")
}
