package piece

import "chess/utility"

type ValidMovefunc func(toX int, toY int) bool

func Pawn(toX, toY int) bool {
	return true
}
func Knight(toX, toY int) bool {
	return true
}
func Bishop(toX, toY int) bool {
	return true
}
func Rook(toX, toY int) bool {
	return true
}
func Queen(toX, toY int) bool {
	return true
}
func King(toX, toY int) bool {
	return true
}
func GetMoveValidator(type_ int) ValidMovefunc {
	switch type_ {
	case utility.Pawn:
		return Pawn
	case utility.Bishop:
		return Bishop
	case utility.King:
		return King
	case utility.Knight:
		return Knight
	case utility.Queen:
		return Queen
	case utility.Rook:
		return Rook
	}
	return nil
}
