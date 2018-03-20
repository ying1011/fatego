package game

import (
	"fmt"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
)

//创建模块
var Module = func() module.Module {
	game := new(Game)
	return game
}

//模块定义
type Game struct{
	basemodule.BaseModule
}


func (m *Game) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Game"
}

func (m *Game) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

//初始化
func (m *Game) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
	//注册远程调用/消息
	m.GetServer().Register("HD_Game", m.game)
}


func (m *Game) Run(closeSig chan bool) {
}


func (m *Game) OnDestroy() {
	//一定别忘了关闭RPC
	m.GetServer().OnDestroy()
}

//消息回调
func (m *Game) game(session gate.Session,msg map[string]interface{}) (result string, err string) {
	//time.Sleep(1)
	if msg["say"] == nil{
		result = "say cannot be nil"
		return
	}
	say := msg["say"].(string)
	fmt.Println(say)

	return "sss",""
}