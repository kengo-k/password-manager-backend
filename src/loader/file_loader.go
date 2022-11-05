package loader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type FileLoader struct{}

func (l *FileLoader) Load() ([]string, error) {
	f, err := os.Open("./password.md")
	if err != nil {
		return nil, fmt.Errorf("failed to open from file")
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 1024)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read from file")
		}
		//s := string(line)
		//fmt.Print(s)
	}

	return nil, fmt.Errorf("foo")
}
