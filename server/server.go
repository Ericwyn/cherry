package server

import (
	"cherry/log"
	"cherry/uploader"
	"cherry/utils"
	"cherry/utils/conf"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func StartCherryServer() {

	config := conf.GetCherryConfig()

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		var reqBody RequestBody
		var imgData []byte
		var uploadUrl string

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			log.D("No JSON provided, trying clipboard...")
			imgData, err = utils.GetClipboardImageData()
			if err == nil && imgData != nil {
				//SaveImage(img, format)
				uploadUrl, err = uploader.Upload(uploader.S3, imgData)
				if err != nil {
					log.E("upload err, ", err.Error())
					c.JSON(400, failResp(err.Error()))
				} else {
					c.JSON(200, successResp([]string{uploadUrl}))
				}
				return
			} else {
				c.JSON(400, failResp("Failed to get image from clipboard"))
				return
			}
		}

		var resultList []string
		for _, filePath := range reqBody.List {
			localFileByte, err := os.ReadFile(filePath)
			if err != nil {
				c.JSON(400, failResp(err.Error()))
				return
			}
			uploadUrl, err = uploader.Upload(uploader.S3, localFileByte)
			if err != nil {
				log.E("upload err, ", err.Error())
				c.JSON(400, failResp(err.Error()))
			} else {
				//c.JSON(200, successResp(true, []string{uploadUrl}))
				resultList = append(resultList, uploadUrl)
			}
		}

		c.JSON(200, successResp(resultList))
	})

	err := router.Run(":" + strconv.Itoa(config.Server.Port))
	if err != nil {
		log.E("启动服务失败: ", err.Error())
		os.Exit(-1)
		return
	}

}
