package types

import (
	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/repo"
)

type ApiCall func(*repo.Repository, context.IContext) gin.HandlerFunc
