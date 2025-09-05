package brandRegistrationController

import (
	"net/http"
	brandRegistrationModel "nivasBackendMain/Brand/model/BrandRegistration"
	brandRegistrationService "nivasBackendMain/Brand/service/BrandRegistration"
	db "nivasBackendMain/DB"

	"github.com/gin-gonic/gin"
)

func StoreBrandRegistrationForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal brandRegistrationModel.StoreBrandRegisterFormDataReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := brandRegistrationService.StoreBrandDetails(dbConnt, reqVal)
		response := gin.H{
			"status":            resVal.Status,
			"message":           resVal.Message,
			"applicationStatus": resVal.ApplicationStatus,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}
