package handler

import (
	"github.com/gin-gonic/gin"
	go_blog "go-blog"
	"net/http"
	"strconv"
)

// CreateLike godoc
// @Summary Create like
// @Security ApiKeyAuth
// @Description Like existing post
// @Tags like
// @ID create-like
// @Accept json
// @Produce json
// @Param input body go_blog.Like true "like"
// @Success 200 {int} int "id"
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/likes [post]
func (h *Handler) createLike(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var input go_blog.Like
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.UserId = userId

	id, err := h.services.Post.CreateLike(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// DeleteLike godoc
// @Summary Delete like
// @Security ApiKeyAuth
// @Description Remove like from post
// @Tags like
// @ID delete-like
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 204 {object} StatusResponse
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/likes/{id} [delete]
func (h *Handler) deleteLike(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "wrong id type")
		return
	}
	if err = h.services.Post.DeleteLike(id, userId); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusNoContent, StatusResponse{"ok"})
}
