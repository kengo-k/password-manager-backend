package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/types"
)

var Publish types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !repo.IsDirty() {
			c.PureJSON(http.StatusAccepted, map[string]bool{"success": false})
			return
		}
		pwds := repo.Serialize()
		context.Save(pwds)
		repo.SetClean()
		c.PureJSON(http.StatusOK, map[string]bool{"success": true})
	}
}
