package repo

import (
	"testing"

	"github.com/kengo-k/password-manager/git"
)

func TestInitRepository(t *testing.T) {
	g := &git.Git{}
	r := newRepositoryImpl()
	passwords, err := g.Checkout()
	if err != nil {
		t.Errorf("failed to checkout passwords: %v", err)
	}
	err = r.Init(passwords)
	if err != nil {
		t.Errorf("failed to init repository: %v", err)
	}
}
