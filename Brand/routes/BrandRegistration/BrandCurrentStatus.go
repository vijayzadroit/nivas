package brandRegistrationRoutes

import (
	brandRegistrationController "nivasBackendMain/Brand/controller/BrandRegistration"

	"github.com/gin-gonic/gin"
)

func GetBrandCurrentStatus(router *gin.RouterGroup) {
	route := router.Group("brandRegistration")
	route.POST("/getStatus", brandRegistrationController.GetBrandCurrentStatus())
}
