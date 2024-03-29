package api

import (
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"eduwave-back-end/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// server serves hhtp requests
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a http server and setup routing
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

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	//requests
	authRoutes.POST("/requests", server.createRequest)
	authRoutes.GET("/requests/:id", server.getRequest)
	authRoutes.GET("/requests", server.listRequests)
	authRoutes.DELETE("/requests/:id", server.deleteRequest)
	authRoutes.PUT("/requests/ :id", server.updateRequest)

	//student
	authRoutes.POST("/students", server.createStudent)
	authRoutes.GET("/students", server.listStudents)
	authRoutes.PUT("/students/:id", server.updateStudent)
	authRoutes.DELETE("/students/:id", server.deleteStudent)

	//teacher
	authRoutes.POST("/teachers", server.createTeacher)
	authRoutes.GET("/teachers/:id", server.getTeacher)
	authRoutes.GET("/teachers", server.listTeachers)
	authRoutes.PUT("/teachers/ :id", server.updateTeacher)
	authRoutes.DELETE("/teachers/ :id", server.deleteTeacher)

	//admin
	authRoutes.GET("/admins/:id", server.getAdmin)
	authRoutes.PUT("/admins/ :id", server.updateAdmin)
	authRoutes.DELETE("/admin/ :id", server.deleteAdmin)

	//assignment
	authRoutes.POST("/assignments", server.createAssignment)
	authRoutes.GET("/assignments/:id", server.getAssignment)
	authRoutes.PUT("/assignments/ :id", server.updateAssignment)
	authRoutes.DELETE("/assignments/ :id", server.deleteAssignment)

	//course_enrolments
	authRoutes.GET("/courseEnrolments", server.listEnrolments)

	//course_progress
	authRoutes.GET("/courseProgress", server.listCourseProgress)
	authRoutes.GET("/courseProgress/:id", server.getCourseProgress)

	//submissions
	authRoutes.GET("/assignments/:id", server.getSubmission)
	authRoutes.GET("/submissions", server.listSubmissions)

	server.router = router
}

// start runs the http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
