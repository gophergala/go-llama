package chessverifier

import (
	"fmt"
)

type BoardState [8][8]string

func main() {
	fmt.Printf("Hello world!\n")

}

func GetValidMoves(board BoardState, piece []rune) [][]rune {

}

func GetAllValidMoves(board BoardState) [][]rune {
	for 
}

func IsMoveValid(board BoardState, move []rune) bool {
	var moveList = GetValidMoves(board, move[0:2])
	for x := range moveList {
		if moveEqual(move, moveList[x]) {
			return true
		}
	}
	return false
}

func GetBoardState(moveList [][]rune) BoardState {
	var board BoardState = startBoardState
	for moveNum := range moveList {
		if IsMoveValid(board, moveList[moveNum]) {
			//@todo finish
		}
	}
}

func getSquareIndexes(squareID []rune) (x, y int) {
	y = int(squareID[1] - '1')
	x = int(squareID[0] - 'a')
	return
}

func moveEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

var startBoardState BoardState = BoardState{
	[8]string{"WR1", "WP1", "", "", "", "", "BP1", "BR1"},
	[8]string{"WK1", "WP2", "", "", "", "", "BP2", "BK1"},
	[8]string{"WB1", "WP3", "", "", "", "", "BP3", "BB1"},
	[8]string{"WQ1", "WP4", "", "", "", "", "BP4", "BK1"},
	[8]string{"WK1", "WP5", "", "", "", "", "BP5", "BK1"},
	[8]string{"WB3", "WP6", "", "", "", "", "BP6", "BB2"},
	[8]string{"WK3", "WP7", "", "", "", "", "BP7", "BK2"},
	[8]string{"WR2", "WP8", "", "", "", "", "BP8", "BR2"},
}
