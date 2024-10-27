package uploader

import (
	"cherry/uploader/s3"
	"cherry/utils"
	"errors"
)

//type Uploader interface {
//	Upload(uploaderName string, data []byte) (string, error)
//}

type UploaderName string

const (
	S3 UploaderName = "aws-s3"
)

func Upload(name UploaderName, date []byte) (string, error) {
	format := utils.DetectImageFormat(date)

	if name == S3 {
		return s3.Upload(date, format)
	}
	return "", errors.New("not support uploader")
}
