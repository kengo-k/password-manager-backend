package git

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/kengo-k/password-manager/model"
)

type Git struct{}

// Gitから最新のパスワードファイル(Markdownを取得する)
func (g *Git) LoadLatestPassword() ([]string, error) {
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

func (g *Git) Load(mdLines []string) ([]*model.Password, []*model.Category, error) {
	foundCategory := false
	foundHeader := false
	foundSeparator := false
	for _, l := range mdLines {
		// 空行の場合はスキップする
		if len(l) == 0 {
			continue
		}
		// #で始まるコメント行の場合はカテゴリ名が記載されている
		if strings.HasPrefix(l, "#") {
			foundCategory = true
			_, category, ok := strings.Cut(l, "#")
			if !ok {
				panic("failed to get category name")
			}
			fmt.Printf("category: %s\n", category)
			continue
		}
		if foundCategory {
			foundCategory = false
			foundHeader = true
			fmt.Printf("header: %s\n", l)
			continue
		}
		if foundHeader {
			foundHeader = false
			foundSeparator = true
			fmt.Println("separator")
			continue
		}
		if foundSeparator {
			columns := splitColumns(l)
			if len(columns) != 7 {
				panic("invalid column length")
			}
			p := model.Password{
				ID:        &columns[0],
				User:      &columns[1],
				Password:  &columns[2],
				Mail:      &columns[3],
				Note:      &columns[4],
				CreatedAt: nil,
				UpdatedAt: nil,
			}
			fmt.Printf("line: %#v\n", p)
		}
	}

	return nil, nil, nil
}

func splitColumns(line string) []string {
	ret := []string{}
	for _, column := range strings.Split(line, "|") {
		column = strings.TrimSpace(column)
		if len(column) > 0 {
			ret = append(ret, column)
		}
	}
	return ret
}
