package handlers

import (
	"coffeeshop/config"
	"coffeeshop/internal/models"
	"coffeeshop/internal/repository"
	"coffeeshop/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repository.RepoUserIF
}

func NewUser(r repository.RepoUserIF) *HandlerUser {
	return &HandlerUser{r}
}

// Create User
func (h *HandlerUser) PostUser(ctx *gin.Context) {
	var err error
	user := models.User{
		Role: "user",
	}

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	user.Photo_profile = ctx.MustGet("profileImage").(string)

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(201, result).Send(ctx)
}

// Get User
func (h *HandlerUser) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := h.FetchUser(id)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

// Update User
func (h *HandlerUser) PatchUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := models.User{}
	var err error

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	user.Photo_profile = ctx.MustGet("profileImage").(string)

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.UpdateUser(id, &user)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

// Delete User
func (h *HandlerUser) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := h.RemoveUser(id)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}
