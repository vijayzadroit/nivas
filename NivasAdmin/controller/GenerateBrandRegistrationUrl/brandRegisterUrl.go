package controller

import (
	"net/http"

	db "nivasBackendMain/DB"
	model "nivasBackendMain/NivasAdmin/model/GenerateBrandRegistrationUrl"
	service "nivasBackendMain/NivasAdmin/service/GenerateBrandRengistrationUrl"

	"github.com/gin-gonic/gin"
)

func BrandRegisterUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal model.BrandRegistrationUrlRequest

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}

		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := service.BrandRegistrationUrlService(dbConnt, reqVal)

		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}
