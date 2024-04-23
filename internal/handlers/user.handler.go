package handlers

import (
	"coffee/internal/models"
	"coffee/internal/repository"
	"coffee/pkg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repository.RepoUser
}

func NewUser(r *repository.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

// Create User
func (h *HandlerUser) PostUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusCreated, result)
}

// Get User
func (h *HandlerUser) GetUser(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}
	offset := (page - 1) * 10

	result, err := h.ReadUser(offset)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}

// Update User
func (h *HandlerUser) PatchUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.UpdateUser(id, &user)
	if err != nil {
		fmt.Println(err)
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}

// Delete User
func (h *HandlerUser) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.RemoveUser(id, &user)
	if err != nil {
		pkg.Response(ctx.Writer, http.StatusBadRequest, err.Error())
		return
	}

	pkg.Response(ctx.Writer, http.StatusOK, result)
}
