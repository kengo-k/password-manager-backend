package loader

import (
	"bufio"
	"fmt"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/kengo-k/password-manager/env"
)

type GitLoader struct{}

// Gitから最新のパスワードファイル(Markdownを取得する)
func (g *GitLoader) Load() ([]string, error) {

	config := env.GetConfig()

	f := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), f, &git.CloneOptions{
		URL:           config.RepositoryURL,
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
	})
	if err != nil {
		return nil, fmt.Errorf("faield to clone repository: %v", err)
	}
	w, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get work tree: %v", err)
	}
	file, err := w.Filesystem.Open(config.PasswordFile)
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

func NewGitLoader() Loader {
	return &GitLoader{}
}
