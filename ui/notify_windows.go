//go:build windows

package ui

import (
	"cherry/utils"
	"cherry/utils/conf"
	"github.com/go-toast/toast"
	"log"
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
	notification := toast.Notification{
		AppID:   "Cherry",
		Title:   "图片成功失败",
		Message: errorMsg,
	}

	err := notification.Push()
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

	notification := toast.Notification{
		AppID:   "Cherry",
		Title:   "图片成功成功",
		Message: strings.Join(urlResult, ", ") + clipboardMsg,
	}

	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
