package main

import (
	"fmt"
	"github.com/gophergala/go-llama/chessai"
)

func main() {
	fmt.Println("Test AI running")

	PropUsername := "testAI"
	PropPassword := "testAI"
	VersesAi := true
	FirstUse := false
	chessai.Make(PropUsername, PropPassword, VersesAi, FirstUse, MySolver)

	addr := "http://192.168.1.25:800/ws"
	host := "http://localhost"

	chessai.Run(addr, host) //this is a blocking call
}

func MySolver(game chessverifier.GameState) []byte {
	//best solver ever
	return nil
}
