package model

import (
	"fmt"
	"testing"

	"github.com/kengo-k/password-manager/loader"
)

func TestSerialize(t *testing.T) {
	gitLoader := loader.NewGitLoader()
	passwords, err := gitLoader.Load()
	if err != nil {
		t.Errorf("failed to load by git loader: %v", err)
	}
	fmt.Printf("len: %v", len(passwords))
	if len(passwords) == 0 {
		t.Errorf("faild to load by git loader, length=%v", len(passwords))
	}
	database := NewDatabase()
	database.Init(passwords)
	serialized := database.Serialize()
	// tables := [][]int{
	// 	{1, 1},
	// 	{1, 2},
	// }
	categorySize := len(serialized)
	expCatSize := 3
	if categorySize != expCatSize {
		t.Errorf("category size: got=%v, expected=%v", categorySize, expCatSize)
	}
}
