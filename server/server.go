package server

import (
	"cherry/log"
	"cherry/ui"
	"cherry/uploader"
	"cherry/utils/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
)

func StartCherryServer() {

	config := conf.GetCherryConfig()

	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to your specific domain if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/upload", func(c *gin.Context) {
		var reqBody RequestBody
		var uploadUrl string

		err := c.ShouldBindJSON(&reqBody)

		if err == nil && len(reqBody.List) != 0 {
			resultList, err := uploader.UploadFromLocalFile(uploader.S3, reqBody.List)
			if err != nil {
				log.E("upload err, ", err.Error())
				c.JSON(400, failResp(err.Error()))
			} else {
				resultList = append(resultList, uploadUrl)
			}

			if len(resultList) != 0 {
				ui.ShowNotify(strconv.Itoa(len(resultList)) + " 张图片上传成功: " + strings.Join(resultList, ", "))
			}
			c.JSON(200, successResp(resultList))
			return
		} else {
			log.D("参数解析失败, 剪贴板图片上传")

			uploadUrl, err = uploader.UploadFromClipboard(uploader.S3)
			if err != nil {
				log.E("upload err, ", err.Error())
				c.JSON(400, failResp(err.Error()))
			} else {
				ui.ShowNotify("1 张图片上传成功: " + uploadUrl)
				c.JSON(200, successResp([]string{uploadUrl}))
			}
			return
		}
	})

	err := router.Run(":" + strconv.Itoa(config.Server.Port))
	if err != nil {
		log.E("启动服务失败: ", err.Error())
		os.Exit(0)
		return
	}
}
