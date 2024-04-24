package handlers

import (
	"coffee/internal/models"
	"coffee/internal/repository"
	"coffee/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	*repository.RepoFavorite
}

func NewFavorite(r *repository.RepoFavorite) *HandlerFavorite {
	return &HandlerFavorite{r}
}

// Create Favorite
func (h *HandlerFavorite) PostFavorite(ctx *gin.Context) {
	var favorite models.Favorite
	if err := ctx.ShouldBind(&favorite); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.CreateFavorite(&favorite)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusCreated, result)
}

// Get Favorite
func (h *HandlerFavorite) GetFavorite(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}
	offset := (page - 1) * 10

	result, err := h.ReadFavorite(user_id, offset)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}

// Update Favorite
func (h *HandlerFavorite) PatchFavorite(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	product_id := ctx.Param("product_id")

	var favorite models.Favorite

	if err := ctx.ShouldBind(&favorite); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.UpdateFavorite(user_id, product_id, &favorite)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}

// Delete Favorite
func (h *HandlerFavorite) DeleteFavorite(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	product_id := ctx.Param("product_id")

	var favorite models.Favorite

	if err := ctx.ShouldBind(&favorite); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.RemoveFavorite(user_id, product_id, &favorite)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}
