package registerFormDocumentController

import (
	"net/http"
	RegisterFormDocumentModel "nivasBackendMain/Brand/model/BrandDocument"
	registerFormDocumentService "nivasBackendMain/Brand/service/BrandDocumentUpload/BrandRegisterFormDocumentUpload"
	db "nivasBackendMain/DB"

	"github.com/gin-gonic/gin"
)

func BrandDocumentUploadUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal RegisterFormDocumentModel.UploadUrlRequestData

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}

		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := registerFormDocumentService.RequestUploadUrl(dbConnt, reqVal)

		response := gin.H{
			"status":    resVal.Status,
			"message":   resVal.Message,
			"uploadUrl": resVal.UploadUrl,
			"fileUrl":   resVal.FileUrl,
			"filePath":  resVal.FilePath,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}
