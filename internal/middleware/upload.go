package middleware

import (
	"coffeeshop/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	profileFile, err := ctx.FormFile("photo_profile")
	if err != nil {
		if err.Error() == "http: no such file" {
			ctx.Set("profileImage", "")
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving profile photo"})
			return
		}
	}

	productFile, err := ctx.FormFile("photo_product")
	if err != nil {
		if err.Error() == "http: no such file" {
			ctx.Set("productImage", "")
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving product photo"})
			return
		}
	}

	var profileImageUrl, productImageUrl string
	if profileFile != nil {
		src, err := profileFile.Open()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error opening profile photo"})
			return
		}
		defer src.Close()

		profileImageUrl, err = pkg.CloudInary(src)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error uploading profile photo"})
			return
		}
	}

	if productFile != nil {
		src, err := productFile.Open()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error opening product photo"})
			return
		}
		defer src.Close()

		productImageUrl, err = pkg.CloudInary(src)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error uploading product photo"})
			return
		}
	}

	ctx.Set("profileImage", profileImageUrl)
	ctx.Set("productImage", productImageUrl)
	ctx.Next()
}
