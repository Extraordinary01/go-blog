package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)


func InitAPI(db *pg.DB) *gin.Engine {
	route := gin.Default()
	route.GET("/ping", func (context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	GenerateAuthorRoutes(db, route)
	GenerateBlogRoutes(db, route)
	return route
}