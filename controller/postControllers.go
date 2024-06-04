package controller

import (
	"jwt-golang/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) PostController {
	return PostController{DB: db}
}

func (pc *PostController) GetPosts(ctx *gin.Context) {

	var posts []models.Post

	// Get page and limit query parameters
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	// Calculate offset
	offset := (page - 1) * limit

	result := pc.DB.Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": posts})
}

func (pc *PostController) GetPost(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var post models.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "post with given postID don't exist"})
		return
	}
	ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": post})
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload models.CreatePostRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	id, _ := uuid.NewV4()

	newPost := models.Post{
		ID:        id,
		Title:     payload.Title,
		Image:     payload.Image,
		Content:   payload.Content,
		User:      currentUser.ID,
		UpdatedAt: now,
		CreatedAt: now,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "post with title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Bad Gateway"})
		return
	}

	ctx.JSON(http.StatusConflict, gin.H{"status": "success", "message": newPost})
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	postId := ctx.Param("postId")

	var payload models.UpdatePostRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var postToBeUpdated models.Post
	result := pc.DB.First(&postToBeUpdated, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "post with given postID don't exist"})
		return
	}

	now := time.Now()

	updatedPost := models.Post{
		ID:        postToBeUpdated.ID,
		Title:     payload.Title,
		Image:     payload.Image,
		Content:   payload.Content,
		User:      currentUser.ID,
		UpdatedAt: now,
		CreatedAt: postToBeUpdated.CreatedAt,
	}

	result = pc.DB.Model(&postToBeUpdated).Updates(updatedPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Duplicate entry"})
		}
		ctx.JSON(http.StatusBadGateway, "badGateway")
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"status": "success", "message": updatedPost})
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("postId")

	result := pc.DB.Delete(&models.Post{}, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "failed"})
		return
	}

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "success"})
		return
	}
}
