package password

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/service/response"
	"github.com/kengo-k/password-manager/types"
	"github.com/kengo-k/password-manager/validators"
)

var CreatePassword types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		// bind parameters
		var req model.PasswordCreateRequest
		if c.ShouldBind(&req) != nil {
			c.PureJSON(http.StatusBadRequest, response.CreateErrorResponse("failed to bind create params", nil))
			return
		}

		// validate parameters
		categories := repo.GetCategories()
		validate := validator.New()
		validate.RegisterValidation("is_valid_category", validators.ValidateCategory(categories))
		err := validate.Struct(&req)
		if err != nil {
			c.PureJSON(http.StatusBadRequest, response.CreateErrorResponse("failed to validate create params", &err))
			return
		}

		// TODO ライブラリのバリデーションを採用したので使うのやめる
		pwd, err := req.Validate(repo.GetCategories())
		if err != nil {
			// TODO return error response (fix in another task)
			panic(fmt.Sprintf("failed to validate create params: %v", err))
		}

		pwd.ID = repo.GetNextPasswordId()
		repo.SavePassword(pwd)
		c.PureJSON(http.StatusOK, pwd)
	}
}
