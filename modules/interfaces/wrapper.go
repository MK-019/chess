package interfaces

type ChessBoardI interface {
	PrintBoard()
	MakeMove(PeiceId string, toX, toY int) error
}

type PeiceI interface {
	GetPeice()
	SetPeice()
}
