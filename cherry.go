package main

import (
	"cherry/server"
	"cherry/tray"
	"cherry/utils/conf"
)

func main() {
	conf.LoadCherryConfig()

	go server.StartCherryServer()

	tray.InitSysTray()
}
