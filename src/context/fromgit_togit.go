//go:build !file2file && !git2file
// +build !file2file,!git2file

package context

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/kengo-k/password-manager/env"
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

	// save password database into file
	saver := &saver.FileSaver{}
	saver.Save(serializedData)

	// push saved file to git repository
	config := env.GetConfig()
	pwdFile, err := os.Open(config.PasswordFile)

	// 読み取り時の例外処理
	if err != nil {
		panic("failed to open password file")
	}
	defer pwdFile.Close()

	contents, err := io.ReadAll(pwdFile)
	if err != nil {
		panic("failed to read password file")
	}

	f := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), f, &git.CloneOptions{
		URL:           config.RepositoryURL,
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
	})
	if err != nil {
		// TODO return error response
		panic(fmt.Sprintf("failed to clone: %v", err))
	}
	w, err := repo.Worktree()
	if err != nil {
		// TODO return error response
		panic("failed to get work tree")
	}
	file, err := w.Filesystem.Create(config.PasswordFile)
	if err != nil {
		panic("failed to create new file in file system")
	}
	defer file.Close()
	file.Write(contents)
	w.Add(config.PasswordFile)
	w.Commit("commit by password manager", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "password-manager",
			Email: "test@example.com",
			When:  time.Now(),
		},
	})

	auth := &http.BasicAuth{
		Username: config.RepositoryUser,
		Password: config.RepositoryPass,
	}
	err = repo.Push(&git.PushOptions{Auth: auth})
	if err != nil {
		panic(fmt.Sprintf("failed to push: %v", err))
	}
	return nil
}
