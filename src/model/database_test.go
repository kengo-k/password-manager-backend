package model

import (
	"testing"

	"github.com/kengo-k/password-manager/loader"
)

func TestSerialize(t *testing.T) {
	gitLoader := loader.NewGitLoader()
	passwords, err := gitLoader.Load()
	if err != nil {
		t.Errorf("failed to load by git loader: %v", err)
	}
	if len(passwords) == 0 {
		t.Errorf("faild to load by git loader, length=%v", len(passwords))
	}
	database := NewDatabase()
	database.Init(passwords)
	serialized := database.Serialize()
	categorySize := len(serialized)
	expCatSize := 5
	if categorySize != expCatSize {
		t.Errorf("category size: got=%v, expected=%v", categorySize, expCatSize)
	}
	for i, passwords := range serialized {
		if len(passwords) == 0 {
			t.Errorf("len[%v] is 0", i)
		}
	}
}
