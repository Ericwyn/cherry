package ui

import (
	"cherry/log"
	"cherry/uploader"
	"cherry/utils"
	"cherry/utils/conf"
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
	mVersion := systray.AddMenuItem("版本: "+conf.Version, "版本: "+conf.Version)
	mUploadFromClipboard := systray.AddMenuItem("剪贴板图片上传", "剪贴板图片上传")
	systray.AddSeparator()

	mReloadConfig := systray.AddMenuItem("重载图床配置", "")
	mOpenDir := systray.AddMenuItem("打开程序目录", "打开程序目录")
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "退出程序")

	for {
		select {
		case <-mVersion.ClickedCh:
			log.I("当前版本: ", conf.Version)
		case <-mUploadFromClipboard.ClickedCh:
			//utils.UploadFromClipboard()
			uploadUrl, err := uploader.UploadFromClipboard(uploader.S3)
			if err != nil {
				log.E("剪贴板上传失败 ", err.Error())
				ShowErrResultNotify("剪贴板读取失败: " + err.Error())
			} else {
				ShowSuccessResultNotify([]string{uploadUrl})
				FlushResultClipboard(uploadUrl)
			}
		case <-mOpenDir.ClickedCh:
			utils.OpenSysDirectory(conf.GetRunnerPath())
		case <-mReloadConfig.ClickedCh:
			log.I("重载配置")
			conf.LoadCherryConfig()
		case <-mQuit.ClickedCh:
			systray.Quit()
		}
	}

}

func onExit() {
	// Cleanly shutdown the systray
	log.I("系统退出")
	os.Exit(0)
}
