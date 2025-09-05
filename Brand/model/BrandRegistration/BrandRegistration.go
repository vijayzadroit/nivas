package brandRegistrationModel

type GetBrandRegistrationRes struct {
	BrandInformation   BrandInformation   `json:"brandInformation"`
	ContactInformation ContactInformation `json:"contactInformation"`
	TaxInformation     TaxInformation     `json:"taxInformation"`
	WareHouseInfo      WareHouseInfo      `json:"wareHouseInfo"`
	ApplicationType    ApplicationType    `json:"applicationType"`
	Status             string             `json:"status"`
	Message            string             `json:"message"`
}

// type GetBrandRegisterFormDataResFromDb struct {
// 	BrandName           string `json:"refBrandName"`
// 	ProductCategory     string `json:"refProductCatageoryId"`
// 	BrandLogoPath       string `json:"refLogo"`
// 	BrandDescription    string `json:"refBrandDesciption"`
// 	WebsiteURL          string `json:"refWebsiteUrl"`
// 	Instragram          string `json:"refInstaUrl"`
// 	ContactPerson       string `json:"refBrandContactPerson"`
// 	Designation         string `json:"refDesignation"`
// 	PhoneNumber         string `json:"refBrandMobile"`
// 	Email               string `json:"refBrandEmail"`
// 	Address             string `json:"refAddress"`
// 	City                string `json:"refCity"`
// 	ZipCode             string `json:"refZipCode"`
// 	State               string `json:"refState"`
// 	ProofDocument       string `json:"refAddressProf"`
// 	GstinNumber         string `json:"refGstin"`
// 	CinNumber           string `json:"refCin"`
// 	GstDocumant         string `json:"refGstinPath"`
// 	PanDocument         string `json:"refPanCarsPath"`
// 	WareHouse           bool   `json:"refIfWareHouse"`
// 	WareHouseAddress    string `json:"wareHouseAddress"`
// 	WareHouseCity       string `json:"wareHouseCity"`
// 	WareHouseDistrict   string `json:"wareHouseDistrict"`
// 	WareHouseZipCode    string `json:"wareHouseZipCode"`
// 	WareHouseState      string `json:"wareHouseState"`
// 	SaveDraft           bool   `json:"refSaveDraft"`
// 	ApplicationStatusId int    `json:"refApplicationStatusId"`
// 	ApplicationStatus   string `json:"refApplicationStatus"`
// }

type GetBrandRegisterFormDataResFromDb struct {
	BrandName           string `json:"refBrandName" gorm:"column:refBrandName"`
	ProductCategory     int    `json:"refProductCatageoryId" gorm:"column:refProductCatageoryId"`
	BrandLogoPath       string `json:"refLogo" gorm:"column:refLogo"`
	BrandDescription    string `json:"refBrandDesciption" gorm:"column:refBrandDesciption"`
	WebsiteURL          string `json:"refWebsiteUrl" gorm:"column:refWebsiteUrl"`
	Instragram          string `json:"refInstaUrl" gorm:"column:refInstaUrl"`
	ContactPerson       string `json:"refBrandContactPerson" gorm:"column:refBrandContactPerson"`
	Designation         string `json:"refDesignation" gorm:"column:refDesignation"`
	PhoneNumber         string `json:"refBrandMobile" gorm:"column:refBrandMobile"`
	Email               string `json:"refBrandEmail" gorm:"column:refBrandEmail"`
	Address             string `json:"refAddress" gorm:"column:refAddress"`
	City                string `json:"refCity" gorm:"column:refCity"`
	ZipCode             string `json:"refZipCode" gorm:"column:refZipCode"`
	State               string `json:"refState" gorm:"column:refState"`
	ProofDocument       string `json:"refAddressProf" gorm:"column:refAddressProf"`
	GstinNumber         string `json:"refGstin" gorm:"column:refGstin"`
	CinNumber           string `json:"refCin" gorm:"column:refCin"`
	GstDocumant         string `json:"refGstinPath" gorm:"column:refGstinPath"`
	PanDocument         string `json:"refPanCars" gorm:"column:refPanCars"`
	WareHouse           bool   `json:"refIfWareHouse" gorm:"column:refIfWareHouse"`
	WareHouseAddress    string `json:"wareHouseAddress" gorm:"column:wareHouseAddress"`
	WareHouseCity       string `json:"wareHouseCity" gorm:"column:wareHouseCity"`
	WareHouseDistrict   string `json:"wareHouseDistrict" gorm:"column:wareHouseDistrict"`
	WareHouseZipCode    string `json:"wareHouseZipCode" gorm:"column:wareHouseZipCode"`
	WareHouseState      string `json:"wareHouseState" gorm:"column:wareHouseState"`
	SaveDraft           bool   `json:"refSaveDraft" gorm:"column:refSaveDraft"`
	ApplicationStatusId int    `json:"refApplicationStatusId" gorm:"column:refApplicationStatusId"`
	ApplicationStatus   string `json:"refApplicationStatus" gorm:"column:refApplicationStatus"`
	BrandLogoName       string `json:"refLogoName" gorm:"column:refLogoName"`
	ProffDocumentName   string `json:"refAddressProfName" gorm:"column:refAddressProfName"`
	GstDocumentName     string `json:"refGstinName" gorm:"column:refGstinName"`
	PanDocumantName     string `json:"refPanCarsName" gorm:"column:refPanCarsName"`
}

type GetBrandRegisterDataResponse struct {
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
	ApplicationType struct {
		SaveDraft              bool   `json:"saveDraft"`
		RefApplicationStatusId int    `json:"refApplicationStatusId"`
		RefApplicationStatus   string `json:"refApplicationStatus"`
	} `json:"applicationType"`
}

type GetBrandRegistrationDataFinalRes struct {
	BrandData GetBrandRegisterDataResponse `json:"brandData"`
	Status    bool                         `json:"status"`
	Message   string                       `json:"message"`
}

type BrandInformation struct {
	BrandName        string `json:"brandName"`
	ProductCategory  int    `json:"productCategory"`
	BrandLogoPath    string `json:"brandLogoPath"`
	BrandDescription string `json:"brandDescription"`
	WebsiteURL       string `json:"websiteURL"`
	Instragram       string `json:"instragram"`
}

type ContactInformation struct {
	ContactPerson string `json:"contactPerson"`
	Designation   string `json:"designation"`
	PhoneNumber   string `json:"phoneNumber"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	City          string `json:"city"`
	ZipCode       string `json:"zipCode"`
	State         string `json:"state"`
	ProofDocument string `json:"proofDocument"`
}

type TaxInformation struct {
	GstinNumber string `json:"gstinNumber"`
	CinNumber   string `json:"cinNumber"`
	GstDocumant string `json:"gstDocumant"`
	PanDocument string `json:"panDocument"`
}

type WareHouseInfo struct {
	WareHouse         bool   `json:"wareHouse"`
	WareHouseAddress  string `json:"wareHouseAddress"`
	WareHouseCity     string `json:"wareHouseCity"`
	WareHouseDistrict string `json:"wareHouseDistrict"`
	WareHouseZipCode  string `json:"wareHouseZipCode"`
	WareHouseState    string `json:"wareHouseState"`
}

type ApplicationType struct {
	SaveDraft bool `json:"saveDraft"`
	ReSubmit  bool `json:"reSubmit"`
}

type GetBrandRegistrationFromDataReq struct {
	ApplicationId string `json:"applicationId"`
}
