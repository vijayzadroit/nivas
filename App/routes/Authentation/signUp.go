package routes

import (
	controller "nivasBackendMain/App/controller/Authentation"

	"github.com/gin-gonic/gin"
)

func SignUpRoutes(router *gin.RouterGroup) {
	route := router.Group("/signUp")
	route.POST("/googleLogin", controller.SignUpController())
	route.GET("/countryCode", controller.CountryCodeController())
	route.POST("/mobileNumberValidation", controller.MobileNumberValidationController())
}
