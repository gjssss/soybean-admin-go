package system

type UploadDTO struct {
	Url         string `json:"url" binding:"required"`
	ObjectKey   string `json:"objectKey" binding:"required"`
	ContentType string `json:"contentType" binding:"required"`
}
