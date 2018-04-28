package base

import (
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/gate"
)

var (
	Stop 	= 1
	Active  = 2
)

type Scene struct{
	sid 		int
	state		int
	players		map[string]*Player
	module		module.RPCModule
}

func (w *Scene)AddPlayer(session gate.Session) {
	uid := session.GetUserid()

	w.players[uid] = &Player{uid:"",name:""}
}

func (w *Scene)RemovePlayer(uid string) {
	delete(w.players, uid)
}