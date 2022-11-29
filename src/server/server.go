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
	server.GET("/api/categories", service.GetCategoryList)
	server.PUT("/api/categories/:id", service.UpdateCategory)
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

	// TODO 暫定処理 後で消す
	serializedData := database.Serialize()
	context.Save(serializedData)

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

	// if pwd's category_id !=  req's category_id,
	// apply category change
	if req.CategoryID != nil && pwd.Category.ID != *req.CategoryID {
		newCat := repo.GetCategory(*req.CategoryID)
		if newCat == nil {
			// TODO return error response (fix in another task)
			panic("failed to get category")
		}
		pwd.Category = newCat
	}
	req.ApplyValuesWithoutCategory(pwd)
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
