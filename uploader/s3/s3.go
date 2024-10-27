package s3

import (
	"bytes"
	"cherry/log"
	"cherry/utils/conf"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/url"
	"strings"
	"time"
)

type s3Uploader struct {
	Bucket string
	Client *s3.S3
}

func newS3Uploader(endpoint, accessKeyID, secretAccessKey, bucket, region string) (*s3Uploader, error) {
	// 创建一个新的会话
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return nil, err
	}

	return &s3Uploader{
		Bucket: bucket,
		Client: s3.New(sess),
	}, nil
}

var uploader *s3Uploader

// Upload 上传接口
func Upload(picData []byte, format string) (string, error) {
	cherryConfig := conf.GetCherryConfig()
	if uploader == nil {
		var err error
		uploader, err = newS3Uploader(
			cherryConfig.S3.Endpoint,
			cherryConfig.S3.AccessKeyID,
			cherryConfig.S3.SecretAccessKey,
			cherryConfig.S3.BucketName,
			cherryConfig.S3.Region,
		)
		if err != nil || uploader == nil {
			return "", errors.New("create s3 uploader err")
		}
	}

	filePath := getFilePath(format)

	_, err := uploader.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(uploader.Bucket),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader(picData),
	})
	if err != nil {
		return "", err
	}
	respUrl := getResponseUrl(filePath)
	log.I("上传图片到: " + respUrl)
	return respUrl, nil
}

func getFilePath(format string) string {
	cherryConfig := conf.GetCherryConfig()

	filePath := strings.Replace(cherryConfig.S3.UploadPath, "{timestampMS}",
		fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)), 1)
	return strings.Replace(filePath, "{extName}", format, 1)
}

func getResponseUrl(filePath string) string {
	cherryConfig := conf.GetCherryConfig()

	var res string
	var err error
	if !cherryConfig.S3.DisableBucketPrefixToURL {
		res, err = url.JoinPath(
			cherryConfig.S3.URLPrefix,
			cherryConfig.S3.BucketName,
			filePath,
		)
	} else {
		res, err = url.JoinPath(cherryConfig.S3.URLPrefix, filePath)
	}
	if err != nil {
		log.E("build file url error, ", err.Error())
		return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", cherryConfig.S3.BucketName, filePath)
	} else {
		return res
	}
}
