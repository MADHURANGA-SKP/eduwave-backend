package api

import (
	"fmt"
    db "eduwave-back-end/db/sqlc"

	"eduwave-back-end/token"

	util "eduwave-back-end/util"

	"github.com/gin-gonic/gin"
)

//server serves hhtp requests
type Server struct {
    config     util.Config
    store      db.Store
    tokenMaker token.Maker
    router     *gin.Engine
}

//NewServer creates a http server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
    tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
    if err != nil {
        return nil, fmt.Errorf("cannot create token maker: %w", err)
    }

    server := &Server{
        config:     config,
        store:      store,
        tokenMaker: tokenMaker,
    }
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	server.router = router
}

//start runs the http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}