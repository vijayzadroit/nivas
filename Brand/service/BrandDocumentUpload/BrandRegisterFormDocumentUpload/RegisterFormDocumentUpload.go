package registerFormDocumentService

import (
	"fmt"
	"math/rand"
	RegisterFormDocumentModel "nivasBackendMain/Brand/model/BrandDocument"
	logger "nivasBackendMain/Helper/Logger"
	service "nivasBackendMain/Helper/MinIo"
	"strings"
	"time"

	"gorm.io/gorm"
)

func RequestUploadUrl(db *gorm.DB, reqVal RegisterFormDocumentModel.UploadUrlRequestData) RegisterFormDocumentModel.UploadUrlResponse {
	log := logger.InitLogger()

	// Get file extension safely
	fileParts := strings.Split(reqVal.FileName, ".")
	extension := ""
	if len(fileParts) > 1 {
		extension = "." + fileParts[len(fileParts)-1]
	}

	// Random 6-digit number
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(900000) + 100000

	// Final filename
	fileName := fmt.Sprintf("%s_%d%s", reqVal.BrandName, number, extension)
	fmt.Println("File Name:", fileName)

	// Generate signed URLs
	uploadUrl, fileUrl, err := service.CreateUploadURL(fileName, 10)
	if err != nil {
		log.Errorf("failed to generate upload URL: %v", err)
		return RegisterFormDocumentModel.UploadUrlResponse{
			Status:    false,
			Message:   "Failed to generate upload URL",
			UploadUrl: "",
			FileUrl:   "",
			FilePath:  "",
		}
	}

	// Success
	return RegisterFormDocumentModel.UploadUrlResponse{
		Status:    true,
		Message:   "Upload URL generated successfully",
		UploadUrl: uploadUrl,
		FileUrl:   fileUrl,
		FilePath:  fileName,
	}
}
