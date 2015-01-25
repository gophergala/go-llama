package chessverifier

import (
	// "encoding/json"
	"fmt"
)

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
	fmt.Println(string(piece))
	if len(piece) == 0 {
		return [][]byte{}
	}
	var newMove = []byte{}
	var white = piece[0] == 'W'
	switch piece[1] { //make sure the king is not in check before anything else
	case 'P': //@todo add on passant
		var ymove = 1
		if !white {
			ymove = -1
		}
		var newSquare = [2]int{x, y + ymove}
		var free, taking = canLand(game, newSquare, white)
		if free && !taking {
			newMove = getMove([2]int{x, y}, newSquare)
			validMoves = append(validMoves, newMove)
		}
		newSquare[0] = newSquare[0] + 1
		free, taking = canLand(game, newSquare, white)
		if free && taking {
			newMove = getMove([2]int{x, y}, newSquare)
			validMoves = append(validMoves, newMove)
		}
		newSquare[0] = newSquare[0] - 2
		free, taking = canLand(game, newSquare, white)
		if free && !taking {
			newMove = getMove([2]int{x, y}, newSquare)
			validMoves = append(validMoves, newMove)
		}

	case 'R':
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{1, 0}, white)...)  //right
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{-1, 0}, white)...) //left
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{0, 1}, white)...)  //up
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{0, -1}, white)...) //down

	case 'N':
		var moveDiffs [8][2]int = [8][2]int{[2]int{2, 1}, [2]int{2, -1}, [2]int{1, 2},
			[2]int{-1, 2}, [2]int{-2, 1}, [2]int{-2, -1}, [2]int{-1, -2}, [2]int{1, -2}}

		for i := range moveDiffs {
			var canLand, _ = canLand(game, [2]int{x + moveDiffs[i][0], y + moveDiffs[i][1]}, white)
			if canLand {
				var newMove = getMove([2]int{x, y}, [2]int{x + moveDiffs[i][0], y + moveDiffs[i][1]})
				var testGame = testMove(*game, &newMove)
				if !isCheck(&testGame, white) {
					validMoves = append(validMoves, newMove)
				}
			}
		}

	case 'B':
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{1, 1}, white)...)   //up-right
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{-1, -1}, white)...) //down-left
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{1, -1}, white)...)  //down-right
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{-1, 1}, white)...)  //down-left

	case 'Q':
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{1, 0}, white)...)   //right
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{-1, 0}, white)...)  //left
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{0, 1}, white)...)   //up
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{0, -1}, white)...)  //down
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{1, 1}, white)...)   //up-right
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{-1, -1}, white)...) //down-left
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{1, -1}, white)...)  //down-right
		validMoves = append(validMoves, moveDirection(game, x, y, [2]int{-1, 1}, white)...)  //down-left

	case 'K':
		var moveDiffs [8][2]int = [8][2]int{[2]int{1, 0}, [2]int{1, 1}, [2]int{0, 1},
			[2]int{-1, 1}, [2]int{-1, 0}, [2]int{-1, -1}, [2]int{0, -1}, [2]int{1, -1}}

		for i := range moveDiffs {
			var canLand, _ = canLand(game, moveDiffs[i], white)
			if canLand {
				var newMove = getMove([2]int{x, y}, [2]int{x + moveDiffs[i][0], y + moveDiffs[i][1]})
				var testGame = testMove(*game, &newMove)
				if !isCheck(&testGame, white) {
					validMoves = append(validMoves, newMove)
				}
			}
		}
	}

	return
}

func moveDirection(game *GameState, x, y int, direction [2]int, white bool) (validMoves [][]byte) {
	for i := 1; i < 8; i++ {
		var canLand, taking = canLand(game, [2]int{x + (direction[0] * i), y + (direction[1] * i)}, white)
		if canLand {
			validMoves = append(validMoves, getMove([2]int{x, y}, [2]int{x + i, y}))
			if taking {
				break
			}
		} else {
			break
		}
	}
	return validMoves
}

func isCheck(game *GameState, white bool) bool {
	return false
}

//used for testing if a move will place the king in check so that functions can be passed only the board state after the move
func testMove(game GameState, move *[]byte) GameState {
	MakeMove(&game, move)
	return game
}

//white is whether or not the current player is white for handling taking
//if canLand is false taking sould be ignored (it will always be false)
func canLand(game *GameState, square [2]int, white bool) (canLand, taking bool) {
	if !(square[0] >= 0 && square[0] <= 7 && square[1] >= 0 && square[1] <= 7) { //is the destination off the board?
		return false, false
	}
	var piece = game.board[square[0]][square[1]]
	var occupied bool = len(piece) != 0

	if !occupied { //on the board and empty
		return true, false
	}
	if (piece[0] == 'W') == white { //occupied with a piece the same colour as the piece being moved
		return false, false
	}
	return true, true //occupied with an oponent's piece
}

//do the line of formatting to save typing and tidy the code
func getMove(source [2]int, dest [2]int) []byte {
	return []byte{byte(source[0] + 'a'), byte(source[1] + '1'), '-', byte(dest[0] + 'a'), byte(dest[1] + '1')}
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
//note that this does not check if the move is valid because it causes and infinite loop XD
//you therefor need to make sure the move is allowed before trying it
func MakeMove(game *GameState, move *[]byte) {
	// if IsMoveValid(game, move) {
	ox, oy := getSquareIndices((*move)[0:2])
	nx, ny := getSquareIndices((*move)[3:5])
	var piece = game.board[ox][oy]
	game.board[ox][oy] = []byte{}
	game.board[nx][ny] = piece
	// }
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

//Knight is actaully spelt Night did you know?
var startBoardState BoardState = BoardState{
	[8][]byte{[]byte{'W', 'R', '1'}, []byte{'W', 'P', '1'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '1'}, []byte{'B', 'R', '1'}},
	[8][]byte{[]byte{'W', 'N', '1'}, []byte{'W', 'P', '2'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '2'}, []byte{'B', 'N', '1'}},
	[8][]byte{[]byte{'W', 'B', '1'}, []byte{'W', 'P', '3'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '3'}, []byte{'B', 'B', '1'}},
	[8][]byte{[]byte{'W', 'Q', '1'}, []byte{'W', 'P', '4'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '4'}, []byte{'B', 'K', '1'}},
	[8][]byte{[]byte{'W', 'K', '1'}, []byte{'W', 'P', '5'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '5'}, []byte{'B', 'K', '1'}},
	[8][]byte{[]byte{'W', 'B', '3'}, []byte{'W', 'P', '6'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '6'}, []byte{'B', 'B', '2'}},
	[8][]byte{[]byte{'W', 'N', '3'}, []byte{'W', 'P', '7'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '7'}, []byte{'B', 'N', '2'}},
	[8][]byte{[]byte{'W', 'R', '2'}, []byte{'W', 'P', '8'}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{'B', 'P', '8'}, []byte{'B', 'R', '2'}},
}

func Runtest() {
	var game = NewGame()
	game.board[1][0] = []byte{}
	game.board[2][0] = []byte{}
	// var board, _ = json.MarshalIndent(game.board, "", "  ")
	// fmt.Println(string(board))
	var moveList = GetValidMoves(&game, 0, 0)
	for _, move := range moveList {
		fmt.Println(string(move))
	}
}
