package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/types"
)

var UpdateCategory types.ApiCall = func(repo *repo.Repository, context context.IContext) gin.HandlerFunc {

	return func(c *gin.Context) {
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
}
