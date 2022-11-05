//go:build file2file
// +build file2file

package context

import (
	"fmt"

	"github.com/kengo-k/password-manager/loader"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/saver"
)

func init() {
	fmt.Println("mode: file2file")
}

func Load() ([]string, error) {
	loader := &loader.FileLoader{}
	return loader.Load()
}

func Save(serializedData [][]*model.Password) error {
	saver := &saver.FileSaver{}
	saver.Save(serializedData)
	return nil
}
