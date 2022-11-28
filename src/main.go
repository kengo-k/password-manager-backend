package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kengo-k/password-manager/context"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
)

func setupRouter() *gin.Engine {

	passwords, err := context.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load initial data: %v", err))
	}
	database := model.NewDatabase()
	database.Init(passwords)

	// TODO 暫定処理 後で消す
	serializedData := database.Serialize()
	context.Save(serializedData)

	repo := repo.NewRepository(database)

	r := gin.Default()

	// パスワードの一覧を返却する
	r.GET("/api/passwords", func(c *gin.Context) {
		data := repo.FindPasswords()
		c.PureJSON(http.StatusOK, data)
	})

	r.PUT("/api/passwords/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			// TODO return error response (fix in another task)
			panic("failed to convert id to number")
		}

		var req model.PasswordUpdateRequest
		if c.ShouldBind(&req) != nil {
			// TODO return error response (fix in another task)
			panic("failed to bind update params")
		}
		pwd := repo.GetPassword(id)
		if pwd == nil {
			// TODO return error response (fix in another task)
			panic("failed to get password")
		}

		// if pwd's category_id !=  req's category_id,
		// apply category change
		if req.CategoryID != nil && pwd.Category.ID != *req.CategoryID {
			newCat := repo.GetCategory(*req.CategoryID)
			if newCat == nil {
				// TODO return error response (fix in another task)
				panic("failed to get category")
			}
			pwd.Category = newCat
		}
		pwd.ApplyUpdateValues(&req)
		repo.SavePassword(pwd)
		c.PureJSON(http.StatusOK, pwd)
	})

	r.GET("/api/categories", func(c *gin.Context) {
		data := repo.FindCategories()
		c.PureJSON(http.StatusOK, data)
	})

	r.PUT("/api/categories/:id", func(c *gin.Context) {
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
	})

	r.POST("/api/passwords", func(c *gin.Context) {
		// var p model.Password
		// if c.ShouldBind(&p) == nil {
		// 	log.Printf("id: %v", p.ID)
		// 	log.Printf("category: %v", p.Category)
		// 	log.Printf("user: %v", p.User)
		// 	log.Printf("password: %v", p.Password)
		// 	log.Printf("note: %v", p.Note)
		// 	// 入力チェック
		// 	// カテゴリが未設定の場合はエラーで終了する
		// 	if p.Category == nil {
		// 		panic("category is required")
		// 	}
		// 	// 設定されているカテゴリーが存在しない場合はエラーで終了する

		// 	// IDが存在しない場合は新規作成
		// 	if p.ID == nil {
		// 		newID, err := uuid.NewUUID()
		// 		if err != nil {
		// 			panic("failed to create id")
		// 		}
		// 		log.Printf("new id: %v", newID)
		// 		id := newID.String()
		// 		p.ID = &id
		// 		repo.SavePassword(&p)
		// 		c.PureJSON(http.StatusOK, map[string]any{})
		// 		return
		// 	}
		// }
		// // その他想定外の理由でPassword生成できなかった場合はエラーで終了する
		// panic("failed to created password")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
