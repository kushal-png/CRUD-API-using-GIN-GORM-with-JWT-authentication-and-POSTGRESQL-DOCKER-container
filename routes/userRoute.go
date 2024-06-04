package routes

import (
	"jwt-golang/controller"
	"jwt-golang/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controller.UserController
}

func NewRouteUserController(userController controller.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
