package app

import (
	"github.com/gin-gonic/gin"
	"github.com/suvamsingh/bookstore_oauth-api/src/domain/access_token"
	"github.com/suvamsingh/bookstore_oauth-api/src/http"
	"github.com/suvamsingh/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.Run(":8080")

}
