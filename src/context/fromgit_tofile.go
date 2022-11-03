//go:build git2file
// +build git2file

package context

import "fmt"

func init() {
	fmt.Println("mode: git2file")
}

func Load() ([]string, error) {
	fmt.Println("load from git")
	return nil, nil
}

func Save() error {
	fmt.Println("save to file")
	return nil
}
