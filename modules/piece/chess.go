package piece

import "strconv"

// Define a struct for a chess piece
type ChessPiece struct {
	Type      int // Type of the piece (Pawn, Knight, etc.)
	Color     int // Color of the piece (White or Black)
	Row       int // Row position on the chessboard
	Col       int // Column position on the chessboard
	ValidMove ValidMovefunc
}
type PeiceId struct {
	Id string
	ChessPiece
}

func NewChessPiece(type_, color, row, col int) *PeiceId {
	return GeneratePeiceId(type_, color, col, GeneratePeice(type_, color, row, col))
}

func NewEmptyChessPiece() *PeiceId {
	return GenerateNewEmptyPeiceId()
}
func GeneratePeiceId(type_, color int, col int, p ChessPiece) *PeiceId {

	return &PeiceId{
		Id:         strconv.Itoa(type_) + strconv.Itoa(color) + strconv.Itoa(col),
		ChessPiece: p,
	}
}

func GeneratePeice(type_, color, row, col int) ChessPiece {
	return ChessPiece{
		Type:      type_,
		Color:     color,
		Row:       row,
		Col:       col,
		ValidMove: GetMoveValidator(type_),
	}
}

func GenerateNewEmptyPeiceId() *PeiceId {
	return &PeiceId{
		Id:         "",
		ChessPiece: ChessPiece{},
	}

}
