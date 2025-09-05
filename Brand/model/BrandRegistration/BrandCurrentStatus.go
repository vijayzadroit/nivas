package brandRegistrationModel

type GetBrandCurrentStatusReq struct {
	ApplicationId string `json:"applicationId"`
}

type GetBrandStatusFromDbRes struct {
	ApplicationStatus int `json:"refApplicationStatus" gorm:"column:refApplicationStatus"`
}

type GetBrandCurrentStatusRes struct {
	AppStatus int    `json:"applicationStatus`
	Status    bool   `json:"status"`
	Message   string `json:"message"`
}
