package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/repo"
)

func setupRouter() *gin.Engine {

	repo, err := repo.NewRepository()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	// パスワードの一覧を返却する
	r.GET("/api/passwords", func(c *gin.Context) {
		data := repo.FindPasswords()
		c.PureJSON(http.StatusOK, data)
	})

	r.GET("/api/categories", func(c *gin.Context) {
		// data := repo.FindCategories()
		// c.PureJSON(http.StatusOK, data)
	})

	r.POST("/api/categories", func(c *gin.Context) {
		// var cat model.Category
		// if c.ShouldBind(&cat) == nil {
		// 	// 入力チェック
		// 	if cat.Name == nil {
		// 		panic("category name is required")
		// 	}
		// 	if cat.ID == nil {
		// 		newID, err := uuid.NewUUID()
		// 		if err != nil {
		// 			panic("failed to create new category id")
		// 		}
		// 		id := newID.String()
		// 		cat.ID = &id
		// 		now := time.Now()
		// 		cat.CreatedAt = &now
		// 		cat.UpdatedAt = &now
		// 		repo.SaveCategory(&cat)
		// 		c.PureJSON(http.StatusOK, map[string]any{})
		// 		return
		// 	}
		// }
		// // 想定外のエラー
		// panic("failed to create category")
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
