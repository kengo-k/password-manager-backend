package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/types"
)

var GetPasswords types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		data := repo.FindPasswords()
		c.PureJSON(http.StatusOK, data)
	}
}
