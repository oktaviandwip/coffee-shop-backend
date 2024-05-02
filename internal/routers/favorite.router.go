package routers

import (
	"coffeeshop/internal/handlers"
	"coffeeshop/internal/middleware"
	"coffeeshop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.POST("/", middleware.AuthJwt("admin", "user"), handler.PostFavorite)
	route.GET("/", middleware.AuthJwt("admin", "user"), handler.GetFavorite)
	route.PATCH("/:user_id/:product_id", middleware.AuthJwt("admin", "user"), handler.PatchFavorite)
	route.DELETE("/:user_id/:product_id", middleware.AuthJwt("admin", "user"), handler.DeleteFavorite)
}
