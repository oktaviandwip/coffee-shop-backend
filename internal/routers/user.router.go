package routers

import (
	"coffeeshop/internal/handlers"
	"coffeeshop/internal/middleware"
	"coffeeshop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", middleware.UploadFile, handler.PostUser)
	route.GET("/:id", middleware.AuthJwt("admin", "user"), handler.GetUser)
	route.PATCH("/:id", middleware.AuthJwt("admin", "user"), middleware.UploadFile, handler.PatchUser)
	route.DELETE("/:id", middleware.AuthJwt("admin", "user"), handler.DeleteUser)
}
