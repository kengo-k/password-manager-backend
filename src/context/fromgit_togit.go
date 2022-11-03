//go:build fromgit_togit
// +build fromgit_togit

package context

import "fmt"

func init() {
	fmt.Println("fromgit_togit")
}

func Load() ([]string, error) {
	fmt.Println("load from git")
	return nil, nil
}

func Save() error {
	fmt.Println("save to git")
	return nil
}
