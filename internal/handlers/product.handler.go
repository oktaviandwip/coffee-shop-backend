package handlers

import (
	"coffee/internal/models"
	"coffee/internal/repository"
	"coffee/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repository.RepoProduct
}

func NewProduct(r *repository.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

// Create Product
func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.CreateProduct(&product)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusCreated, result)
}

// Get Product
func (h *HandlerProduct) GetProduct(ctx *gin.Context) {
	search := ctx.Query("search")
	sort := ctx.Query("sort")
	pageStr := ctx.Query("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}
	offset := (page - 1) * 10

	if len(search) > 0 {
		result, err := h.SearchProduct(search, offset)
		if err != nil {
			pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
			return
		}
		pkg.Response(ctx.Writer, http.StatusOK, result)
		return
	}

	if len(sort) > 0 {
		result, err := h.SortProduct(sort, offset)
		if err != nil {
			pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
			return
		}
		pkg.Response(ctx.Writer, http.StatusOK, result)
		return
	}

	result, err := h.ReadProduct(offset)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}

// Update Product
func (h *HandlerProduct) PatchProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.UpdateProduct(id, &product)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}

// Delete Product
func (h *HandlerProduct) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.RemoveProduct(id, &product)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}
