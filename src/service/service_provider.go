package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/service/api/category"
	"github.com/kengo-k/password-manager/service/api/password"
	"github.com/kengo-k/password-manager/types"
)

type IServiceProvider interface {
	GetPasswordList() gin.HandlerFunc
	UpdatePassword() gin.HandlerFunc
	CreatePassword() gin.HandlerFunc
	DeletePassword() gin.HandlerFunc
	GetCategoryList() gin.HandlerFunc
	UpdateCategory() gin.HandlerFunc
	Publish() gin.HandlerFunc
}

type ServiceProvider struct {
	repo              *repo.Repository
	context           context.IContext
	createPasswordFn  types.ApiCall
	getPasswordListFn types.ApiCall
	updatePasswordFn  types.ApiCall
	deletePasswordFn  types.ApiCall
	getCategoryListFn types.ApiCall
	updateCategoryFn  types.ApiCall
	publishFn         types.ApiCall
}

func (provider *ServiceProvider) GetPasswordList() gin.HandlerFunc {
	return provider.getPasswordListFn(provider.repo, provider.context)
}
func (provider *ServiceProvider) UpdatePassword() gin.HandlerFunc {
	return provider.updatePasswordFn(provider.repo, provider.context)
}
func (provider *ServiceProvider) CreatePassword() gin.HandlerFunc {
	return provider.createPasswordFn(provider.repo, provider.context)
}
func (provider *ServiceProvider) DeletePassword() gin.HandlerFunc {
	return provider.deletePasswordFn(provider.repo, provider.context)
}
func (provider *ServiceProvider) GetCategoryList() gin.HandlerFunc {
	return provider.getCategoryListFn(provider.repo, provider.context)
}
func (provider *ServiceProvider) UpdateCategory() gin.HandlerFunc {
	return provider.updateCategoryFn(provider.repo, provider.context)
}
func (provider *ServiceProvider) Publish() gin.HandlerFunc {
	return provider.publishFn(provider.repo, provider.context)
}

func NewServiceProvider(context context.IContext) IServiceProvider {
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
	return &ServiceProvider{
		repo: repo,

		createPasswordFn:  password.CreatePassword,
		getPasswordListFn: password.GetPasswords,
		updatePasswordFn:  password.UpdatePassword,
		deletePasswordFn:  password.DeletePassword,
		publishFn:         password.Publish,

		getCategoryListFn: category.GetCategories,
		updateCategoryFn:  category.UpdateCategory,
	}
}
