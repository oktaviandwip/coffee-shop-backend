package middleware

import (
	"coffeeshop/config"
	"coffeeshop/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewRes(401, &config.Result{
				Data: "Please login",
			}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			pkg.NewRes(401, &config.Result{
				Data: "Invalid Header Type",
			}).Send(ctx)
			return
		}

		token := strings.Replace(header, "Bearer ", "", -1)
		check, err := pkg.VerifyToken(token)
		if err != nil {
			pkg.NewRes(401, &config.Result{
				Data: err.Error(),
			}).Send(ctx)
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.NewRes(401, &config.Result{
				Data: "You don't have permission",
			}).Send(ctx)
			return
		}

		ctx.Set("userId", check.Id)
		ctx.Next()
	}
}
