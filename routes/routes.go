package routes

import (
	routes "nivasBackendMain/App/routes/Authentation"
	brandRegistrationRoutes "nivasBackendMain/Brand/routes/BrandRegistration"
	brandRoutes "nivasBackendMain/Brand/routes/RegisterFormDocument"
	nivasAdminRoutes "nivasBackendMain/NivasAdmin/routes/GenerateBrandRegisterUrl"

	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.Engine) {
	route := router.Group("api/v1/app")
	routes.SignUpRoutes(route)
	adminRoute := router.Group(("api/v1/nivasAdmin"))
	nivasAdminRoutes.BrandRegisterUrl(adminRoute)
	brandRoute := router.Group("api/v1/brand")
	brandRoutes.RegisterFormDocumentUploadUrl(brandRoute)
	brandRegistrationRoutes.GetBrandRegistrationFormData(brandRoute)
	brandRegistrationRoutes.GetBrandCurrentStatus(brandRoute)
	brandRegistrationRoutes.StoreBrandDetails(brandRoute)
	brandRegistrationRoutes.GetBrandRegistrationStatusData(brandRoute)
}
