package routes

import (
	"jwt-golang/controller"
	"jwt-golang/middleware"

	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controller.PostController
}

func NewPostRouteController (pc controller.PostController)(PostRouteController) {
   return PostRouteController{
	postController: pc,
   }
}
func (pc *PostRouteController) PostRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/posts")
	router.GET("/", middleware.DeserializeUser(), pc.postController.GetPosts)
	router.GET("/:postId", middleware.DeserializeUser(), pc.postController.GetPost)
	router.POST("/", middleware.DeserializeUser(), pc.postController.CreatePost)
	router.DELETE("/:postId", middleware.DeserializeUser(), pc.postController.DeletePost)
	router.PUT("/:postId", middleware.DeserializeUser(), pc.postController.UpdatePost)
}
