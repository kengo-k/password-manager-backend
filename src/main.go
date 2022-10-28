package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kengo-k/password-manager/model"
)

type Repository interface {
	Find() []model.Password
	Save()
	Delete()
}

type Database struct {
	db *model.Database
}

func (c *Database) Find() []model.Password {
	return []model.Password{}
}

func (c *Database) Save() {
}

func (c *Database) Delete() {
}

func initRepository() Repository {
	db := &model.Database{}
	conn := &Database{db: db}
	return conn
}

func setupRouter() *gin.Engine {

	repo := initRepository()
	r := gin.Default()

	// パスワードの一覧を返却する
	r.GET("/api/passwords", func(c *gin.Context) {
		data := repo.Find()
		c.PureJSON(http.StatusOK, data)
	})

	r.POST("/api/passwords", func(c *gin.Context) {
		var p model.Password
		if c.ShouldBind(&p) == nil {
			log.Printf("id: %v", p.ID)
			log.Printf("category: %v", p.Category)
			log.Printf("user: %v", p.User)
			log.Printf("password: %v", p.Password)
			log.Printf("note1: %v", p.Note1)
			log.Printf("note2: %v", p.Note2)
			log.Printf("note3: %v", p.Note3)
			// IDが存在しない場合は新規作成
			if p.ID == 0 {
				newID, err := uuid.NewUUID()
				if err != nil {
					panic("failed to create id")
				}
				log.Printf("new id: %v", newID)
			}
		}
		c.PureJSON(http.StatusOK, map[string]any{})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
