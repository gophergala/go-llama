package intchess

import (
	"code.google.com/p/websocket"
	"encoding/json"
	"fmt"
)

type Connection struct {
	ws           *websocket.Conn
	User         *User
	GameIndex    *int
	sendMessages chan string
}

func (c *Connection) Reader() {
	for {
		var reply string
		err := websocket.Message.Receive(c.ws, &reply)
		if err != nil {
			// if c.User == nil {
			// 	fmt.Println("Connection lost for " + c.User.Username)
			// } else {
			// 	fmt.Println("Connection lost for anon user")
			// }
			break
		}
		//packet := &GamePacket{Message: reply, User: c.Player}

		//GS.broadcast <- packet
		c.DecodeMessage([]byte(reply))
	}
	c.ws.Close()
}

func (c *Connection) Writer() {
Loop:
	for {
		for msg := range c.sendMessages {
			err := websocket.Message.Send(c.ws, msg)
			if err != nil {
				// if c.User == nil {
				// 	fmt.Println("Connection lost for " + c.User.Username)
				// } else {
				// 	fmt.Println("Connection lost for anon user")
				// }
				break Loop
			}
		}
		//how to detect if broken if no messages to send?
		if !c.ws.IsClientConn() {
			// if c.User == nil {
			// 	fmt.Println("Connection lost for " + c.User.Username)
			// } else {
			// 	fmt.Println("Connection lost for anon user")
			// }
			break Loop
		}
	}
	c.ws.Close()
}

func (c *Connection) DecodeMessage(message []byte) {
	var t APITypeOnly
	if err := json.Unmarshal(message, &t); err != nil {
		fmt.Println("Just receieved a message I couldn't decode:")
		fmt.Println(string(message))
		fmt.Println("Exact error: " + err.Error())
		fmt.Println()
		return
	}

	switch t.Type {
	case "authentication_request":
		if c.User != nil {
			//they are already authenticated
			return
		}
		var a APIAuthenticationRequest
		if err := json.Unmarshal(message, &a); err != nil {
			return
		}

		if u := AttemptLogin(a.Username, a.UserToken); u != nil {
			c.User = u
			fmt.Println("Anon user authenticates as " + c.User.Username)
		}
		c.SendAuthenticationResponse()

	case "game_response":
		var r APIGameResponse
		if err := json.Unmarshal(message, &r); err != nil {
			return
		}
		if r.Response == "ok" {
			c.AcceptGameRequest()
		}

	default:
		fmt.Println("I'm not familiar with type " + t.Type)
	}
}

func (c *Connection) SendAuthenticationResponse() {
	var ResponseStr string
	if c.User == nil {
		ResponseStr = "incorrect username/access token combo"
	} else {
		ResponseStr = "ok"
	}

	resp := APIAuthenticationResponse{
		Type:     "authentication_response",
		Response: ResponseStr,
		User:     c.User,
	}

	serverMsg, _ := json.Marshal(resp)
	c.sendMessages <- string(serverMsg)
	return
}

func (c *Connection) SendGameRequest(versesConnection *Connection, gameIndex int) {
	gameReq := APIGameRequest{
		Type:     "game_request",
		Opponent: versesConnection.User,
	}
	serverMsg, _ := json.Marshal(gameReq)
	c.sendMessages <- string(serverMsg)
	c.GameIndex = &gameIndex
	return
}

func (c *Connection) SendGameExpired() {
	serverOut := APIGameResponse{
		Type:     "game_response_rejection",
		Response: "expired",
	}
	serverMsg, _ := json.Marshal(serverOut)
	c.sendMessages <- string(serverMsg)
	c.GameIndex = nil
}

func (c *Connection) AcceptGameRequest() {
	if c.GameIndex == nil {
		c.SendGameExpired()
	}
	//mark them as accepting the game, and alert them that we're happy
	serverOut := APIGameResponse{
		Type:     "game_response_accepted",
		Response: "ok",
	}
	serverMsg, _ := json.Marshal(serverOut)
	c.sendMessages <- string(serverMsg)

	//if the game is ready to start, it will start on its own
	Games[*c.GameIndex].PlayerAccepts(c.User)

}

func (c *Connection) SendGameUpdate(g *ChessGame, Type string) {
	serverOut := APIGameOutput{
		Type: Type,
		Game: g,
	}
	serverMsg, _ := json.Marshal(serverOut)
	c.sendMessages <- string(serverMsg)
}
