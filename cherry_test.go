package main

import (
	"cherry/log"
	"cherry/ui"
	"cherry/uploader/s3"
	"cherry/utils/conf"
	"testing"
)

func TestUpload(t *testing.T) {
	conf.GetCherryConfig().S3.UploadPath = "typora/cherry_test.{extName}"
	uploadUrl, err := s3.Upload(ui.CherryIcon, "ico")
	if err != nil {
		log.E("upload err, ", err.Error())
	} else {
		log.I("upload success, ", uploadUrl)
	}
}
