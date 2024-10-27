package tray

import (
	"cherry/log"
	"cherry/utils"
	"cherry/utils/conf"
	"fmt"
	"github.com/getlantern/systray"
	"os"
)

func InitSysTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(CherryIcon)
	systray.SetTitle("Cherry")

	systray.AddMenuItem("程序运行中...", "程序运行中...")
	systray.AddSeparator()

	mReloadConfig := systray.AddMenuItem("重载图床配置", "")
	mOpenDir := systray.AddMenuItem("打开程序目录", "打开程序目录")
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "退出程序")
	// Sets the icon of a menu item. Only available on Mac and Windows.
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	for {
		select {
		case <-mOpenDir.ClickedCh:
			utils.OpenSysDirectory(conf.GetRunnerPath())
		case <-mReloadConfig.ClickedCh:
			log.I("重载配置")
			conf.LoadCherryConfig()
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}

}

func onExit() {
	// Cleanly shutdown the systray
	log.I("系统退出")
	os.Exit(0)
}
