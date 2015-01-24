package intchess

import (
	"code.google.com/p/websocket"
	"encoding/json"
	"fmt"
)

type Connection struct {
	ws           *websocket.Conn
	User         *User
	Game         *Game
	sendMessages chan string
}

func (c *Connection) Reader() {
	for {
		var reply string
		err := websocket.Message.Receive(c.ws, &reply)
		if err != nil {
			fmt.Println("Connection lost for " + c.User.Username)
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
				fmt.Println("Connection lost for " + c.User.Username)
				break Loop
			}
		}
		//how to detect if broken if no messages to send?
		if !c.ws.IsClientConn() {
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
		var a APIAuthenticationRequest
		if err := json.Unmarshal(message, &a); err != nil {
			fmt.Println("Just receieved a message I couldn't decode:")
			fmt.Println(string(message))
			fmt.Println("Exact error: " + err.Error())
			fmt.Println()
			return
		}

		if u := AttemptLogin(a.Username, a.UserToken); u != nil {
			c.User = u
		}
		c.SendAuthenticationResponse()

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

func (c *Connection) SendGameRequest(versesConnection *Connection) {
	gameReq := APIGameRequest{
		Type:     "game_request",
		Opponent: versesConnection.User,
	}
	serverMsg, _ := json.Marshal(gameReq)
	c.sendMessages <- string(serverMsg)
	return
}
