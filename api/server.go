package api

import (
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"eduwave-back-end/util"
	"fmt"

	_ "eduwave-back-end/docs"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("user/signup", server.createUser)
	router.POST("user/login", server.loginUser)
	router.POST("tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker, []string{util.AdminRole, util.TeacherRole, util.StudentRole}))

	//requests
	authRoutes.POST("create/requests", server.createRequest)
	authRoutes.GET("get/request/:request_id", server.getRequest)
	authRoutes.GET("list/requests", server.listRequests)
	authRoutes.DELETE("delete/request/:request_id", server.deleteRequest)
	authRoutes.PUT("update/request/:student_id", server.updateRequest)

	//student
	// authRoutes.POST("create/students", server.createStudent)//
	authRoutes.GET("list/students", server.listStudents)//
	authRoutes.PUT("update/students/:student_id", server.updateStudent)//
	authRoutes.DELETE("delete/students/:student_id", server.deleteStudent)//

	//teacher
	authRoutes.POST("create/teacher", server.createTeacher)
	authRoutes.GET("get/teacher/:teacher_id", server.getTeacher)
	authRoutes.GET("list/teachers", server.listTeachers)
	authRoutes.PUT("update/teachers/:teacher_id", server.updateTeacher)
	authRoutes.DELETE("delete/teachers/:teacher_id", server.deleteTeacher)

	//admin
	authRoutes.POST("create/admins", server.createAdmin)	
	authRoutes.GET("get/admins/:admin_id", server.getAdmin)
	authRoutes.PUT("update/admins/:admin_id", server.updateAdmin)
	authRoutes.DELETE("delete/admins/:admin_id", server.deleteAdmin)

	//material
	authRoutes.POST("create/material", server.CreateMaterial)	
	authRoutes.GET("get/material/:material_id", server.GetMaterials)
	authRoutes.PUT("update/material/:material_id", server.UpdateMaterial)
	authRoutes.DELETE("delete/material/:material_id", server.DeleteMaterial)

	//resource
	authRoutes.POST("create/resource", server.createResource)	
	authRoutes.GET("get/resource/:resource_id", server.getResource)
	authRoutes.PUT("update/resource/:resource_id", server.updateResource)
	authRoutes.DELETE("delete/resource/:resource_id", server.deleteResource)

	//createcourse
	authRoutes.POST("create/course", server.CreateCourse)	
	authRoutes.GET("get/course/:course_id", server.GetCourse)
	authRoutes.PUT("list/courses", server.ListCourses)
	authRoutes.PUT("update/course/:course_id", server.UpdateCourse)
	authRoutes.DELETE("delete/course/:course_id", server.DeleteCourse)

	//assignment
	authRoutes.POST("create/assignments", server.createAssignment)
	authRoutes.GET("get/assignments/:assignment_id", server.getAssignment)
	authRoutes.PUT("update/assignments/:assignment_id", server.updateAssignment)
	authRoutes.DELETE("delete/assignments/:assignment_id", server.deleteAssignment)

	//course_enrolments
	authRoutes.GET("list/courseEnrolments", server.listEnrolments)

	//course_progress
	authRoutes.POST("create/courseProgress", server.createCourseProgress)
	authRoutes.GET("list/courseProgress", server.listCourseProgress)
	authRoutes.GET("get/courseProgress/:courseProgress_id", server.getCourseProgress)

	//submissions
	// authRoutes.POST("create/submissions/", server.createsubmission)
	authRoutes.GET("get/submissions/:submission_id", server.getSubmission)
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
