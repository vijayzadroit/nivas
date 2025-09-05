package brandRegistrationRoutes

import (
	brandRegistrationController "nivasBackendMain/Brand/controller/BrandRegistration"

	"github.com/gin-gonic/gin"
)

func GetBrandRegistrationFormData(router *gin.RouterGroup) {
	route := router.Group("brandRegistration")
	route.POST("/getFormData", brandRegistrationController.GetBrandRegistrationData())
	// route.POST("/DeleteDocument", registerFormDocumentController.DocumentDelete())
}
