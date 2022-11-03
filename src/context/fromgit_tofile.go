//go:build fromgit_tofile
// +build fromgit_tofile

package context

import "fmt"

func init() {
	fmt.Println("fromgit_tofile")
}

func Load() ([]string, error) {
	fmt.Println("load from git")
	return nil, nil
}

func Save() error {
	fmt.Println("save to file")
	return nil
}
