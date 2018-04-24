package chat

import (
	//"fmt"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"

)

//创建模块
var Module = func() module.Module {
	chat := new(Chat)
	return chat
}

//模块定义
type Chat struct{
	basemodule.BaseModule
}


func (m *Chat) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Chat"
}

func (m *Chat) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

//初始化
func (m *Chat) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
	//注册远程调用/消息
	m.GetServer().Register("HD_Chess", m.chat)
}

//返回一次
func (m *Chat) Run(closeSig chan bool) {
}


func (m *Chat) OnDestroy() {
	//一定别忘了关闭RPC
	m.GetServer().OnDestroy()
}

//消息回调
func (m *Chat) chat(session gate.Session,msg map[string]interface{}) (result string, err string) {
	//time.Sleep(1)
	return "sss",""
}