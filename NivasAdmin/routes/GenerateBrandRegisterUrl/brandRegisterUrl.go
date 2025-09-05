package nivasAdminRoutes

import (
	controller "nivasBackendMain/NivasAdmin/controller/GenerateBrandRegistrationUrl"

	"github.com/gin-gonic/gin"
)

func BrandRegisterUrl(router *gin.RouterGroup) {
	route := router.Group("brandRegistration")
	route.POST("/generateUrl", controller.BrandRegisterUrl())
}
