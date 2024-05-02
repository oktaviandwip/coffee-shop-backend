package handlers

import (
	"coffeeshop/config"
	"coffeeshop/internal/repository"
	"coffeeshop/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password" form:"password"`
}

type HandlerAuth struct {
	*repository.RepoUser
}

func NewAuth(r *repository.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	data := User{}

	if err := ctx.ShouldBind(&data); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user, err := h.GetAuthData(data.Email)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(user.Password, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Wrong password",
		}).Send(ctx)
		return
	}

	jwt := pkg.NewToken(user.User_id, user.Role)
	token, err := jwt.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: token}).Send(ctx)
}
