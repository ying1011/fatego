package main

import (
	//"fmt"
	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/module/modules"
	"server/gate"
	"server/game"
	"server/login"
)

func main()  {
	app:= mqant.CreateApp()

	app.Run(true,
			modules.MasterModule(),
			gate.Module(),  //这是默认网关模块,是必须的支持 TCP,websocket,MQTT协议
			game.Module(),
			login.Module(),
		)
}