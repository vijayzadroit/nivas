package model

type BrandRegistrationUrlRequest struct {
	MailId       string `json:"mailId" binding :"required"`
	BrandName    string `json:"brandName" binding :"required"`
	CustomerName string `json:"customerName" binding :"required"`
}

type BrandEmailCheckReq struct {
	Mail string `json:"brandEmail"`
}

type BrandMailCheckQueryResponse struct {
	ApplicationCustId string `json:"refApplicationCustId" gorm:"column:refApplicationCustId"`
	BrandName         string `json:"refBrandName" gorm:"column:refBrandName"`
	MailId            string `json:"refMailId" gorm:"column:refMailId"`
}

type BrandEmailCheckRes struct {
	ApplicationId     int    `json:"refApplicationId" gorm:"refApplicationId"`
	BrandId           string `json:"refBrandId" gorm:"refBrandId"`
	ApplicationCustId string `json:"refApplicationCustId" gorm:"refApplicationCustId"`
	BrandName         string `json:"refBrandName" gorm:"refBrandName"`
	MailId            string `json:"refMailId" gorm:"refMailId"`
	Status            bool   `json:"status" binding : "required"`
	Message           string `json:"message" binding :"required"`
}

type BrandApplicationIdResponse struct {
	ApplicationCustId string `json:"refApplicationCustId" gorm:"column:refApplicationCustId"`
}
type SocialMediaRes struct {
	SocialMediaId string `json:"refSocialMediaId" gorm:"column:refSocialMediaId"`
}
type ContactRes struct {
	ContactId string `json:"refBrandContactId" gorm:"column:refBrandContactId"`
}
type LocationRes struct {
	LocationId string `json:"refLocationId" gorm:"column:refLocationId"`
}
type WareHouseRes struct {
	WareHouseId string `json:"refWareHouseId" gorm:"column:refWareHouseId"`
}
type DocumentRes struct {
	DocumentId string `json:"refDocumentsId" gorm:"column:refDocumentsId"`
}
