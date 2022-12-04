package password

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/service/response"
	"github.com/kengo-k/password-manager/types"
)

var UpdatePassword types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		// get update target id from path param
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.PureJSON(
				http.StatusBadRequest,
				response.CreateErrorResponse(fmt.Sprintf("failed to update, id: `%v` is not a number", idStr), nil),
			)
			return
		}

		// bind parameters
		var req model.PasswordUpdateRequest
		if c.ShouldBind(&req) != nil {
			c.PureJSON(http.StatusBadRequest, response.CreateErrorResponse("failed to bind bind params", nil))
			return
		}

		// check update target password exists
		pwd := repo.GetPassword(id)
		if pwd == nil {
			c.PureJSON(
				http.StatusBadRequest,
				response.CreateErrorResponse(fmt.Sprintf("failed to update password, id: %v not found", id), nil))
			return
		}

		// validate update params and create updated password
		err = req.Validate(pwd, repo.GetCategories())
		if err != nil {
			c.PureJSON(
				http.StatusBadRequest,
				response.CreateErrorResponse("failed to update password, validation error", &err))
			return
		}

		// TODO 変更が一切ないデータが来た場合はcommitしたくないので変更が存在する場合のみSaveする

		repo.SavePassword(pwd)
		c.PureJSON(http.StatusOK, pwd)
	}
}
