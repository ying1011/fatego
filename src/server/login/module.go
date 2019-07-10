package login

import (
	"fmt"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
)

//创建模块
var Module = func() module.Module {
	login := new(Login)
	return login
}

//模块定义
type Login struct{
	basemodule.BaseModule
}


func (m *Login) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Login"
}

func (m *Login) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

//初始化
func (m *Login) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
	//注册远程调用/消息
	m.GetServer().Register("HD_Login", m.login)


}

//返回一次
func (m *Login) Run(closeSig chan bool) {
}


func (m *Login) OnDestroy() {
	//一定别忘了关闭RPC
	m.GetServer().OnDestroy()
}

//消息回调
func (m *Login) login(session gate.Session,msg map[string]interface{}) (result string, err string) {
	//time.Sleep(1)
	if msg["userId"] == nil {
		result = "userId cannot be nil"
		return "", result
	}
	userId := msg["userId"].(string)

	err = session.Bind(userId)
	if err != "" {
		return
	}
	fmt.Println(session.GetUserId())
	session.Set("logining", "true")
	session.Push()
	return fmt.Sprint("login success ", userId),""
}