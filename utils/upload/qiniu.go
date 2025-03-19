package upload

import (
	"fmt"

	"github.com/gjssss/soybean-admin-go/utils/config"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

// GenerateQiniuSingleUseToken 生成七牛云单次使用的上传凭证
// 该凭证使用一次后失效
func GenerateQiniuSingleUseToken(qiniuConfig *config.QiniuConfig, fileKey string) (string, error) {
	// 创建上传策略
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", qiniuConfig.Bucket, fileKey), // 指定文件key，确保只能上传到这个路径
		Expires:    3600,                                              // 凭证有效期1小时
		InsertOnly: 1,                                                 // 1表示只能新增，不能覆盖，这确保了凭证只能使用一次
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(fname)","mimeType":"$(mimeType)"}`,
	}

	// 创建七牛云认证对象
	mac := auth.New(qiniuConfig.AccessKey, qiniuConfig.SecretKey)

	// 生成上传凭证
	upToken := putPolicy.UploadToken(mac)

	return upToken, nil
}

// GenerateQiniuUploadCredentials 生成完整的七牛云上传凭证信息
func GenerateQiniuUploadCredentials(qiniuConfig *config.QiniuConfig, fileKey, mimeType string) (map[string]string, error) {
	token, err := GenerateQiniuSingleUseToken(qiniuConfig, fileKey)
	if err != nil {
		return nil, err
	}

	// 构建上传信息
	credentials := map[string]string{
		"token":       token,
		"key":         fileKey,
		"uploadUrl":   qiniuConfig.Endpoint, // 根据您的存储区域可能需要修改
		"mimeType":    mimeType,
		"resourceUrl": fmt.Sprintf("%s/%s", qiniuConfig.PublicDomain, fileKey),
	}

	return credentials, nil
}
