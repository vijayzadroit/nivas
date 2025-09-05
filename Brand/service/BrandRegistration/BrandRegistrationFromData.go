package brandRegistrationService

import (
	"fmt"
	brandRegistrationModel "nivasBackendMain/Brand/model/BrandRegistration"
	brandRegistrationQuery "nivasBackendMain/Brand/query/BrandRegistration"
	logger "nivasBackendMain/Helper/Logger"

	"gorm.io/gorm"
)

// MapBrandDetails maps flat DB response to nested JSON
func MapBrandDetails(brandDetails brandRegistrationModel.GetBrandRegisterFormDataResFromDb) brandRegistrationModel.GetBrandRegisterDataResponse {
	var resp brandRegistrationModel.GetBrandRegisterDataResponse

	// Brand Information
	resp.BrandInformation.BrandName = brandDetails.BrandName
	resp.BrandInformation.ProductCategory = brandDetails.ProductCategory
	resp.BrandInformation.BrandLogoPath = brandDetails.BrandLogoPath
	resp.BrandInformation.BrandDescription = brandDetails.BrandDescription
	resp.BrandInformation.WebsiteURL = brandDetails.WebsiteURL
	resp.BrandInformation.Instragram = brandDetails.Instragram
	resp.BrandInformation.BrandLogoName = brandDetails.BrandLogoName

	// Contact Information
	resp.ContactInformation.ContactPerson = brandDetails.ContactPerson
	resp.ContactInformation.Designation = brandDetails.Designation
	resp.ContactInformation.PhoneNumber = brandDetails.PhoneNumber
	resp.ContactInformation.Email = brandDetails.Email
	resp.ContactInformation.Address = brandDetails.Address
	resp.ContactInformation.City = brandDetails.City
	resp.ContactInformation.ZipCode = brandDetails.ZipCode
	resp.ContactInformation.State = brandDetails.State
	resp.ContactInformation.ProofDocument = brandDetails.ProofDocument
	resp.ContactInformation.ProffDocumentName = brandDetails.ProffDocumentName

	// Tax Information
	resp.TaxInformation.GstinNumber = brandDetails.GstinNumber
	resp.TaxInformation.CinNumber = brandDetails.CinNumber
	resp.TaxInformation.GstDocumant = brandDetails.GstDocumant
	resp.TaxInformation.PanDocument = brandDetails.PanDocument
	resp.TaxInformation.GstDocumentName = brandDetails.GstDocumentName
	resp.TaxInformation.PanDocumantName = brandDetails.PanDocumantName

	// Warehouse Information
	resp.WareHouseInfo.WareHouse = brandDetails.WareHouse
	resp.WareHouseInfo.WareHouseAddress = brandDetails.WareHouseAddress
	resp.WareHouseInfo.WareHouseCity = brandDetails.WareHouseCity
	resp.WareHouseInfo.WareHouseDistrict = brandDetails.WareHouseDistrict
	resp.WareHouseInfo.WareHouseZipCode = brandDetails.WareHouseZipCode
	resp.WareHouseInfo.WareHouseState = brandDetails.WareHouseState

	// Application Type
	resp.ApplicationType.SaveDraft = brandDetails.SaveDraft
	resp.ApplicationType.RefApplicationStatusId = brandDetails.ApplicationStatusId
	resp.ApplicationType.RefApplicationStatus = brandDetails.ApplicationStatus

	return resp
}

// GetBrandRegistrationFormData fetches brand registration data from DB and maps it
func GetBrandRegistrationFormData(db *gorm.DB, reqVal brandRegistrationModel.GetBrandRegistrationFromDataReq) brandRegistrationModel.GetBrandRegistrationDataFinalRes {
	log := logger.InitLogger()

	var brandDetails brandRegistrationModel.GetBrandRegisterFormDataResFromDb

	err := db.Raw(brandRegistrationQuery.GetBrandRegisterFormData, reqVal.ApplicationId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Error in getting the Brand Register Form Data: " + err.Error())
		return brandRegistrationModel.GetBrandRegistrationDataFinalRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}
	fmt.Println("Db Data : ", brandDetails)
	response := MapBrandDetails(brandDetails)
	fmt.Println("format Data : ", response)

	return brandRegistrationModel.GetBrandRegistrationDataFinalRes{
		Status:    true,
		Message:   "Brand registration data fetched successfully",
		BrandData: response,
	}
}
