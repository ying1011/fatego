package base

import (
	"github.com/liangdas/mqant/gate"
)

type Player struct{
	uid			string
	name		string
	session		gate.Session
}

func NewPlayer()  {
	
}

func (w *Player)Reset()  {
	
}

func (w *Player)GetUid() string {
	return w.uid
}

func (w *Player)Bind(session gate.Session) bool {
	if session != nil {
		w.session = session
		return true
	}else {
		return false
	}
}

func (w *Player)UnBind() {
	w.session = nil
}