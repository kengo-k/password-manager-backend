package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/model"
)

type Connection struct {
	db *model.Database
}

func (c *Connection) Find() []model.Password {
	return []model.Password{}
}

func (c *Connection) Save() {
}

func (c *Connection) Delete() {
}

func initDatabase() *Connection {
	db := &model.Database{}
	conn := &Connection{db: db}
	return conn
}

func setupRouter() *gin.Engine {

	conn := initDatabase()
	r := gin.Default()

	// パスワードの一覧を返却する
	r.GET("/api/passwords", func(c *gin.Context) {
		data := conn.Find()
		c.PureJSON(http.StatusOK, data)
	})

	r.POST("/api/passwords", func(c *gin.Context) {
		conn.Save()
		c.PureJSON(http.StatusOK, map[string]any{})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
