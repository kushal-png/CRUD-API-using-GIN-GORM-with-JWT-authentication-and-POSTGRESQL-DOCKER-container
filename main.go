package main

import (
	"jwt-golang/controller"
	"jwt-golang/initializers"
	"jwt-golang/routes"

	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controller.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controller.UserController
	UserRouteController routes.UserRouteController

	PostController      controller.PostController
	PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load environment variables")
	}

	initializers.ConnectDB(&config)
	AuthController = controller.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controller.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controller.NewPostController(initializers.DB)
	PostRouteController = routes.NewPostRouteController(PostController)
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoutes(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
