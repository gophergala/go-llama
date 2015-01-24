package chessverifier

// import (
// 	"fmt"
// )

type BoardState [8][8][]byte
type GameState struct {
	board    BoardState
	moveList [][]byte
}

func newGame() GameState {
	var game GameState
	game.board = startBoardState
	return game
}

// func main() {
// 	fmt.Printf("Hello world!\n")

// }

func GetValidMoves(game *GameState, piece *[]byte) (validMoves [][]byte) {
	return [][]byte{}
}

func GetAllValidMoves(game *GameState) (validMoves [][]byte) {
	for x := range game.board {
		for y := range game.board[x] {
			if len(game.board[x][y]) != 0 {
				validMoves = append(validMoves, GetValidMoves(game, &game.board[x][y])...)
			}
		}
	}
	return
}

func IsMoveValid(game *GameState, move *[]byte) bool {
	var x, y = getSquareIndices((*move)[0:1])
	var moveList = GetValidMoves(game, &game.board[x][y])
	for i := range moveList {
		if moveEqual(move, &moveList[i]) {
			return true
		}
	}
	return false
}

func GetBoardState(moveList *[][]byte) GameState {
	var game GameState = newGame()
	for moveNum := range *moveList {
		MakeMove(&game, &(*moveList)[moveNum])
	}
	return game
}

func MakeMove(game *GameState, move *[]byte) {
	if IsMoveValid(game, move) {
		var x, y = getSquareIndices((*move)[0:1])
		var piece = game.board[x][y]
		game.board[(*move)[0]][(*move)[1]] = []byte{}
		game.board[(*move)[4]][(*move)[5]] = piece
	}
}

func getSquareIndices(squareID []byte) (x, y int) {
	y = int(squareID[1] - '1')
	x = int(squareID[0] - 'a')
	return
}

func moveEqual(a, b *[]byte) bool {
	if len(*a) != len(*b) {
		return false
	}

	for i := range *a {
		if (*a)[i] != (*b)[i] {
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
