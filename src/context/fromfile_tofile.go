//go:build fromfile_tofile
// +build fromfile_tofile

package context

import "fmt"

func init() {
	fmt.Println("fromfile_tofile")
}

func Load() ([]string, error) {
	fmt.Println("load from file")
	return nil, nil
}

func Save() error {
	fmt.Println("save to file")
	return nil
}
