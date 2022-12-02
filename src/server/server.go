package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
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
	server.DELETE("/api/passwords/:id", service.DeletePassword)
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

func (service *Service) DeletePassword(c *gin.Context) {
	repo := service.repo
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// TODO return error response (fix in another task)
		panic("failed to convert id to number")
	}
	pwd := repo.GetPassword(id)
	if pwd == nil {
		c.PureJSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("failed to delete, id: `%v` was not found", id),
		})
		return
	}
	repo.DeletePassword(pwd)
	c.PureJSON(http.StatusOK, pwd)
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

	// TODO 変更が一切ないデータが来た場合はcommitしたくないので変更が存在する場合のみSaveする

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

func (service *Service) Publish(c *gin.Context) {
	if !service.repo.IsDirty() {
		c.PureJSON(http.StatusAccepted, map[string]bool{"success": false})
		return
	}
	pwds := service.repo.Serialize()
	context.Save(pwds)
	service.repo.SetClean()
	c.PureJSON(http.StatusOK, map[string]bool{"success": true})
}
