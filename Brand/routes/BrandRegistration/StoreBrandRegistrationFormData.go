package brandRegistrationRoutes

import (
	brandRegistrationController "nivasBackendMain/Brand/controller/BrandRegistration"

	"github.com/gin-gonic/gin"
)

func StoreBrandDetails(router *gin.RouterGroup) {
	route := router.Group("brandRegistration")
	route.POST("/storeData", brandRegistrationController.StoreBrandRegistrationForm())
}
