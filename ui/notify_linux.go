//go:build linux

package ui

import (
	"cherry/utils"
	"cherry/utils/conf"
	"log"
	"os/exec"
	"strings"
)

func FlushResultClipboard(result string) {
	if !conf.GetCherryConfig().Server.FlushResultClipboard {
		return
	}
	utils.WriteUrlToClipboard(result)
}

func ShowErrResultNotify(errorMsg string) {
	if !conf.GetCherryConfig().Server.ShowSysNotify {
		return
	}
	cmd := exec.Command("notify-send", "Cherry", "图片发送失败: "+errorMsg)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func ShowSuccessResultNotify(urlResult []string) {
	if !conf.GetCherryConfig().Server.ShowSysNotify {
		return
	}
	clipboardMsg := ""
	if conf.GetCherryConfig().Server.FlushResultClipboard {
		clipboardMsg = " ，已复制到剪贴板"
	}

	message := strings.Join(urlResult, ", ") + clipboardMsg
	cmd := exec.Command("notify-send", "Cherry", "图片发送成功: "+message)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
