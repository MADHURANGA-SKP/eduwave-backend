package main

import (
	"database/sql"
	"eduwave-back-end/api"
	db "eduwave-back-end/db/sqlc"
	util "eduwave-back-end/util"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*", "*", "https://testnet.bethelnet.io"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	store := db.NewStore(conn)
	runGinServer(config, *store)
}

func runGinServer(config util.Config, store db.Store){
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}