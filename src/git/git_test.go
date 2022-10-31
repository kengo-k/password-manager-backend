package git

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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
	file, err := w.Filesystem.Create("test.txt")
	if err != nil {
		t.Errorf("failed to create new file in file system")
	}
	file.Write([]byte("test"))
	defer file.Close()

	w.Add("test.txt")
	w.Commit("test commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "password-manager",
			Email: "test@example.com",
			When:  time.Now(),
		},
	})

	auth := &http.BasicAuth{
		Username: "root",
		Password: "root",
	}
	err = repo.Push(&git.PushOptions{Auth: auth})
	if err != nil {
		t.Errorf("failed to push: %v", err)
	}
}

func TestLoadPassword(t *testing.T) {
	g := &Git{}
	list, err := g.Checkout()
	if err != nil {
		t.Errorf("failed to load password: %v", err)
	}
	len := len(list)
	if !(len > 0) {
		t.Errorf("contents is empty")
	}
}

func TestLoad(t *testing.T) {
	// g := &Git{}
	// mdLines := []string{
	// 	"# category1:  category 1 description",
	// 	"| id | user | password | mail | note | created_at | updated_at |",
	// 	"|----|------|----------|------|------|------------|------------|",
	// 	"| id1| user1| password1| mail1| note1| created_at1| updated_at1|",
	// 	"| id2| user2| password2| mail2| note2| created_at2| updated_at2|",
	// 	"",
	// 	"# category2:  category 2 description",
	// 	"| id | user | password | mail | note | created_at | updated_at |",
	// 	"|----|------|----------|------|------|------------|------------|",
	// 	"| id3| user3| password3| mail3| note3| created_at3| updated_at3|",
	// 	"| id4| user4| password4| mail4| note4| created_at4| updated_at4|",
	// 	"",
	// }
	// g.Load(mdLines)
}
