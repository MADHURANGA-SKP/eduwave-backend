package api

import (
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"eduwave-back-end/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("user_name", validUsername)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("create/user", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker, []string{util.AdminRole, util.TeacherRole, util.StudentRole}))

	//requests
	authRoutes.POST("create/requests", server.createRequest)
	authRoutes.GET("get/requests/:id", server.getRequest)
	authRoutes.GET("list/requests", server.listRequests)
	authRoutes.DELETE("remove/requests/:id", server.deleteRequest)
	authRoutes.PUT("update/requests/:id", server.updateRequest)

	//student
	// authRoutes.POST("create/students", server.createStudent)//
	authRoutes.GET("list/students", server.listStudents)//
	authRoutes.PUT("update/students/:id", server.updateStudent)//
	authRoutes.DELETE("delete/students/:id", server.deleteStudent)//

	//teacher
	authRoutes.POST("create/teachers", server.createTeacher)
	authRoutes.GET("get/teachers/:teacher_id", server.getTeacher)
	authRoutes.GET("list/teachers", server.listTeachers)
	authRoutes.PUT("/update/teachers/:teacher_id", server.updateTeacher)
	authRoutes.DELETE("delete/teachers/:teacher_id", server.deleteTeacher)

	//admin
	authRoutes.POST("create/admins", server.createAdmin)	
	authRoutes.GET("get/admins/:id", server.getAdmin)
	authRoutes.PUT("update/admins/:id", server.updateAdmin)
	authRoutes.DELETE("delete/admin/:id", server.deleteAdmin)

	//assignment
	authRoutes.POST("create/assignments", server.createAssignment)
	authRoutes.GET("get/assignments/:id", server.getAssignment)
	authRoutes.PUT("update/assignments/:id", server.updateAssignment)
	authRoutes.DELETE("delete/assignments/:id", server.deleteAssignment)

	//course_enrolments
	authRoutes.GET("list/courseEnrolments", server.listEnrolments)

	//course_progress
	authRoutes.GET("list/courseProgress", server.listCourseProgress)
	authRoutes.GET("get/courseProgress/:id", server.getCourseProgress)

	//submissions
	// authRoutes.POST("create/submissions/", server.createsubmission)
	authRoutes.GET("get/submissions/:id", server.getSubmission)
	authRoutes.GET("list/submissions", server.listSubmissions)

	server.router = router
}

// start runs the http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
