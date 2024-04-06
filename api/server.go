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
		AllowOrigins:   []string{"http://localhost:3001", "https://testnet.bethelnet.io", "http://*", "https://*", "*"},
		AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:   []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//public routes
	router.POST("/signup", server.createUser)
	router.POST("/login", server.loginUser)
	router.POST("tokens/renew_access", server.renewAccessToken)


	//ADMIN - RBAC
		authroute := router.Group("/").Use(authMiddleware(server.tokenMaker, []string{util.AdminRole, util.TeacherRole, util.StudentRole}))
		
		router.POST("/admin/signup", server.createAdminUser)
		//requests
			authroute.POST("/requests", server.createRequest)//
			authroute.GET("/request/:request_id", server.getRequest)
			authroute.GET("/requests", server.listRequests)
			authroute.DELETE("/request/:student_id/:request_id", server.deleteRequest)
			authroute.PUT("/request/:student_id", server.updateRequest)
		//material
			authroute.POST("/material", server.CreateMaterial)	
			authroute.GET("/material/:material_id", server.GetMaterials)
			authroute.PUT("/material/:material_id", server.UpdateMaterial)
			authroute.DELETE("/material/:material_id", server.DeleteMaterial)
		//resource
			authroute.POST("/resource", server.createResource)	
			authroute.GET("/resource/:resource_id", server.getResource)
			authroute.PUT("/resource/:resource_id", server.updateResource)
			authroute.DELETE("/resource/:resource_id", server.deleteResource)
		//createcourse
			authroute.POST("/course", server.CreateCourse)	
			authroute.GET("/course/:course_id", server.GetCourse)
			authroute.PUT("/courses", server.ListCourses)
			authroute.PUT("/course/:course_id", server.UpdateCourse)
			authroute.DELETE("/course/:course_id", server.DeleteCourse)
		//assignment
			authroute.POST("/assignments", server.createAssignment)
			authroute.GET("/assignments/:assignment_id", server.getAssignment)
			authroute.PUT("/assignments/:assignment_id", server.updateAssignment)
			authroute.DELETE("/assignments/:assignment_id/:resource_id", server.deleteAssignment)
		//submissions
			// authRoutes.POST("create/submissions/", server.createsubmission)
			authroute.GET("/submissions/:submission_id", server.getSubmission)
			authroute.GET("/submissions", server.listSubmissions)
		//course_enrolments
			authroute.POST("/enrolments", server.CreateCourseEnrolment)
			authroute.GET("/enrolments", server.listEnrolments)
		//course_progress
			authroute.POST("/courseProgress", server.createCourseProgress)
			authroute.GET("/courseProgress", server.listCourseProgress)
			authroute.GET("/courseProgress/:courseProgress_id", server.getCourseProgress)

	server.router = router
}

// start runs the http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}


// func main() {
// 	r := gin.Default()

// 	// Public route
// 	r.GET("/public", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "This is a public route"})
// 	})

// 	// Protected route for admins
// 	adminRoutes := r.Group("/admin").Use(AuthorizeRole("admin"))
// 	adminRoutes.GET("/dashboard", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Welcome, Admin!"})
// 	})

// 	// Protected route for users
// 	userRoutes := r.Group("/user").Use(AuthorizeRole("user"))
// 	userRoutes.GET("/profile", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Welcome, User!"})
// 	})

// 	r.Run(":8080")
// }