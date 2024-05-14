package handlers

import (
	"coffeeshop/config"
	"coffeeshop/internal/models"
	"coffeeshop/internal/repository"
	"coffeeshop/pkg"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	repository.RepoProductIF
}

func NewProduct(r repository.RepoProductIF) *HandlerProduct {
	return &HandlerProduct{r}
}

// Create Product
func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	var err error
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if result, ok := ctx.Get("productImage"); ok && result != nil {
		product.Photo_product = result.(string)
	}

	_, err = govalidator.ValidateStruct(&product)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.CreateProduct(&product)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(201, result).Send(ctx)
}

// Get Product
func (h *HandlerProduct) GetProduct(ctx *gin.Context) {
	search := ctx.Query("search")
	sort := ctx.Query("sort")
	pageStr := ctx.Query("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	offset := (page - 1) * 10

	// Search Products
	if len(search) > 0 {
		result, err := h.SearchProduct(search, page, offset)
		if err != nil {
			pkg.NewRes(400, &config.Result{
				Data: err.Error(),
			}).Send(ctx)
			return
		}
		pkg.NewRes(200, result).Send(ctx)
		return
	}

	// Sort Products
	if len(sort) > 0 {
		result, err := h.SortProduct(sort, page, offset)
		if err != nil {
			pkg.NewRes(400, &config.Result{
				Data: err.Error(),
			}).Send(ctx)
			return
		}
		pkg.NewRes(200, result).Send(ctx)
		return
	}

	// All Products
	result, err := h.FetchProduct(page, offset)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	pkg.NewRes(200, result).Send(ctx)
}

// Update Product
func (h *HandlerProduct) PatchProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product := models.Product{}
	var err error

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if result, ok := ctx.Get("productImage"); ok && result != nil {
		product.Photo_product = result.(string)
	}

	_, err = govalidator.ValidateStruct(&product)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.UpdateProduct(id, &product)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

// Delete Product
func (h *HandlerProduct) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := h.RemoveProduct(id)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}
