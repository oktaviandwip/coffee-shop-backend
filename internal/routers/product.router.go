package routers

import (
	"coffee/internal/handlers"
	"coffee/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repository.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.POST("/", handler.PostProduct)
	route.GET("/", handler.GetProduct)
	route.PATCH("/:id", handler.PatchProduct)
	route.DELETE("/:id", handler.DeleteProduct)
}
