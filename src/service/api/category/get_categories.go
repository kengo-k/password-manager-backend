package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/types"
)

var GetCategories types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		data := repo.GetCategories()
		c.PureJSON(http.StatusOK, data)
	}
}
