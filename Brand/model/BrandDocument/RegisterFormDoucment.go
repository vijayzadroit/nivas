package RegisterFormDocumentModel

type UploadUrlRequestData struct {
	FileName  string `json:"fileName"`
	BrandName string `json:"brandName"`
}

type UploadUrlResponse struct {
	UploadUrl string `json:"uploadUrl"`
	FileUrl   string `json:"fileUrl"`
	FilePath  string `json:"filePath`
	Status    bool   `json:"status"`
	Message   string `json:"message"`
}

