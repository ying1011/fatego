package main

import (
	//"fmt"
	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/module/modules"
	"server/gate"
	"server/webapp"
)

func main()  {
	app:= mqant.CreateApp()

	app.Run(true,
		modules.MasterModule(),
		gateAlpha.Module(),  //这是默认网关模块,是必须的支持 TCP,websocket,MQTT协议
		//game.Module(),
		//login.Module(),
		webapp.Module(),
	)
}