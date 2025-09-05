package brandRegistrationModel

type StoreBrandRegisterFormDataReq struct {
	BrandInformation struct {
		BrandName        string `json:"brandName"`
		ProductCategory  int    `json:"productCategory"`
		BrandLogoPath    string `json:"brandLogoPath"`
		BrandDescription string `json:"brandDescription"`
		WebsiteURL       string `json:"websiteURL"`
		Instragram       string `json:"instragram"`
		BrandLogoName    string `json:"brandLogaFileName"`
	} `json:"brandInformation"`
	ContactInformation struct {
		ContactPerson     string `json:"contactPerson"`
		Designation       string `json:"designation"`
		PhoneNumber       string `json:"phoneNumber"`
		Email             string `json:"email"`
		Address           string `json:"address"`
		City              string `json:"city"`
		ZipCode           string `json:"zipCode"`
		State             string `json:"state"`
		ProofDocument     string `json:"proofDocument"`
		ProffDocumentName string `json:"addressProofFileName"`
	} `json:"contactInformation"`
	TaxInformation struct {
		GstinNumber     string `json:"gstinNumber"`
		CinNumber       string `json:"cinNumber"`
		GstDocumant     string `json:"gstDocumant"`
		PanDocument     string `json:"panDocument"`
		GstDocumentName string `json:"gstDocumentFileName"`
		PanDocumantName string `json:"panDocumentFileName"`
	} `json:"taxInformation"`
	WareHouseInfo struct {
		WareHouse         bool   `json:"wareHouse"`
		WareHouseAddress  string `json:"wareHouseAddress"`
		WareHouseCity     string `json:"wareHouseCity"`
		WareHouseDistrict string `json:"wareHouseDistrict"`
		WareHouseZipCode  string `json:"wareHouseZipCode"`
		WareHouseState    string `json:"wareHouseState"`
	} `json:"wareHouseInfo"`
	ApplicationId string `json:"applicationId"`
	SaveDraft     bool   `json:"saveDraft"`
}

type BrandReferenceIdRes struct {
	SocialMediaId int `json:"refSocialMediaId" gorm:"column:refSocialMediaId"`
	DocumentId    int `json:"refDocumentsId" gorm:"column:refDocumentsId"`
	LocationId    int `json:"refBrandLocationId" gorm:"column:refBrandLocationId"`
	ContactId     int `json:"refBrandContactId" gorm:"column:refBrandContactId"`
	WareHouseId   int `json:"refWareHoueId" gorm:"column:refWareHoueId"`
}

type StoreBrandRegisterFormDetailsRes struct {
	Status            bool   `json:"status"`
	Message           string `json:"message"`
	ApplicationStatus int    `json:"applicationStatus"`
}
