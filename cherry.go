package main

import (
	"cherry/server"
	"cherry/ui"
	"cherry/utils/conf"
)

func main() {
	conf.LoadCherryConfig()

	go server.StartCherryServer()

	ui.InitSysTray()
}
