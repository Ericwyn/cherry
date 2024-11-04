package conf

import (
	"cherry/log"
	"path"
	"path/filepath"
	"strings"
)

const configFileName = "cherry-setting.json"

var runnerPath = ""

// GetRunnerPath
// 获取项目所在路径, 就是我们正常理解里面的 ./
func GetRunnerPath() string {
	if runnerPath == "" {
		dir, err := filepath.Abs(filepath.Dir("./"))
		if err != nil {
			log.E("获取 ./ 目录绝对路径失败")
			log.E(err)
		}

		runnerPath = strings.Replace(dir, "\\", "/", -1)
		log.D("程序运行目录更新为: " + runnerPath)
	}

	return runnerPath
}

func GetRunnerConfigPath() string {
	return path.Join(GetRunnerPath(), configFileName)
}
