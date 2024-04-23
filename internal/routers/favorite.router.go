package routers

import (
	"coffee/internal/handlers"
	"coffee/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.POST("/", handler.PostFavorite)
	route.GET("/", handler.GetFavorite)
	route.PATCH("/:user_id/:product_id", handler.PatchFavorite)
	route.DELETE("/:user_id/:product_id", handler.DeleteFavorite)
}
