package intchess

import (
	"code.google.com/p/websocket"
	"fmt"
)

type Connection struct {
	ws           *websocket.Conn
	Player       User
	sendMessages chan string
}

func (c *Connection) Reader() {
	for {
		var reply string
		err := websocket.Message.Receive(c.ws, &reply)
		if err != nil {
			fmt.Println("Connection lost for " + c.Player.Username)
			break
		}
		packet := &GamePacket{Message: reply, User: &c.Player}

		GS.broadcast <- packet
	}
	c.ws.Close()
}

func (c *Connection) Writer() {
Loop:
	for {
		for msg := range c.sendMessages {
			err := websocket.Message.Send(c.ws, msg)
			if err != nil {
				fmt.Println("Connection lost for " + c.Player.Username)
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
