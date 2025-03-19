package upload

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gjssss/soybean-admin-go/global"
)

// GetS3Client returns an S3 client
func GetS3Client() (*s3.Client, error) {
	s3Config := global.Config.S3

	// 根据配置创建AWS凭证提供者
	credProvider := credentials.NewStaticCredentialsProvider(
		s3Config.AccessKey,
		s3Config.SecretKey,
		"",
	)

	// 创建AWS配置
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credProvider),
		config.WithRegion(s3Config.Region),
	)
	if err != nil {
		return nil, err
	}

	// 创建S3客户端，使用自定义端点
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(s3Config.Endpoint)
		o.UsePathStyle = true
	})

	return client, nil
}

// GeneratePresignedURL 生成预签名URL用于上传文件
func GeneratePresignedURL(objectKey string, contentType string, expires time.Duration) (string, error) {
	client, err := GetS3Client()
	if err != nil {
		return "", err
	}

	presignClient := s3.NewPresignClient(client)

	putObjectInput := &s3.PutObjectInput{
		Bucket:      aws.String(global.Config.S3.Bucket),
		Key:         aws.String(objectKey),
		ContentType: aws.String(contentType),
	}

	presignedReq, err := presignClient.PresignPutObject(context.TODO(), putObjectInput, func(opts *s3.PresignOptions) {
		opts.Expires = expires
	})
	if err != nil {
		return "", err
	}

	return presignedReq.URL, nil
}
