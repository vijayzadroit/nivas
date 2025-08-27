package controller

import (
	"net/http"
	model "nivasBackendMain/App/model/Authentation"
	service "nivasBackendMain/App/service/Authentation"
	accesstoken "nivasBackendMain/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func SignUpController() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqVal model.GoogleToken
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}

		resVal := service.SignUpService(reqVal)
		token := accesstoken.CreateToken()

		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
			"name":    resVal.Name,
			"mail":    resVal.Mail,
			"picture": resVal.Profile,
			"token":   token,
		}
		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}

func CountryCodeController() gin.HandlerFunc {
	return func(c *gin.Context) {
		resVal := service.CountryCode()
		token := accesstoken.CreateToken()

		responce := gin.H{
			"status":      resVal.Status,
			"Message":     resVal.Message,
			"countryData": resVal.CountryData,
			"token":       token,
		}
		c.JSON(http.StatusOK, gin.H{
			"data": responce,
		})
	}
}

func MobileNumberValidationController() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqVal model.MobileNumberValidationRequest
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}

		resVal := service.MobileNumberValidationService(reqVal)
		
		token := accesstoken.CreateToken()
		responce := gin.H{
			"status":  resVal.Status,
			"Message": resVal.Message,
			"Code":    resVal.Code,
			"token":   token,
		}
		c.JSON(http.StatusOK, gin.H{
			"data": responce,
		})

	}
}
