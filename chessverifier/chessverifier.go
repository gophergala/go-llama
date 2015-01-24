package chessverifier

// import (
// 	"fmt"
// )

type BoardState [8][8][]byte
type GameState struct {
	board    BoardState
	moveList [][]byte
}

func NewGame() GameState {
	var game GameState
	game.board = startBoardState
	return game
}

// func main() {
// 	fmt.Printf("Hello world!\n")

// }

func GetValidMoves(game *GameState, x, y int) (validMoves [][]byte) {
	var piece = game.board[x][y]
	if len(piece) == 0 {
		return [][]byte{}
	}
	var newMove = []byte{}
	switch piece[1] {
	case 'P':
		if piece[0] == 'w' {
			newMove = []byte{byte(x + '1'), byte(y + 'a'), '-', byte(x + '1' + 1), byte(y + 'a')}
		} else {
			newMove = []byte{byte(x + '1'), byte(y + 'a'), '-', byte(x + '1' - 1), byte(y + 'a')}
		}
		validMoves = append(validMoves, newMove)
	}

	return [][]byte{[]byte{'a', '1', '-', 'b', '2'}}
}

func GetAllValidMoves(game *GameState) (validMoves [][]byte) {
	for x := range game.board {
		for y := range game.board[x] {
			if len(game.board[x][y]) != 0 {
				validMoves = append(validMoves, GetValidMoves(game, x, y)...)
			}
		}
	}
	return
}

func IsMoveValid(game *GameState, move *[]byte) bool {
	var x, y = getSquareIndices((*move)[0:2])
	var moveList = GetValidMoves(game, x, y)
	for i := range moveList {
		if moveEqual(move, &moveList[i]) {
			return true
		}
	}
	return false
}

func GetBoardState(moveList *[][]byte) GameState {
	var game GameState = NewGame()
	for moveNum := range *moveList {
		MakeMove(&game, &(*moveList)[moveNum])
	}
	return game
}

//This function is used to apply a move (eg "a2-a4") to the board
//it will verify first if the move is valid
func MakeMove(game *GameState, move *[]byte) {
	if IsMoveValid(game, move) {
		ox, oy := getSquareIndices((*move)[0:2])
		nx, ny := getSquareIndices((*move)[3:5])
		var piece = game.board[ox][oy]
		game.board[ox][oy] = []byte{}
		game.board[nx][ny] = piece
	}
}

//This function is used to convert a piece location in Algebraic notation
//to a piece location in the internal board 2d slice
func getSquareIndices(squareID []byte) (x, y int) {
	y = int(squareID[1] - '1')
	x = int(squareID[0] - 'a')
	return
}

//This function is used to test the equality of two byte slices
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
