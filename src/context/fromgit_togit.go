//go:build !file2file && !git2file
// +build !file2file,!git2file

package context

import (
	"fmt"

	"github.com/kengo-k/password-manager/loader"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/saver"
)

func init() {
	fmt.Println("mode: git2git")
}

func Load() ([]string, error) {
	g := &loader.GitLoader{}
	passwords, err := g.Load()
	if err != nil {
		return nil, err
	}
	return passwords, nil
}

func Save(serializedData [][]*model.Password) error {
	saver := &saver.FileSaver{}
	saver.Save(serializedData)
	return nil
}
