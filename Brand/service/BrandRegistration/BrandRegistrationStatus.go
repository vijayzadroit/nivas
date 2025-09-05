package brandRegistrationService

import (
	"fmt"
	brandRegistrationModel "nivasBackendMain/Brand/model/BrandRegistration"
	brandRegistrationQuery "nivasBackendMain/Brand/query/BrandRegistration"
	logger "nivasBackendMain/Helper/Logger"
	service "nivasBackendMain/Helper/MinIo"
	"time"

	"gorm.io/gorm"
)

func parseAndFormatDate(dateStr string) string {
	if dateStr == "" {
		return ""
	}

	// Try parsing full timestamp first
	if t, err := time.Parse("2006-01-02 15:04:05", dateStr); err == nil {
		return t.Format("Jan 02, 2006")
	}

	// Try parsing date only
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t.Format("Jan 02, 2006")
	}

	// If all fails, return raw string
	return dateStr
}

// FormatBrandApplication converts DB response -> API response format
func FormatBrandApplication(resp brandRegistrationModel.BrandApplicationResponse) brandRegistrationModel.BrandApplicationFormatted {
	var formatted brandRegistrationModel.BrandApplicationFormatted

	// ---- Brand Application Status ----
	formatted.BrandApplicationStatus.Status = resp.ApplicationStatusId
	formatted.BrandApplicationStatus.StatusName = resp.ApplicationStatus
	formatted.BrandApplicationStatus.ApplicationId = resp.ApplicationCustId
	formatted.BrandApplicationStatus.SubmitDate = parseAndFormatDate(resp.CreateAt)
	formatted.BrandApplicationStatus.LastDate = parseAndFormatDate(resp.UpdateAt)
	formatted.BrandApplicationStatus.ProcessTime = resp.ProcessTime

	// ---- Brand Information ----
	brandLogoUrl, err := service.GetFileURL(resp.Logo, 30)
	if err != nil {
		fmt.Println("Error generating BrandLogo URL:", err)
		brandLogoUrl = ""
	}
	formatted.BrandInformation.BrandLogo = brandLogoUrl
	formatted.BrandInformation.BrandName = resp.BrandName
	formatted.BrandInformation.BrandCategory = resp.BrandCategoryName
	formatted.BrandInformation.CinNumber = resp.Cin
	formatted.BrandInformation.ContactPerson = resp.BrandContactPerson
	formatted.BrandInformation.SubmitDate = parseAndFormatDate(resp.CreateAt)
	formatted.BrandInformation.PhoneNumber = resp.BrandMobile
	formatted.BrandInformation.Email = resp.BrandEmail

	// ---- Documents ----
	formatted.Document.ShowDocument = resp.ApplicationStatusId == 3 || resp.ApplicationStatusId == 5

	addViewUrl, _ := service.GetFileURL(resp.AddressProf, 30)
	addDownloadUrl, _ := service.CreateDownloadURL(resp.AddressProf, resp.AddressProfName, 30)
	addressDocSize, _ := service.GetFileInfo(resp.AddressProf)

	formatted.Document.AddressProof = brandRegistrationModel.DocumentLink{
		Url:         addViewUrl,
		DownloadUrl: addDownloadUrl,
		FileSize:    addressDocSize,
	}

	gsrViewUrl, _ := service.GetFileURL(resp.GstinDoc, 30)
	gstDownloadUrl, _ := service.CreateDownloadURL(resp.GstinDoc, resp.GstDocName, 30)
	gstDocSize, _ := service.GetFileInfo(resp.GstinDoc)

	formatted.Document.GstDocument = brandRegistrationModel.DocumentLink{
		Url:         gsrViewUrl,
		DownloadUrl: gstDownloadUrl,
		FileSize:    gstDocSize,
	}

	panViewUrl, _ := service.GetFileURL(resp.PanDoc, 30)
	panDownloadUrl, _ := service.CreateDownloadURL(resp.PanDoc, resp.PanDocName, 30)
	panDocSize, _ := service.GetFileInfo(resp.PanDoc)
	formatted.Document.PanDocument = brandRegistrationModel.DocumentLink{
		Url:         panViewUrl,
		DownloadUrl: panDownloadUrl,
		FileSize:    panDocSize,
	}

	// ---- Feedback ----
	formatted.Feedback.CurrentStatus = resp.CurrentStatus
	formatted.Feedback.ReviewContent = resp.ApplicationStatusDescription

	return formatted
}

// GetBrandStatusData fetches brand application status from DB
func GetBrandStatusData(db *gorm.DB, reqVal brandRegistrationModel.GetBrandRegistrationStatusReq) brandRegistrationModel.BrandRegistrationStatus {
	log := logger.InitLogger()

	var brandDetails []brandRegistrationModel.BrandApplicationResponse

	// Execute query
	err := db.Raw(brandRegistrationQuery.GetBrandRegistrationStatusData, reqVal.ApplicationId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("❌ Error fetching brand status from DB: " + err.Error())
		return brandRegistrationModel.BrandRegistrationStatus{
			Status:  false,
			Message: "Something went wrong while fetching data. Please try again.",
		}
	}
	fmt.Println("DB Data : ", brandDetails)

	// If no data found
	if len(brandDetails) == 0 {
		log.Warn("⚠️ No brand details found for ApplicationId: " + fmt.Sprint(reqVal.ApplicationId))
		return brandRegistrationModel.BrandRegistrationStatus{
			Status:  false,
			Message: "No brand data found for the provided application ID",
		}
	}

	log.Info(fmt.Sprintf("✅ Brand data fetched successfully for ApplicationId: %v", reqVal.ApplicationId))

	// Format and return response
	return brandRegistrationModel.BrandRegistrationStatus{
		Status:    true,
		Message:   "Brand registration Status Passed Successfully",
		BrandData: FormatBrandApplication(brandDetails[0]),
	}
}
