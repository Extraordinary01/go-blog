package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	go_blog "go-blog"
	"net/http"
	"strconv"
)

// CreatePost godoc
// @Summary Create post
// @Security ApiKeyAuth
// @Description Create new post
// @Tags post
// @ID create-post
// @Accept json
// @Produce json
// @Param input body go_blog.Post true "post"
// @Success 200 {int} int "id"
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/post [post]
func (h *Handler) createPost(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var post go_blog.Post

	if err := ctx.BindJSON(&post); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	post.UserId = userId

	id, err := h.services.Post.CreatePost(post)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type getAllPostsResponse struct {
	Data []*go_blog.Post `json:"data"`
}

// GetAllPosts godoc
// @Summary Get all posts
// @Security ApiKeyAuth
// @Tags post
// @ID get-all-posts
// @Accept json
// @Produce json
// @Success 200 {object} getAllPostsResponse
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/post [get]
func (h *Handler) getAllPosts(ctx *gin.Context) {
	posts, err := h.services.Post.GetAllPosts()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllPostsResponse{posts})
}

// GetPost godoc
// @Summary Get post
// @Security ApiKeyAuth
// @Description Get post by id
// @Tags post
// @ID get-post
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} go_blog.Post
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/post/{id} [get]
func (h *Handler) getPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "wrong id type")
		return
	}

	post, err := h.services.Post.GetPost(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update post
// @Security ApiKeyAuth
// @Description Update post by id
// @Tags post
// @ID update-post
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param input body go_blog.PostUpdateInput true "post"
// @Success 200 {object} go_blog.PostUpdateInput
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/post/{id} [put]
func (h *Handler) updatePost(ctx *gin.Context) {
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

	var input go_blog.PostUpdateInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Post.UpdatePost(id, userId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{"ok"})
}

// DeletePost godoc
// @Summary Delete post
// @Security ApiKeyAuth
// @Description Delete post by id
// @Tags post
// @ID delete-post
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} StatusResponse
// @Failure 400,401,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/post/{id} [delete]
func (h *Handler) deletePost(ctx *gin.Context) {
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

	err = h.services.Post.DeletePost(id, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, StatusResponse{"ok"})
}

func getUserId(ctx *gin.Context) (int, error) {
	rawId, ok := ctx.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")

	}
	userId, ok := rawId.(int)
	if !ok {
		return 0, errors.New("user id is not integer")
	}
	return userId, nil
}
