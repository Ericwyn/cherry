package uploader

import (
	"cherry/log"
	"cherry/uploader/s3"
	"cherry/utils"
	"cherry/utils/conf"
	"errors"
	"os"
)

//type Uploader interface {
//	Upload(uploaderName string, data []byte) (string, error)
//}

type UploaderName string

const (
	S3 UploaderName = "aws-s3"
)

func UploadFromLocalFile(name UploaderName, filePaths []string) ([]string, error) {
	log.I("开始本地图片上传")

	var resultList []string

	for _, filePath := range filePaths {
		log.D("本地文件上传: " + filePath)
		localFileByte, err := os.ReadFile(filePath)
		if err != nil {
			log.E("本地文件读取失败 ", err.Error())
			continue
		}
		uploadUrl, err := Upload(name, localFileByte)
		if err != nil {
			log.E("本地文件上传失败 ", err.Error())
			continue
		}
		resultList = append(resultList, uploadUrl)
	}

	return resultList, nil
}

func UploadFromClipboard(name UploaderName) (string, error) {
	log.I("开始剪贴板图片上传")
	var imgData []byte
	var err error

	imgData, err = utils.GetClipboardImageData()
	if err != nil || imgData == nil || len(imgData) == 0 {
		if err != nil {
			log.E("剪贴板图片读取失败, ", err.Error())
			return "剪贴板图片读取失败", err
		} else {
			return "剪贴板图片读取失败", errors.New("剪贴板图片读取失败")
		}
	}
	upload, err := Upload(name, imgData)

	if err == nil {
		if conf.GetCherryConfig().Server.ShowSysNotify {
			utils.WriteUrlToClipboard(upload)
		}
	}
	return upload, err
}

func Upload(name UploaderName, date []byte) (string, error) {
	format := utils.DetectImageFormat(date)

	if name == S3 {
		return s3.Upload(date, format)
	}
	return "", errors.New("not support uploader")
}
