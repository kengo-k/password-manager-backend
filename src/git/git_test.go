package git

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func TestClone(t *testing.T) {
	f := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), f, &git.CloneOptions{
		URL:           "http://gitbucket.mynet/git/private/password.git",
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
	})
	if err != nil {
		t.Errorf("failed to clone: %v", err)
	}
	_, err = repo.Worktree()
	if err != nil {
		t.Errorf("failed to get work tree")
	}
}
