package system

type UploadAWSDTO struct {
	Url         string `json:"url" binding:"required"`
	ObjectKey   string `json:"objectKey" binding:"required"`
	ContentType string `json:"contentType" binding:"required"`
}

type UploadQiniuDTO struct {
	Token     string `json:"token" binding:"required"`
	Key       string `json:"key" binding:"required"`
	UploadUrl string `json:"uploadUrl" binding:"required"`
	Url       string `json:"url" binding:"required"`
}
