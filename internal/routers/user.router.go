package routers

import (
	"coffee/internal/handlers"
	"coffee/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.PostUser)
	route.GET("/", handler.GetUser)
	route.PATCH("/:id", handler.PatchUser)
	route.DELETE("/:id", handler.DeleteUser)
}
