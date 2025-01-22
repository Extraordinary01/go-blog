package api

import (
	"go-blog/database/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type AuthorController struct {
	db *pg.DB
}

func (controller *AuthorController) List(context *gin.Context) {
	authors, err := models.GetAuthors(controller.db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, authors)
}

func (controller *AuthorController) Create(context *gin.Context) {
	var author models.Author
	context.BindJSON(&author)
	err := author.Create(controller.db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, author)
}

func (controller *AuthorController) Retrieve(context *gin.Context) {
	authorId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}
	author := &models.Author{}
	err = author.Get(controller.db, int64(authorId))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, author)
}

func (controller *AuthorController) Update(context *gin.Context) {
	authorId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}
	author := &models.Author{}
	context.BindJSON(&author)
	err = author.Update(controller.db, int64(authorId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, author)
}

func (controller *AuthorController) Delete(context *gin.Context) {
	authorId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}
	author := &models.Author{}
	err = author.Delete(controller.db, int64(authorId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}

func GenerateAuthorRoutes(db *pg.DB, router *gin.Engine) {
	authors := router.Group("/authors")
	authorHandler := &AuthorController{db}
	authors.GET("/", authorHandler.List)
	authors.POST("/", authorHandler.Create)
	authors.GET("/:id", authorHandler.Retrieve)
	authors.PUT("/:id", authorHandler.Update)
	authors.DELETE("/:id", authorHandler.Delete)
}