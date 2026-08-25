package board

import (
	"chess/modules/interfaces"
	"chess/modules/piece"
	"chess/utility"
	"fmt"
)

var _ interfaces.ChessBoardI = (*ChessBoard)(nil)

// Define a chessboard
type ChessBoard struct {
	Squares     [8][8]*piece.PeiceId
	pieceMapper map[string]*piece.ChessPiece
	*piece.ChessPiece
}

// Initialize a chessboard, factory function for chessboard
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{}
	board.pieceMapper = make(map[string]*piece.ChessPiece, 0)
	// Set up the starting position (simplified for demonstration)
	board.Squares[0] = [8]*piece.PeiceId{
		piece.NewChessPiece(utility.Rook, utility.Black, 0, 0),
		piece.NewChessPiece(utility.Knight, utility.Black, 0, 1),
		piece.NewChessPiece(utility.Bishop, utility.Black, 0, 2),
		piece.NewChessPiece(utility.Queen, utility.Black, 0, 3),
		piece.NewChessPiece(utility.King, utility.Black, 0, 4),
		piece.NewChessPiece(utility.Bishop, utility.Black, 0, 5),
		piece.NewChessPiece(utility.Knight, utility.Black, 0, 6),
		piece.NewChessPiece(utility.Rook, utility.Black, 0, 7),
	}
	board.Squares[1] = [8]*piece.PeiceId{
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 0),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 1),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 2),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 3),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 4),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 5),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 6),
		piece.NewChessPiece(utility.Pawn, utility.Black, 1, 7),
	}
	emptyPeiceId := make([]*piece.PeiceId, 0)
	for i := 0; i <= 7; i++ {
		emptyPeiceId = append(emptyPeiceId, piece.NewEmptyChessPiece())
	}
	for i := 2; i <= 5; i++ {
		board.Squares[i] = [8]*piece.PeiceId(emptyPeiceId)
	}

	board.Squares[7] = [8]*piece.PeiceId{
		piece.NewChessPiece(utility.Rook, utility.White, 7, 0),
		piece.NewChessPiece(utility.Knight, utility.White, 7, 1),
		piece.NewChessPiece(utility.Bishop, utility.White, 7, 2),
		piece.NewChessPiece(utility.Queen, utility.White, 7, 3),
		piece.NewChessPiece(utility.King, utility.White, 7, 4),
		piece.NewChessPiece(utility.Bishop, utility.White, 7, 5),
		piece.NewChessPiece(utility.Knight, utility.White, 7, 6),
		piece.NewChessPiece(utility.Rook, utility.White, 7, 7),
	}
	board.Squares[6] = [8]*piece.PeiceId{
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 0),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 1),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 2),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 3),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 4),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 5),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 6),
		piece.NewChessPiece(utility.Pawn, utility.White, 1, 7),
	}
	peiceInitailLoc := []int{0, 1, 6, 7}
	for _, loc := range peiceInitailLoc {
		for _, value1 := range board.Squares[loc] {
			board.pieceMapper[value1.Id] = &value1.ChessPiece
		}
	}

	return board
}

func (board *ChessBoard) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%v ", *board.Squares[i][j])
		}
		fmt.Println()
	}
}

// Function to make a move on the chessboard
func (board *ChessBoard) MakeMove(PeiceId string, toX, toY int) error {
	// Check if the move is valid (simplified for demonstration)
	if toX < 0 || toX >= 8 || toY < 0 || toY >= 8 {
		return fmt.Errorf("invalid move: Out of bounds")
	}
	// Implement more move validation logic here (e.g., piece-specific rules)

	// Perform the move
	peice := board.pieceMapper[PeiceId]
	if peice.ValidMove(toX, toY) {
		temp := board.Squares[toX][toY]
		board.Squares[toX][toY] = board.Squares[peice.Row][peice.Col]
		board.Squares[peice.Row][peice.Col] = temp
		peice.Col = toY
		peice.Row = toX
	} else {
		return fmt.Errorf("invalid move: Out of bounds")
	}

	return nil
}
