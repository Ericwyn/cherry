package conf

import (
	"cherry/log"
	"encoding/json"
	"os"
)

const Version = "0.0.2"

type CherryConfig struct {
	Server Server `json:"server"`
	S3     S3     `json:"s3"`
}
type Server struct {
	Port int `json:"port"`
	//Host   string `json:"host"`
	//Enable bool   `json:"enable"`
}
type S3 struct {
	AccessKeyID              string `json:"accessKeyID"`
	SecretAccessKey          string `json:"secretAccessKey"`
	BucketName               string `json:"bucketName"`
	UploadPath               string `json:"uploadPath"`
	Region                   string `json:"region"`
	Endpoint                 string `json:"endpoint"`
	URLPrefix                string `json:"urlPrefix"`
	RejectUnauthorized       bool   `json:"rejectUnauthorized"`
	DisableBucketPrefixToURL bool   `json:"disableBucketPrefixToURL"`
}

var cherryConfig *CherryConfig

func GetCherryConfig() *CherryConfig {
	if cherryConfig != nil {
		return cherryConfig
	}

	LoadCherryConfig()

	checkCherryConfig()

	return cherryConfig
}

func LoadCherryConfig() {
	path := GetRunnerConfigPath()
	log.I("读取配置: " + path)

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.E("配置文件读取失败 ", err.Error())
		os.Exit(-1)
	}

	config := &CherryConfig{}
	err = json.Unmarshal(fileBytes, config)
	if err != nil {
		log.E("配置文件解析失败 ", err.Error())
		os.Exit(-1)
	}

	cherryConfig = config
}

func checkCherryConfig() {
	if cherryConfig == nil {
		log.E("配置校验失败: 没有配置文件")
		os.Exit(-1)
	}
	if cherryConfig.Server.Port <= 1000 {
		log.E("配置校验失败: Server.Port <= 1000, port: ", cherryConfig.Server.Port)
		os.Exit(-1)
	}

	if cherryConfig.S3.BucketName == "" {
		log.E("配置校验失败: S3.BucketName 为空")
		os.Exit(-1)
	}

	if cherryConfig.S3.SecretAccessKey == "" {
		log.E("配置校验失败: S3.SecretAccessKey 为空")
		os.Exit(-1)
	}

	if cherryConfig.S3.AccessKeyID == "" {
		log.E("配置校验失败: S3.AccessKeyID 为空")
		os.Exit(-1)
	}
}
