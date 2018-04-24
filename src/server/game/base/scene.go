package base

import "github.com/liangdas/mqant/module"

var (
	Stop 	= 1
	Active  = 2
)

type Scene struct{
	sid 		int
	state		int
	players		[]*Player
	module		module.RPCModule
}