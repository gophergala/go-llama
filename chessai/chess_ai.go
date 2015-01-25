package chessai

import (
	"encoding/json"
	"fmt"
	"github.com/gophergala/go-llama/chessverifier"
	"github.com/gophergala/go-llama/intchess"
	"golang.org/x/net/websocket"
	"log"
)

type AI struct {
	PropUsername     string
	PropPassword     string
	VersesAi         bool
	FirstUse         bool
	ws               *websocket.Conn
	User             *intchess.User
	sendMessages     chan string
	receivedMessages chan string
	Solve            func(chessverifier.GameState) []byte
	Chat             func(messageId)
}

var ai AI

func Make(PropUsername string, PropPassword string, VersesAi bool, FirstUse bool, Solve func(chessverifier.GameState) []byte, Chat func(messageId)) {
	ai.PropPassword = PropPassword
	ai.PropUsername = PropUsername
	ai.FirstUse = FirstUse
	ai.Solve = Solve
	ai.sendMessages = make(chan string, 256)
	ai.receivedMessages = make(chan string, 256)
}

//addr is url of server
//host can be http://localhost
func Run(addr string, host string) {

	ws, err := websocket.Dial(addr, "", host)

	if err != nil {
		log.Fatal(err)
	}

	ai.ws = ws

	defer close(ai.sendMessages)
	defer close(ai.receivedMessages)

	go ai.Reader()
	go ai.Writer()
	ai.Runner()
}

func SendChat(messageId int) {
	request := intchess.APIChatRequest{
		Type:      "chat_request",
		messageId: messageId,
	}
	msg, _ := json.Marshal(request)
	ai.sendMessages <- string(msg)
}

func (a *AI) Reader() {
	for {
		var reply string
		err := websocket.Message.Receive(a.ws, &reply)
		if err != nil {
			break
		}

		a.receivedMessages <- reply
	}
	a.ws.Close()
}

func (a *AI) Writer() {
Loop:
	for {
		for msg := range a.sendMessages {
			err := websocket.Message.Send(a.ws, msg)
			if err != nil {
				break Loop
			}
		}
		//how to detect if broken if no messages to send?
		if !a.ws.IsClientConn() {
			break Loop
		}
	}
	a.ws.Close()
}

func (a *AI) Runner() {
	if a.FirstUse {
		a.attemptCreateAndAuthenticate(a.PropUsername, a.PropPassword, a.VersesAi)
	} else {
		a.attemptAuthentication(a.PropUsername, a.PropPassword)
	}
	for {
		for msg := range a.receivedMessages {
			a.DecodeMessage([]byte(msg))
		}
	}

}

func (a *AI) attemptAuthentication(username string, proposed_password string) {
	request := intchess.APIAuthenticationRequest{
		Type:      "authentication_request",
		Username:  username,
		UserToken: proposed_password,
	}
	msg, _ := json.Marshal(request)
	ai.sendMessages <- string(msg)
	return
}

func (a *AI) attemptCreateAndAuthenticate(username string, proposed_password string, verses_ai bool) {
	r := intchess.APISignupRequest{
		Type:      "signup_request",
		Username:  username,
		UserToken: proposed_password,
		IsAi:      true,
		VersesAi:  verses_ai,
	}
	msg, _ := json.Marshal(r)
	ai.sendMessages <- string(msg)
	return
}

func (a *AI) DecodeMessage(message []byte) {
	var t intchess.APITypeOnly
	if err := json.Unmarshal(message, &t); err != nil {
		fmt.Println("Just receieved a message I couldn't decode:")
		fmt.Println(string(message))
		fmt.Println("Exact error: " + err.Error())
		fmt.Println()
		return
	}

	switch t.Type {
	case "authentication_response":
		var r intchess.APIAuthenticationResponse
		if err := json.Unmarshal(message, &r); err != nil {
			return
		}
		if r.Response == "ok" && r.User != nil {
			a.User = r.User
		} else {
			log.Fatalf("Could not sign in.")
		}
	case "game_request":

		//accept all game requests
		a.SendGameAccept()
	case "game_response_rejection":
		//bleh, ignore it
	case "game_response_accepted":
		//bleh, ignore it
	case "game_move_update":
		var r intchess.APIGameOutput
		if err := json.Unmarshal(message, &r); err != nil {
			return
		}
		a.processMove(r.Game)
	case "game_over":
		//do we need to do anything?
	case "game_chat":
		var r intchess.APIGameChat
		if err := json.Unmarshal(message, &r); err != nil {
			return
		}
		a.Chat(r.MessageId)
	default:
		log.Printf("I'm not familiar with type %s\n", t.Type)
	}
}

func (a *AI) SendGameAccept() {
	r := intchess.APIGameResponse{
		Type:     "game_response",
		Response: "ok",
	}
	msg, _ := json.Marshal(r)
	ai.sendMessages <- string(msg)
}

func (a *AI) processMove(g *intchess.ChessGame) {
	if len(g.GameMoves)%2 == 0 && g.WhitePlayer.Username == a.User.Username || len(g.GameMoves)%2 == 1 && g.BlackPlayer.Username == a.User.Username {
		//is yer turn

		//now verify that the move is allowable
		//first, we need to convert all our moves from []GameMove to [][]Byte
		bMoves := make([][]byte, len(g.GameMoves))
		for index, move := range g.GameMoves {
			bMoves[index] = []byte(move.Move)
		}

		//get the current board state
		curGameState := chessverifier.GetBoardState(&bMoves)

		//get what to do
		r := intchess.APIGameMoveRequest{
			Type: "game_move_request",
			Move: string(a.Solve(curGameState)),
		}

		msg, _ := json.Marshal(r)
		ai.sendMessages <- string(msg)
	}
}
