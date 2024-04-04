package api

import (
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"eduwave-back-end/util"
	"fmt"
	"time"

	_ "eduwave-back-end/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	
	router.Use(cors.New(cors.Config{
		AllowOrigins:   []string{"http://localhost:3001", "https://testnet.bethelnet.io", "http://*", "https://*", "*"},
		AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:   []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/signup", server.createUser)
	router.POST("/login", server.loginUser)
	router.POST("tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker, []string{util.AdminRole, util.TeacherRole, util.StudentRole}))

	//requests
	authRoutes.POST("/requests", server.createRequest)
	authRoutes.GET("/request/:request_id", server.getRequest)
	authRoutes.GET("/requests", server.listRequests)
	authRoutes.DELETE("/request/:student_id/:request_id", server.deleteRequest)
	authRoutes.PUT("/request/:student_id", server.updateRequest)

	//student
	// authRoutes.POST("create/students", server.createStudent)//
	authRoutes.GET("/students", server.listStudents)//
	authRoutes.PUT("/students/:student_id", server.updateStudent)//
	authRoutes.DELETE("/students/:student_id", server.deleteStudent)//

	//teacher
	authRoutes.POST("/teacher", server.createTeacher)
	authRoutes.GET("/teacher/:teacher_id", server.getTeacher)
	authRoutes.GET("/teachers", server.listTeachers)
	authRoutes.PUT("/teachers/:teacher_id", server.updateTeacher)
	authRoutes.DELETE("/teachers/:teacher_id", server.deleteTeacher)

	//admin
	authRoutes.POST("/admins", server.createAdmin)	
	authRoutes.GET("/admins/:admin_id", server.getAdmin)
	authRoutes.PUT("/admins/:admin_id", server.updateAdmin)
	authRoutes.DELETE("/admins/:admin_id", server.deleteAdmin)

	//material
	authRoutes.POST("/material", server.CreateMaterial)	
	authRoutes.GET("/material/:material_id", server.GetMaterials)
	authRoutes.PUT("/material/:material_id", server.UpdateMaterial)
	authRoutes.DELETE("/material/:material_id", server.DeleteMaterial)

	//resource
	authRoutes.POST("/resource", server.createResource)	
	authRoutes.GET("/resource/:resource_id", server.getResource)
	authRoutes.PUT("/resource/:resource_id", server.updateResource)
	authRoutes.DELETE("/resource/:resource_id", server.deleteResource)

	//createcourse
	authRoutes.POST("/course", server.CreateCourse)	
	authRoutes.GET("/course/:course_id", server.GetCourse)
	authRoutes.PUT("/courses", server.ListCourses)
	authRoutes.PUT("/course/:course_id", server.UpdateCourse)
	authRoutes.DELETE("/course/:course_id", server.DeleteCourse)

	//assignment
	authRoutes.POST("/assignments", server.createAssignment)
	authRoutes.GET("/assignments/:assignment_id", server.getAssignment)
	authRoutes.PUT("/assignments/:assignment_id", server.updateAssignment)
	authRoutes.DELETE("/assignments/:assignment_id/:resource_id", server.deleteAssignment)

	//course_enrolments
	authRoutes.GET("/courseEnrolments", server.listEnrolments)

	//course_progress
	authRoutes.POST("/courseProgress", server.createCourseProgress)
	authRoutes.GET("/courseProgress", server.listCourseProgress)
	authRoutes.GET("/courseProgress/:courseProgress_id", server.getCourseProgress)

	//submissions
	// authRoutes.POST("create/submissions/", server.createsubmission)
	authRoutes.GET("/submissions/:submission_id", server.getSubmission)
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
