package main

import (
	"database/sql"
	"eduwave-back-end/api"
	db "eduwave-back-end/db/sqlc"
	util "eduwave-back-end/util"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/lib/pq"

	_ "eduwave-back-end/docs"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
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
	
	rundDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	runGinServer(config, *store, router)
}

//run db migration when server start
func rundDBMigration(MigrationURL string, DBSource string){
	//create new migrate instance
	migration, err := migrate.New(MigrationURL,DBSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance ", err)
	}
	//run migrate up
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange  {
		log.Fatal("failed to run migrate up", err)
	}

	log.Println("db migrated succesfully")
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


