package main

import (
	"database/sql"
	"eduwave-back-end/api"
	db "eduwave-back-end/db/sqlc"
	util "eduwave-back-end/util"

	_ "github.com/lib/pq"

	"log"

	_ "eduwave-back-end/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}                                                               
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db", err)
	}
	
	store := db.NewStore(conn)
	
	runGinServer(config, *store, router)
}

func runGinServer(config util.Config, store db.Store, router *gin.Engine) {
	// Serve Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}


