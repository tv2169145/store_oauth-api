package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/store_oauth-api/clients/cassandra"
	"github.com/tv2169145/store_oauth-api/src/domain/access_token"
	"github.com/tv2169145/store_oauth-api/src/http"
	"github.com/tv2169145/store_oauth-api/src/repositories/db"
	"log"
)

var (
	router = gin.Default()
)


func StartApplication() {
	// connect cassandraDB
	// 測試 cassandra 可否正常啟動
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	log.Println("cassandra connection successfully")
	session.Close()

	// make repository
	dbRepository := db.NewRepository()
	//make Service
	atService := access_token.NewService(dbRepository)
	// make handler(router)
	atHandler := http.NewHandler(atService)
	// mapping routers
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	// start app
	router.Run(":8080")
}
