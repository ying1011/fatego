package base

import (
	"github.com/liangdas/mqant/gate"
)

type Player struct{
	uid			int
	name		string
	session		gate.Session
}