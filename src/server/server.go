package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/service"
)

func NewServer(service service.IServiceProvider) *gin.Engine {
	server := gin.Default()

	server.POST("/api/passwords", service.CreatePassword())
	server.GET("/api/passwords", service.GetPasswordList())
	server.PUT("/api/passwords/:id", service.UpdatePassword())
	server.DELETE("/api/passwords/:id", service.DeletePassword())
	server.POST("/api/passwords/publish", service.Publish())

	server.GET("/api/categories", service.GetCategoryList())
	server.PUT("/api/categories/:id", service.UpdateCategory())

	return server
}

// func (service *Service) GetPasswordList(c *gin.Context) {
// 	data := service.repo.FindPasswords()
// 	c.PureJSON(http.StatusOK, data)
// }

// func (service *Service) CreatePassword(c *gin.Context) {
// 	repo := service.repo

// 	// bind parameters
// 	var req model.PasswordCreateRequest
// 	if c.ShouldBind(&req) != nil {
// 		c.PureJSON(http.StatusBadRequest, createErrorResult("failed to bind create params", nil))
// 		return
// 	}

// 	// validate parameters
// 	categories := repo.GetCategories()
// 	validate := validator.New()
// 	validate.RegisterValidation("is_valid_category", validators.ValidateCategory(categories))
// 	err := validate.Struct(&req)
// 	if err != nil {
// 		c.PureJSON(http.StatusBadRequest, createErrorResult("failed to validate create params", &err))
// 		return
// 	}

// 	// TODO ライブラリのバリデーションを採用したので使うのやめる
// 	pwd, err := req.Validate(repo.GetCategories())
// 	if err != nil {
// 		// TODO return error response (fix in another task)
// 		panic(fmt.Sprintf("failed to validate create params: %v", err))
// 	}

// 	pwd.ID = repo.GetNextPasswordId()
// 	repo.SavePassword(pwd)
// 	c.PureJSON(http.StatusOK, pwd)
// }

// func (service *Service) UpdatePassword(c *gin.Context) {
// 	repo := service.repo

// 	// get update target id from path param
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.PureJSON(
// 			http.StatusBadRequest,
// 			createErrorResult(fmt.Sprintf("failed to update, id: `%v` is not a number", idStr), nil),
// 		)
// 		return
// 	}

// 	// bind parameters
// 	var req model.PasswordUpdateRequest
// 	if c.ShouldBind(&req) != nil {
// 		c.PureJSON(http.StatusBadRequest, createErrorResult("failed to bind bind params", nil))
// 		return
// 	}

// 	// check update target password exists
// 	pwd := repo.GetPassword(id)
// 	if pwd == nil {
// 		c.PureJSON(
// 			http.StatusBadRequest,
// 			createErrorResult(fmt.Sprintf("failed to update password, id: %v not found", id), nil))
// 		return
// 	}

// 	// validate update params and create updated password
// 	err = req.Validate(pwd, repo.GetCategories())
// 	if err != nil {
// 		c.PureJSON(
// 			http.StatusBadRequest,
// 			createErrorResult("failed to update password, validation error", &err))
// 		return
// 	}

// 	// TODO 変更が一切ないデータが来た場合はcommitしたくないので変更が存在する場合のみSaveする

// 	repo.SavePassword(pwd)
// 	c.PureJSON(http.StatusOK, pwd)
// }

// func (service *Service) DeletePassword(c *gin.Context) {
// 	repo := service.repo

// 	// get delete target id from path param
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.PureJSON(
// 			http.StatusBadRequest,
// 			createErrorResult(fmt.Sprintf("failed to delete, id: `%v` is not a number", idStr), nil),
// 		)
// 		return
// 	}

// 	// if password exists, delete that password
// 	pwd := repo.GetPassword(id)
// 	if pwd == nil {
// 		c.PureJSON(
// 			http.StatusNotFound,
// 			createErrorResult(fmt.Sprintf("failed to delete, id: `%v` was not found", id), nil),
// 		)
// 		return
// 	}
// 	repo.DeletePassword(pwd)

// 	c.PureJSON(http.StatusOK, pwd)
// }

// func (service *Service) GetCategoryList(c *gin.Context) {
// 	data := service.repo.GetCategories()
// 	c.PureJSON(http.StatusOK, data)
// }

// func (service *Service) UpdateCategory(c *gin.Context) {
// 	repo := service.repo
// 	catID := c.Param("id")
// 	var req model.CategoryUpdateRequest
// 	if c.ShouldBind(&req) == nil {
// 		cat := repo.GetCategory(catID)
// 		if cat != nil {
// 			if req.Name != nil {
// 				cat.Name = *req.Name
// 			}
// 			if req.Order != nil {
// 				cat.Order = *req.Order
// 			}
// 			repo.SaveCategory(cat)
// 		}
// 		c.PureJSON(http.StatusOK, cat)
// 	}
// }

// func (service *Service) Publish(c *gin.Context) {
// 	if !service.repo.IsDirty() {
// 		c.PureJSON(http.StatusAccepted, map[string]bool{"success": false})
// 		return
// 	}
// 	pwds := service.repo.Serialize()
// 	context.Save(pwds)
// 	service.repo.SetClean()
// 	c.PureJSON(http.StatusOK, map[string]bool{"success": true})
// }
