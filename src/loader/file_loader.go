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

	lines := make([]string, 1024)
	reader := bufio.NewReaderSize(f, 1024)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read from file")
		}
		lines = append(lines, string(line))
	}

	return lines, nil
}
