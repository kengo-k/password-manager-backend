//go:build !file2file && !git2file
// +build !file2file,!git2file

package context

import "fmt"

func init() {
	fmt.Println("mode: git2git")
}

func Load() ([]string, error) {
	fmt.Println("load from git")
	return nil, nil
}

func Save() error {
	fmt.Println("save to git")
	return nil
}
