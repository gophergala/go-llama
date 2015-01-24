package chessverifier

import (
	"fmt"
)

type BoardState [8][8][]byte

func main() {
	fmt.Printf("Hello world!\n")

}

func GetValidMoves(board BoardState, piece []byte) (moveList [][]byte) {

}

func GetAllValidMoves(board BoardState) (moveList [][]byte) {
	for x := range board {
		for y := range board[x] {
			if len(board[x][y]) == 0 {
				moveList = append(moveList, GetValidMoves(board, board[x][y])...)
			}
		}
	}
	return
}

func IsMoveValid(board BoardState, move []byte) bool {
	var moveList = GetValidMoves(board, move[0:2])
	for x := range moveList {
		if moveEqual(move, moveList[x]) {
			return true
		}
	}
	return false
}

func GetBoardState(moveList [][]byte) BoardState {
	var board BoardState = startBoardState
	for moveNum := range moveList {
		if IsMoveValid(board, moveList[moveNum]) {
			//@todo finish
		}
	}
	return board
}

func getSquareIndexes(squareID []byte) (x, y int) {
	y = int(squareID[1] - '1')
	x = int(squareID[0] - 'a')
	return
}

func moveEqual(a, b []byte) bool {
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
	[8][]byte{[]byte{'W', 'R', '1'}, []byte{'W', 'P', '1'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '1'}, []byte{'B', 'R', '1'}},
	[8][]byte{[]byte{'W', 'K', '1'}, []byte{'W', 'P', '2'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '2'}, []byte{'B', 'K', '1'}},
	[8][]byte{[]byte{'W', 'B', '1'}, []byte{'W', 'P', '3'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '3'}, []byte{'B', 'B', '1'}},
	[8][]byte{[]byte{'W', 'Q', '1'}, []byte{'W', 'P', '4'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '4'}, []byte{'B', 'K', '1'}},
	[8][]byte{[]byte{'W', 'K', '1'}, []byte{'W', 'P', '5'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '5'}, []byte{'B', 'K', '1'}},
	[8][]byte{[]byte{'W', 'B', '3'}, []byte{'W', 'P', '6'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '6'}, []byte{'B', 'B', '2'}},
	[8][]byte{[]byte{'W', 'K', '3'}, []byte{'W', 'P', '7'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '7'}, []byte{'B', 'K', '2'}},
	[8][]byte{[]byte{'W', 'R', '2'}, []byte{'W', 'P', '8'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '8'}, []byte{'B', 'R', '2'}},
}
