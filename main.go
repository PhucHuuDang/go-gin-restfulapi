package main

import (
	"golang-restapi/db"
	_ "golang-restapi/docs"
	"golang-restapi/handler"
	"golang-restapi/models"
	"golang-restapi/repositories"
	"log"
	"net/http"
	"os"
	"time"
	
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	
	// mux "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	
	uuid "github.com/google/uuid"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "localhost:*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
	}
	
}

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		now := time.Now()
		start := now.Format(time.RFC822)
		path := c.Request.URL.Path
		method := c.Request.Method
		
		log.Printf("[%s] [%s] [%s]\n", method, path, start)
		
		c.Next()
		
	}
}

// func TokenAuthMiddleware() *gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		auth.Token
// 		c.Next()
//
// 	}
// }

// var auth = new(controller.AuthController)

// @title GOLANG GIN RESTFULL API

// @version         1.0
// @description     A RESTful API boilerplate with Gin Framework, PostgreSQL, Redis and JWT authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env file")
	}
	
	// fmt.Println("DB_URL =", os.Getenv("DB_URL"))
	
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	router := gin.Default()
	
	// custom
	// router.TrustedPlatform = "X-CDN-Client-IP"
	
	router.Use(CORSMiddleware())
	router.Use(RequestIdMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(LoggerMiddleware())
	
	//NOTE: we have to pay attention to the order of the model cause it has relationship,
	// so we need to priority the model have main key and put it first
	DB := db.Init(&models.AuthModel{}, &models.UserModel{}, &models.PostModel{})
	
	// User
	userRepo := repositories.NewUserRepository(DB)
	userHandler := handler.NewUserHandler(userRepo)
	
	v1 := router.Group("/v1")
	{
		// user
		v1.GET("/", Helloworld)
		authRoute := v1.Group("auth")
		
		{
			authRoute.POST("/login")
			authRoute.POST("/register")
			authRoute.POST("/refresh")
		}
		
		userRoute := v1.Group("/user")
		{
			userRoute.GET("", userHandler.GetAllUsers)
			userRoute.POST("", userHandler.CreateUser)
			userRoute.DELETE("/:id", userHandler.DeleteUser)
			userRoute.PUT("/:id", userHandler.UpdateUser)
			userRoute.PATCH("/:id", userHandler.UpdateUser)
		}
		
	}
	
	router.Run("localhost:8080")
}
