package api

import (
	"go-blog/database/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)


type BlogController struct {
	db *pg.DB
}


func (controller *BlogController) List(context *gin.Context) {
	blogs, err := models.GetBlogs(controller.db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, blogs)
}


func (controller *BlogController) Create(context *gin.Context) {
	var blog models.Blog
	if err := context.BindJSON(&blog);  err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := blog.Create(controller.db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, blog)
}


func (controller *BlogController) Retrieve(context *gin.Context) {
	blogId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	blog := &models.Blog{}
	err = blog.Get(controller.db, int64(blogId))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, blog)
}


func (controller *BlogController) Update(context *gin.Context) {
	blogId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	blog := &models.Blog{}
	context.BindJSON(&blog)
	err = blog.Update(controller.db, int64(blogId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, blog)
}


func (controller *BlogController) Delete(context *gin.Context) {
	blogId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	blog := &models.Blog{}
	err = blog.Delete(controller.db, int64(blogId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}


func GenerateBlogRoutes(db *pg.DB, router *gin.Engine) {
	blogs := router.Group("/blogs")
	controller := BlogController{db: db}
	blogs.GET("/", controller.List)
	blogs.POST("/", controller.Create)
	blogs.GET("/:id", controller.Retrieve)
	blogs.PUT("/:id", controller.Update)
	blogs.DELETE("/:id", controller.Delete)
}