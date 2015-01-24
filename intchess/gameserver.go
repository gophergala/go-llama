package intchess

import (
	"code.google.com/p/websocket"
	//	"encoding/json"
	"fmt"
	//	"strconv"
	"time"
)

type GamePacket struct {
	Message string
	User    *User
}

type GameServer struct {
	connections map[*Connection]bool // Inbound messages from the connections.
	broadcast   chan *GamePacket     // Register requests from the connections.
	register    chan *Connection     // Unregister requests from connections.
	unregister  chan *Connection
	count       int
}

var GS = GameServer{
	connections: make(map[*Connection]bool),
	broadcast:   make(chan *GamePacket),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	count:       0,
}

func (GS *GameServer) run() {
	for {
		select {
		case c := <-GS.register:
			GS.activateConnection(c)
		case c := <-GS.unregister:
			GS.deactivateConnection(c)
		case m := <-GS.broadcast:
			//m.User.SetRequest(m.Message)
			GS.broadcastAll(m.Message)
			//default:
			//this is neccessarry so that this select statement is nonblocking
		}
		time.Sleep(500 * time.Nanosecond)
		//GS.gameLoop()
	}
}

func (GS *GameServer) activateConnection(c *Connection) {
	//is this a new player?
	if _, ok := GS.connections[c]; ok {
		//this is an old player reconnecting (does this ever happen!?!)
		GS.connections[c] = true
	} else {
		//this is a new player
		GS.connections[c] = true
		GS.count++
	}
}

func (GS *GameServer) deactivateConnection(c *Connection) {
	//delete(h.connections, c)
	GS.connections[c] = false
	close(c.sendMessages)
}

func (GS *GameServer) broadcastAll(msg string) {
	fmt.Println(msg)
	for conn := range GS.connections {
		if GS.connections[conn] == true {
			select {
			case conn.sendMessages <- msg:
			default:
				GS.deactivateConnection(conn)
			}
		}
	}
}

func (GS *GameServer) NumActiveConnections() int {
	count := 0
	for conn := range GS.connections {
		if GS.connections[conn] == true {
			count++
		}
	}
	return count
}

func WsHandler(ws *websocket.Conn) {
	//a new websocket has been created
	Client := &Connection{sendMessages: make(chan string, 256), ws: ws}

	GS.activateConnection(Client)
	defer GS.deactivateConnection(Client)

	go Client.Writer()
	Client.Reader()

}

func StartGameServer() {
	go GS.run()
}
