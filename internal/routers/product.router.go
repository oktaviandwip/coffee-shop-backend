package routers

import (
	"coffeeshop/internal/handlers"
	"coffeeshop/internal/middleware"
	"coffeeshop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repository.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.POST("/", middleware.AuthJwt("admin"), middleware.UploadFile, handler.PostProduct)
	route.GET("/", middleware.AuthJwt("admin", "user"), handler.GetProduct)
	route.PATCH("/:id", middleware.AuthJwt("admin"), middleware.UploadFile, handler.PatchProduct)
	route.DELETE(":id", middleware.AuthJwt("admin"), handler.DeleteProduct)
}
