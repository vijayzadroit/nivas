package brandRegistrationModel

type GetBrandRegistrationStatusReq struct {
	ApplicationId string `json:"applicationId"`
}

type BrandApplicationResponse struct {
	ApplicationStatus            string `json:"applicationStatus" gorm:"column:refApplicationStatus"`
	ApplicationStatusId          int    `json:"applicationStatusId" gorm:"column:refApplicationStatusId"`
	ApplicationCustId            string `json:"applicationCustId" gorm:"column:refApplicationCustId"`
	ProcessTime                  string `json:"processTime" gorm:"column:refProcessTime"`
	Logo                         string `json:"logo" gorm:"column:refLogo"`
	AddressProf                  string `json:"addressProf" gorm:"column:refAddressProf"`
	PanDoc                       string `json:"panDoc" gorm:"column:refPanDoc"`
	GstinDoc                     string `json:"gstinDoc" gorm:"column:refGstinDoc"`
	BrandName                    string `json:"brandName" gorm:"column:refBrandName"`
	BrandCategoryName            string `json:"brandCategoryName" gorm:"column:refBrandCategoryName"`
	Cin                          string `json:"cin" gorm:"column:refCin"`
	Gstin                        string `json:"gstin" gorm:"column:refGstin"`
	BrandContactPerson           string `json:"brandContactPerson" gorm:"column:refBrandContactPerson"`
	BrandMobile                  string `json:"brandMobile" gorm:"column:refBrandMobile"`
	BrandEmail                   string `json:"brandEmail" gorm:"column:refBrandEmail"`
	CreateAt                     string `json:"createAt" gorm:"column:refCreateAT"`
	UpdateAt                     string `json:"updateAt" gorm:"column:refUpdateAt"`
	CurrentStatus                string `json:"currentStatus" gorm:"column:currentStatus"`
	ApplicationStatusDescription string `json:"applicationStatusDescription" gorm:"column:refApplicationStatusDesciption"`
	AddressProfName              string `json:"refAddressProfName" gorm:"column:refAddressProfName"`
	GstDocName                   string `json:"refGstinName" gorm:"column:refGstinName"`
	PanDocName                   string `json:"refPanCarsName" gorm:"column:refPanCarsName"`
	LogoName                     string `json:"refLogoName" gorm:"column:refLogoName"`
}

type DocumentLink struct {
	Url         string `json:"url"`
	DownloadUrl string `json:"downloadUrl"`
	FileSize    string `json:"fileSize"`
}

type BrandApplicationFormatted struct {
	BrandApplicationStatus struct {
		Status        int    `json:"status"`
		StatusName    string `json:"statusName"`
		ApplicationId string `json:"applicationId"`
		SubmitDate    string `json:"submitDate"`
		LastDate      string `json:"lastDate"`
		ProcessTime   string `json:"processTime"`
	} `json:"brandApplicationStatus"`

	BrandInformation struct {
		BrandLogo     string `json:"brandLogo"`
		BrandName     string `json:"brandName"`
		BrandCategory string `json:"brandCategory"`
		CinNumber     string `json:"cinNumber"`
		ContactPerson string `json:"contactPerson"`
		SubmitDate    string `json:"submitDate"`
		PhoneNumber   string `json:"phoneNumber"`
		Email         string `json:"email"`
	} `json:"brandInformation"`

	Document struct {
		ShowDocument bool         `json:"showDocument"`
		AddressProof DocumentLink `json:"addressProof"`
		GstDocument  DocumentLink `json:"gstDocument"`
		PanDocument  DocumentLink `json:"panDocument"`
	} `json:"document"`

	Feedback struct {
		CurrentStatus string `json:"currentStatus"`
		ReviewContent string `json:"reviewContent"`
	} `json:"feedback"`
}

type BrandRegistrationStatus struct {
	Status    bool                      `json:"status"`
	Message   string                    `json:"message"`
	BrandData BrandApplicationFormatted `json:"brandData"`
}
