package brandRegistrationRoutes

import (
	brandRegistrationController "nivasBackendMain/Brand/controller/BrandRegistration"

	"github.com/gin-gonic/gin"
)

func GetBrandRegistrationStatusData(router *gin.RouterGroup) {
	route := router.Group("brandRegistration")
	route.POST("/getRegistrationStatus", brandRegistrationController.GetBrandRegistrationStatusData())
}
