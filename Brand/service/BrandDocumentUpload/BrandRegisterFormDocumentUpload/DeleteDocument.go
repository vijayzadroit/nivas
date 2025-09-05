package registerFormDocumentService

import (
	"fmt"
	RegisterFormDocumentModel "nivasBackendMain/Brand/model/BrandDocument"
	logger "nivasBackendMain/Helper/Logger"
	service "nivasBackendMain/Helper/MinIo"

	"gorm.io/gorm"
)

func RequestToDeleteOldDocument(db *gorm.DB, reqVal RegisterFormDocumentModel.DocumentDeleteReq) RegisterFormDocumentModel.DocumentDeleteResponse {
	log := logger.InitLogger()

	// Check if OldFile is empty
	if reqVal.OldFile == "" {
		log.Warn("No old file provided for deletion")
		return RegisterFormDocumentModel.DocumentDeleteResponse{
			Status:  false,
			Message: "No old file provided for deletion",
		}
	}

	// Attempt to delete the old file
	err := service.DeleteFile(reqVal.OldFile)
	if err != nil {
		log.Errorf("Failed to delete old file %s: %v", reqVal.OldFile, err)
		return RegisterFormDocumentModel.DocumentDeleteResponse{
			Status:  false,
			Message: fmt.Sprintf("Failed to delete old document: %v", err),
		}
	}

	// Success
	log.Infof("Old document deleted successfully: %s", reqVal.OldFile)
	return RegisterFormDocumentModel.DocumentDeleteResponse{
		Status:  true,
		Message: "Old document deleted successfully",
	}
}
