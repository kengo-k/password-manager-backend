//go:build git2file
// +build git2file

package context

import (
	"fmt"

	"github.com/kengo-k/password-manager/loader"
)

func init() {
	fmt.Println("mode: git2file")
}

func Load() ([]string, error) {
	g := &loader.GitLoader{}
	passwords, err := g.Load()
	if err != nil {
		return nil, err
	}
	return passwords, nil
}

func Save() error {
	fmt.Println("save to file")
	return nil
}
