package context

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/model"
)

type IContext interface {
	Load() ([]string, error)
	Save(serializedData [][]*model.Password) error
}

type Context struct {
	LoadFn func() ([]string, error)
	SaveFn func(serializeData [][]*model.Password) error
}

func (ctx *Context) Load() ([]string, error) {
	return ctx.LoadFn()
}

func (ctx *Context) Save(serializeData [][]*model.Password) error {
	return ctx.SaveFn(serializeData)
}

func loadFile() ([]string, error) {
	f, err := os.Open("./password.md")
	if err != nil {
		return nil, fmt.Errorf("failed to open from file")
	}
	defer f.Close()

	lines := make([]string, 1024)
	reader := bufio.NewReaderSize(f, 1024)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read from file")
		}
		lines = append(lines, string(line))
	}

	return lines, nil
}

func loadRepository() ([]string, error) {

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

func saveFile(serializedData [][]*model.Password) error {
	config := env.GetConfig()
	var sb strings.Builder
	ifNil := func(sp *string) string {
		if sp == nil {
			return ""
		} else {
			return *sp
		}
	}

	catCount := len(serializedData)
	for i, passwords := range serializedData {
		head := passwords[0]
		fmt.Fprintf(&sb, "# %s: name=%s,order=%d\n", head.Category.ID, head.Category.Name, head.Category.Order)
		fmt.Fprint(&sb, "| 名称 | 説明 | ユーザ | パスワード | メール | 備考 |\n")
		fmt.Fprint(&sb, "|------|------|--------|------------|--------|------|\n")
		for _, p := range passwords {
			fmt.Fprintf(&sb, "| %s | %s | %s | %s | %s | %s |\n",
				p.Name, ifNil(p.Desc), ifNil(p.User), ifNil(p.Password), ifNil(p.Mail), ifNil(p.Note))
		}
		if i < catCount-1 {
			fmt.Fprint(&sb, "\n")
		}
	}
	f, err := os.Create(config.PasswordFile)
	if err != nil {
		panic("failed to open file for write")
	}
	defer f.Close()
	f.Write([]byte(sb.String()))
	return nil
}

func saveRepository(serializedData [][]*model.Password) error {

	// save password database into file
	saveFile(serializedData)

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

func NewContext(mode runmode.RunMode) IContext {
	ctx := &Context{}
	switch mode {
	case runmode.FILE_TO_FILE:
		ctx.LoadFn = loadFile
		ctx.SaveFn = saveFile
	case runmode.GIT_TO_FILE:
		ctx.LoadFn = loadRepository
		ctx.SaveFn = saveFile
	case runmode.GIT_TO_GIT:
		ctx.LoadFn = loadRepository
		ctx.SaveFn = saveRepository
	}
	return ctx
}
