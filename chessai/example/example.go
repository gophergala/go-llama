package main

import (
	"fmt"
	"github.com/gophergala/go-llama/chessai"
	"github.com/gophergala/go-llama/chessverifier"
)

func main() {
	fmt.Println("Test AI running")

	PropUsername := "testAI"
	PropPassword := "testAI"
	VersesAi := true
	FirstUse := false
	chessai.Make(PropUsername, PropPassword, VersesAi, FirstUse, MySolver, IncomingChat)

	addr := "ws://192.168.1.25:8080/ws"
	host := "http://localhost"

	chessai.Run(addr, host) //this is a blocking call
}

func MySolver(game chessverifier.GameState) []byte {
	//best solver ever
	return []byte("a2-a3")
	//valid moves - who cares about those PFFT
}

func IncomingChat(messageId int) {
	//echo the same message back
	chessai.SendChat(messageId)
}
