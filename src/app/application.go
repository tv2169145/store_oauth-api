package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/store_oauth-api/src/http"
	"github.com/tv2169145/store_oauth-api/src/repositories/db"
	"github.com/tv2169145/store_oauth-api/src/repositories/rest"
	"github.com/tv2169145/store_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)


func StartApplication() {
	// connect cassandraDB
	// 測試 cassandra 可否正常啟動
	//session := cassandra.GetSession()
	//log.Println("cassandra connection successfully")
	//session.Close()

	// make repository
	dbRepository := db.NewRepository()
	//make usersRepository
	userRepository := rest.NewRepository()
	//make Service
	atService := access_token.NewService(userRepository, dbRepository)
	// make handler(router)
	atHandler := http.NewHandler(atService)
	// mapping routers
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	// start app
	router.Run(":8080")
}
