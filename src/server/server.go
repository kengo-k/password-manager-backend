package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
)

type Service struct {
	repo *repo.Repository
}

func NewServer(service *Service) *gin.Engine {
	server := gin.Default()
	server.GET("/api/passwords", service.GetPasswordList)
	server.PUT("/api/passwords/:id", service.UpdatePassword)
	server.POST("/api/passwords", service.CreatePassword)
	server.GET("/api/categories", service.GetCategoryList)
	server.PUT("/api/categories/:id", service.UpdateCategory)
	server.POST("/api/passwords/publish", service.Publish)
	return server
}

func NewService() *Service {
	// load data
	passwords, err := context.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load initial data: %v", err))
	}

	// init database
	database := model.NewDatabase()
	if err := database.Init(passwords); err != nil {
		panic(fmt.Sprintf("failed to init database: %v", err))
	}

	repo := repo.NewRepository(database)
	return &Service{repo: repo}
}

func (service *Service) GetPasswordList(c *gin.Context) {
	data := service.repo.FindPasswords()
	c.PureJSON(http.StatusOK, data)
}

func (service *Service) UpdatePassword(c *gin.Context) {
	repo := service.repo
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// TODO return error response (fix in another task)
		panic("failed to convert id to number")
	}

	var req model.PasswordUpdateRequest
	if c.ShouldBind(&req) != nil {
		// TODO return error response (fix in another task)
		panic("failed to bind update params")
	}
	pwd := repo.GetPassword(id)
	if pwd == nil {
		// TODO return error response (fix in another task)
		panic("failed to get password")
	}

	err = req.Validate(pwd, repo.GetCategories())
	if err != nil {
		// TODO return error response (fix in another task)
		panic("failed to validate update request")
	}
	repo.SavePassword(pwd)
	c.PureJSON(http.StatusOK, pwd)
}

func (service *Service) GetCategoryList(c *gin.Context) {
	data := service.repo.GetCategories()
	c.PureJSON(http.StatusOK, data)
}

func (service *Service) UpdateCategory(c *gin.Context) {
	repo := service.repo
	catID := c.Param("id")
	var req model.CategoryUpdateRequest
	if c.ShouldBind(&req) == nil {
		cat := repo.GetCategory(catID)
		if cat != nil {
			if req.Name != nil {
				cat.Name = *req.Name
			}
			if req.Order != nil {
				cat.Order = *req.Order
			}
			repo.SaveCategory(cat)
		}
		c.PureJSON(http.StatusOK, cat)
	}
}

func (service *Service) CreatePassword(c *gin.Context) {
	repo := service.repo

	var req model.PasswordCreateRequest
	if c.ShouldBind(&req) != nil {
		// TODO return error response (fix in another task)
		panic("failed to bind create params")
	}

	pwd, err := req.Validate(repo.GetCategories())
	if err != nil {
		// TODO return error response (fix in another task)
		panic("failed to validate create params")
	}

	pwd.ID = repo.GetNextPasswordId()
	repo.SavePassword(pwd)
	c.PureJSON(http.StatusOK, pwd)
}

func (service *Service) Publish(c *gin.Context) {

	config := env.GetConfig()

	pwds := service.repo.Serialize()
	context.Save(pwds)

	pwdFile, err := os.Open(config.PasswordFile)
	// 読み取り時の例外処理
	if err != nil {
		panic("failed to open password file")
	}
	defer pwdFile.Close()

	contents, err := ioutil.ReadAll(pwdFile)
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

	auth := &githttp.BasicAuth{
		Username: config.RepositoryUser,
		Password: config.RepositoryPass,
	}
	err = repo.Push(&git.PushOptions{Auth: auth})
	if err != nil {
		panic(fmt.Sprintf("failed to push: %v", err))
	}
	c.PureJSON(http.StatusOK, map[string]bool{"success": true})
}
