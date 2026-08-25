package main

import (
	"chess/modules/board"
	"fmt"
)

func main() {
	board := board.NewChessBoard()
	board.PrintBoard()
	board.MakeMove("120", 3, 1)

	fmt.Println()
	board.PrintBoard()

}
