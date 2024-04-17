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
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
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
		AllowOrigins:   []string{"http://localhost:3000", "https://testnet.bethelnet.io", "http://*", "https://*", "*"},
		AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//public routes
	router.POST("/signup", server.createUser)
	router.POST("/login", server.loginUser)
	router.POST("tokens/renew_access", server.renewAccessToken)

	//verify email
	router.POST("/verify-email", server.VerifyEmailHandler)

	//RBAC auth routes
		authroute := router.Group("/").Use(authMiddleware(server.tokenMaker))
		
		authroute.POST("/admin/signup", server.createAdminUser)
		authroute.PUT("/user/edit", server.UpdateUser)
		authroute.GET("/listadmin", server.ListUser)
		authroute.GET("/liststudent", server.ListUserStudent)
		authroute.GET("/listteacher", server.ListUserTeacher)
		//requests
			authroute.POST("/requests", server.createRequest)//
			authroute.GET("/request/get", server.getRequest)
			authroute.GET("/requests", server.listRequests)
			authroute.DELETE("/request/delete", server.deleteRequest)
			authroute.PUT("/request/edit", server.updateRequest)
		//material
			authroute.POST("/material", server.CreateMaterial)	
			authroute.GET("/material/get", server.GetMaterials)
			authroute.PUT("/material/edit", server.UpdateMaterial)
			authroute.DELETE("/material/delete", server.DeleteMaterial)
		//resource
			authroute.POST("/resource", server.createResource)	
			authroute.GET("/resource/get", server.getResource)
			authroute.PUT("/resource/edit", server.updateResource)
			authroute.DELETE("/resource/delete", server.deleteResource)
		//createcourse
			authroute.POST("/course", server.CreateCourse)	
			authroute.GET("/course/get", server.GetCourse)
			authroute.GET("/courses", server.ListCourses)
			authroute.PUT("/course/edit", server.UpdateCourses)
			authroute.DELETE("/course/delete", server.DeleteCourse)
		//assignment
			authroute.POST("/assignments", server.createAssignment)
			authroute.GET("/assignment/get", server.getAssignment)
			authroute.PUT("/assignments/edit", server.updateAssignment)
			authroute.DELETE("/assignment/delete", server.deleteAssignment)
		//submissions
			authroute.POST("/submissions", server.CreateSubmission)
			authroute.GET("/submission/byassignment", server.GetSubmissionsByAssignment)
			authroute.GET("/submission/byuser", server.GetSubmissionsByUser)
			authroute.GET("/submissions", server.listSubmissions)
		//course_enrolments
			authroute.POST("/enrol", server.CreateCourseEnrolment)
			authroute.GET("/enrolments", server.listEnrolments)
		//course_progress
			authroute.POST("/createprogress", server.createCourseProgress)
			authroute.GET("/courseprogress", server.listCourseProgress)
			authroute.GET("/courseprogress/get", server.getCourseProgress)

	server.router = router
}

// start runs the http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}