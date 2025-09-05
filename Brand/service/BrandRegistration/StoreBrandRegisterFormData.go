package brandRegistrationService

import (
	"fmt"
	brandRegistrationModel "nivasBackendMain/Brand/model/BrandRegistration"
	brandRegistrationQuery "nivasBackendMain/Brand/query/BrandRegistration"
	getDate "nivasBackendMain/Helper/DateFormat"
	logger "nivasBackendMain/Helper/Logger"

	"gorm.io/gorm"
)

func StoreBrandDetails(db *gorm.DB, reqVal brandRegistrationModel.StoreBrandRegisterFormDataReq) brandRegistrationModel.StoreBrandRegisterFormDetailsRes {
	log := logger.InitLogger()
	var date = getDate.GetCurrentDate("")

	var brandDetails brandRegistrationModel.BrandReferenceIdRes

	var brandStatus int
	if reqVal.SaveDraft {
		brandStatus = 2
	} else {
		brandStatus = 3
	}

	err1 := db.Raw(brandRegistrationQuery.StoreBrandInfoQuery, reqVal.BrandInformation.BrandName, reqVal.BrandInformation.ProductCategory, reqVal.BrandInformation.BrandDescription, brandStatus, reqVal.TaxInformation.GstinNumber, reqVal.TaxInformation.CinNumber, date, reqVal.ApplicationId).
		Scan(&brandDetails).Error
	if err1 != nil {
		log.Error("Error in Updating the Brand InFormation Data: " + err1.Error())
		return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	fmt.Println("Db Data : ", brandDetails)
	fmt.Println("Db Data : ", brandDetails.ContactId)
	fmt.Println("Db Data : ", reqVal.ContactInformation.ContactPerson)
	fmt.Println("Db Data : ", reqVal.ContactInformation.Email)
	fmt.Println("Db Data : ", reqVal.ContactInformation.PhoneNumber)
	fmt.Println("Db Data : ", reqVal.ContactInformation.Designation)

	err2 := db.Exec(
		brandRegistrationQuery.StoreBrandContactQuery,
		brandDetails.ContactId,
		reqVal.ContactInformation.ContactPerson,
		reqVal.ContactInformation.Email,
		reqVal.ContactInformation.PhoneNumber,
		reqVal.ContactInformation.Designation,
	).Error
	if err2 != nil {
		log.Error("Error in Updating Brand Contact Details: ", err2)
		return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	fmt.Println("Db Data : ", brandDetails.LocationId)
	fmt.Println("Db Data : ", reqVal.ContactInformation.Address)
	fmt.Println("Db Data : ", reqVal.ContactInformation.City)
	fmt.Println("Db Data : ", reqVal.ContactInformation.State)
	fmt.Println("Db Data : ", reqVal.ContactInformation.ZipCode)

	err3 := db.Exec(
		brandRegistrationQuery.StoreLocationQuery,
		brandDetails.LocationId,
		reqVal.ContactInformation.Address,
		reqVal.ContactInformation.City,
		reqVal.ContactInformation.State,
		reqVal.ContactInformation.ZipCode,
	).Error
	if err3 != nil {
		log.Error("Error in Updating the Brand Location Data: ", err3)
		return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	err4 := db.Exec(
		brandRegistrationQuery.StoreSocialMedia,
		brandDetails.SocialMediaId,
		reqVal.BrandInformation.Instragram,
		reqVal.BrandInformation.WebsiteURL,
	).Error
	if err4 != nil {
		log.Error("Error in Updating the Brand Social Media Data: ", err4)
		return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	err5 := db.Exec(
		brandRegistrationQuery.StoreDocumentQuery,
		brandDetails.DocumentId,
		reqVal.ContactInformation.ProofDocument,
		reqVal.TaxInformation.GstDocumant,
		reqVal.BrandInformation.BrandLogoPath,
		reqVal.TaxInformation.PanDocument,
		reqVal.ContactInformation.ProffDocumentName,
		reqVal.TaxInformation.PanDocumantName,
		reqVal.TaxInformation.GstDocumentName,
		reqVal.BrandInformation.BrandLogoName,
	).Error
	if err5 != nil {
		log.Error("Error in Updating the Brand Document Data: ", err5)
		return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	err6 := db.Exec(
		brandRegistrationQuery.StoreWareHouseQuery,
		brandDetails.WareHouseId,
		reqVal.WareHouseInfo.WareHouseAddress,
		reqVal.WareHouseInfo.WareHouseCity,
		reqVal.WareHouseInfo.WareHouseDistrict,
		reqVal.WareHouseInfo.WareHouse,
		reqVal.WareHouseInfo.WareHouseState,
		reqVal.WareHouseInfo.WareHouseZipCode,
	).Error
	if err6 != nil {
		log.Error("Error in Updating the Brand DocWare Houseument Data: ", err6)
		return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	return brandRegistrationModel.StoreBrandRegisterFormDetailsRes{
		Status:            true,
		Message:           "Brand registration data fetched successfully",
		ApplicationStatus: brandStatus,
	}
}
