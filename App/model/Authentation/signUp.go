package model

type SignUpResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Profile string `json:"profileUrl"`
}

type GoogleToken struct {
	GoogleToken string `json:"googleToken" binding:"required`
}

type CountryInfo struct {
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	DialCode    string `json:"dial_code"`
	NSNMin      int    `json:"nsn_min"`
	NSNMax      int    `json:"nsn_max"`
}

type CountryInfoResponse struct {
	Status      bool          `json:"status"`
	Message     string        `json:"message"`
	CountryData []CountryInfo `json:"countryData"`
}

type MobileNumberValidationRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required"`
}

type MobileNumberValidationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int64    `json:"code"`
}
