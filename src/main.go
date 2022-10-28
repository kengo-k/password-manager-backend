package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		data := map[string]any{
			"message": "Hello,Gin!",
			"success": true,
			"values":  []int{1, 2, 3, 4, 5},
		}
		c.PureJSON(http.StatusOK, data)
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
