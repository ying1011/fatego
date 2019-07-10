package chess

type Piece struct{
	x int
	y int
	colorId int		//颜色
	typeId int		//类型
}

type Board struct {
	wide	int		//宽度
	high	int		//高度
	chessPieces []Piece
}

