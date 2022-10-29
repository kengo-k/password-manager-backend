package git

import (
	"fmt"
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
	w, err := repo.Worktree()
	if err != nil {
		t.Errorf("failed to get work tree")
	}
	fileList, err := w.Filesystem.ReadDir(".")
	if err != nil {
		t.Errorf("failed to read dir")
	}
	for _, f := range fileList {
		fmt.Printf("filename: %s\n", f.Name())
	}
}
