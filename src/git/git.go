package git

import (
	"bufio"
	"fmt"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Git struct{}

// Gitから最新のパスワードファイル(Markdownを取得する)
func (g *Git) Checkout() ([]string, error) {
	f := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), f, &git.CloneOptions{
		URL:           "http://gitbucket.mynet/git/private/password.git",
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
	})
	if err != nil {
		return nil, fmt.Errorf("faield to clone repository: %v", err)
	}
	w, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get work tree: %v", err)
	}
	file, err := w.Filesystem.Open("password.md")
	if err != nil {
		return nil, fmt.Errorf("failed to open password file: %v", err)
	}
	defer file.Close()

	var l []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l = append(l, line)
	}
	return l, nil
}
