package ui

import (
	"cherry/log"
	"cherry/utils/conf"
	"github.com/gen2brain/beeep"
)

func ShowNotify(notifyMsg string) {
	if !conf.GetCherryConfig().Server.ShowSysNotify {
		return
	}

	err := beeep.Notify("Cherry", notifyMsg, "")
	if err != nil {
		log.E("通知显示异常: " + notifyMsg)
		return
	}
}
