package brandRoutes

import (
	registerFormDocumentController "nivasBackendMain/Brand/controller/BrandDocument"

	"github.com/gin-gonic/gin"
)

func RegisterFormDocumentUploadUrl(router *gin.RouterGroup) {
	route := router.Group("registerFormDocument")
	route.POST("/generateUploadUrl", registerFormDocumentController.BrandDocumentUploadUrl())
	route.POST("/DeleteDocument", registerFormDocumentController.DocumentDelete())
}
