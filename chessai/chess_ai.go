package chessai

import (
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
	Solve            func(*[][]byte) []byte
}

var ai AI

func Make(PropUsername string, PropPassword string, VersesAi bool, FirstUse bool, Solve func(*[][]byte) []byte) {
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

func (a *AI) Reader() {
	for {
		var reply string
		err := websocket.Message.Receive(a.ws, &reply)
		if err != nil {
			break
		}

		a.sendMessages <- reply
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

func (a *AI) UseCredentialsAndAuthenticate(username string, proposed_password string, verses_ai bool, first_use bool) {

}

func (a *AI) attemptAuthentication(username string, proposed_password string) {

}

func (a *AI) attemptCreateAndAuthenticate(username string, proposed_password string, verses_ai bool) {

}

func (a *AI) DecodeMessage(message []byte) {

}
