package service

import (
	"fmt"
	"os"

	getDate "nivasBackendMain/Helper/DateFormat"
	logger "nivasBackendMain/Helper/Logger"
	mailService "nivasBackendMain/Helper/MailService"
	model "nivasBackendMain/NivasAdmin/model/GenerateBrandRegistrationUrl"
	query "nivasBackendMain/NivasAdmin/query/GenerateBranchRequestUrl"

	"gorm.io/gorm"
)

func BrandRegistrationUrlService(db *gorm.DB, reqVal model.BrandRegistrationUrlRequest) model.BrandEmailCheckRes {
	log := logger.InitLogger()
	webSiteUrl := os.Getenv("BRAND_REGISTRATION_URL")

	var brandDetails []model.BrandMailCheckQueryResponse
	// 1. Check if brand email already exists
	err := db.Raw(query.CheckBrandEmail, reqVal.MailId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Brand Email Unique Check Query Failed: " + err.Error())
		return model.BrandEmailCheckRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	var applicationId model.BrandApplicationIdResponse
	var brandEmail, brandName, brandOwnerName, Message string

	if len(brandDetails) != 0 {
		fmt.Println("Brand Details:", brandDetails)

		applicationId.ApplicationCustId = brandDetails[0].ApplicationCustId
		brandName = brandDetails[0].BrandName
		brandOwnerName = brandDetails[0].BrandName // (maybe should be BrandOwnerName if you have that field later?)
		brandEmail = brandDetails[0].MailId
		Message = "Brand Already Registered, Mail Resend To that Brand"
	}

	// 2. Insert new brand application
	if len(brandDetails) == 0 {
		var date = getDate.GetCurrentDate("")
		fmt.Println(date)
		err := db.Raw(query.CreateNewApplication, reqVal.BrandName, reqVal.MailId, date).
			Scan(&applicationId).Error
		if err != nil {
			log.Error("Create New Application Insert Failed: " + err.Error())
			return model.BrandEmailCheckRes{
				Status:  false,
				Message: "Failed to create brand application",
			}
		}
		brandEmail = reqVal.MailId
		brandName = reqVal.BrandName
		brandOwnerName = reqVal.CustomerName
		Message = "Brand application created"
	}

	registrationURL := webSiteUrl + "/#id:" + applicationId.ApplicationCustId
	subject := fmt.Sprintf("Welcome %s - Complete Your Registration", brandName)

	htmlContent := fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; line-height: 1.6;">
				<h2>Hello %s,</h2>
				<p>We’re excited to have <b>%s</b> on board!</p>
				<p>Please complete your registration by clicking the link below:</p>
				<p>
					<a href="%s" 
					   style="display:inline-block; padding:10px 20px; background:#4CAF50; 
							  color:#fff; text-decoration:none; border-radius:5px;">
						Complete Registration
					</a>
				</p>
				<p>If the button doesn’t work, copy and paste this link into your browser:</p>
				<p><a href="%s">%s</a></p>
				<br>
				<p>Best Regards,<br>Brand Support Team</p>
			</body>
		</html>
	`, brandOwnerName, brandName, registrationURL, registrationURL, registrationURL)

	fmt.Println("brandEmail : " + brandEmail)

	// 4. Send mail
	mailSent := mailService.MailService(reqVal.MailId, htmlContent, subject)
	if !mailSent {
		log.Error("Mail sending failed for brand: " + brandName)
		return model.BrandEmailCheckRes{
			Status:  false,
			Message: Message + "but mail sending failed",
		}
	}

	log.Info("Brand application created and mail sent successfully to: " + brandEmail)
	return model.BrandEmailCheckRes{
		Status:  true,
		Message: Message + "successfully",
	}
}
