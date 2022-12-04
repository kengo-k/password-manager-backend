package password

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/service/response"
	"github.com/kengo-k/password-manager/types"
)

var DeletePassword types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		// get delete target id from path param
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.PureJSON(
				http.StatusBadRequest,
				response.CreateErrorResponse(fmt.Sprintf("failed to delete, id: `%v` is not a number", idStr), nil),
			)
			return
		}

		// if password exists, delete that password
		pwd := repo.GetPassword(id)
		if pwd == nil {
			c.PureJSON(
				http.StatusNotFound,
				response.CreateErrorResponse(fmt.Sprintf("failed to delete, id: `%v` was not found", id), nil),
			)
			return
		}
		repo.DeletePassword(pwd)

		c.PureJSON(http.StatusOK, pwd)
	}
}
