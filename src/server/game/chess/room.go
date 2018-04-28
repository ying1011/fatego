package chess

import (
	"server/game/base"
)

type Room struct{
	base.Room
	seats 		[]*base.Player
	maxSeat 	int
	chessBoard 	Board
}

func (w *Room)init()  {

}

