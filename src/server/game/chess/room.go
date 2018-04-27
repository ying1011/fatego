package chess

import (
	"server/game/base"
)

type Room struct{
	base.Room
	chessBoard Board
}

func (w *Room)init()  {

}

