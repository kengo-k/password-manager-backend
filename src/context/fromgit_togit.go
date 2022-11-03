//go:build !fromfile_tofile && ignore && !fromgit_tofile
// +build !fromfile_tofile,ignore,!fromgit_tofile

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
