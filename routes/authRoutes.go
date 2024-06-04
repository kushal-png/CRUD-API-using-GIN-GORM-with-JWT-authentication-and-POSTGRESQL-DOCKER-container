package routes

import (
	"jwt-golang/controller"
	"jwt-golang/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controller.AuthController
}

func NewAuthRouteController(authController controller.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
