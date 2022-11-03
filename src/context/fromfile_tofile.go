//go:build file2file
// +build file2file

package context

import "fmt"

func init() {
	fmt.Println("mode: file2file")
}

func Load() ([]string, error) {
	fmt.Println("load from file")
	return nil, nil
}

func Save() error {
	fmt.Println("save to file")
	return nil
}
