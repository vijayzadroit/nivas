package RegisterFormDocumentModel

type DocumentDeleteReq struct {
	OldFile string `json: "oldFile"`
}

type DocumentDeleteResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
