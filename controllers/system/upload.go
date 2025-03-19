package system

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
	"github.com/gjssss/soybean-admin-go/utils/upload"
)

// UploadController 处理上传相关的请求
type UploadController struct{}

// GetUploadToken 获取S3上传凭证
// @Summary 获取S3上传凭证
// @Description 获取S3预签名URL用于直接上传文件
// @Tags 上传
// @Accept json
// @Produce json
// @Param object_key query string true "对象键名"
// @Param content_type query string true "文件内容类型"
// @Success 200 {object} utils.Response "返回预签名URL"
// @Router /upload/aws [get]
func (uc *UploadController) GetUploadToken(c *gin.Context) {
	objectKey := c.Query("object_key")
	contentType := c.Query("content_type")

	// 如果没有提供object_key，则返回错误
	if objectKey == "" {
		utils.Fail(c, http.StatusBadRequest, "对象键名不能为空")
		return
	}

	// 如果没有提供content_type，则尝试根据文件扩展名猜测
	if contentType == "" {
		ext := filepath.Ext(objectKey)
		switch ext {
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		case ".gif":
			contentType = "image/gif"
		case ".pdf":
			contentType = "application/pdf"
		case ".doc", ".docx":
			contentType = "application/msword"
		case ".xls", ".xlsx":
			contentType = "application/vnd.ms-excel"
		default:
			contentType = "application/octet-stream"
		}
	}

	// 生成有效期为15分钟的预签名URL
	presignedURL, err := upload.GeneratePresignedURL(objectKey, contentType, 1*time.Minute)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "生成上传凭证失败: "+err.Error())
		return
	}

	utils.Success(c, system.UploadDTO{
		Url:         presignedURL,
		ObjectKey:   objectKey,
		ContentType: contentType,
	})
}
