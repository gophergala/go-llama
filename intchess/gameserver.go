package intchess

import (
	"code.google.com/p/websocket"
	//	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
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

var Games []ChessGame

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
	go GS.matchCreationLoop()
}

func (GS *GameServer) matchCreationLoop() {
	for {
		//find out if enough time has passed since the last game loop. If it hasn't, then don't run

		fmt.Println("Server: There is currently " + strconv.Itoa(GS.NumActiveConnections()) + " players online.")
		GS.clearExpiredGames()
		GS.attemptMatchCreations()

		time.Sleep(1 * time.Second)
	}
}

func (GS *GameServer) attemptMatchCreations() {
	//make a list of all connected players that are not in a match
	var outOfGameConnections []*Connection

	for conn := range GS.connections {
		if GS.connections[conn] == true && conn.InGame == false && conn.User != nil {
			outOfGameConnections = append(outOfGameConnections, conn)
		}
	}
	lenOutGamers := len(outOfGameConnections)
	//make sure there's at least two players
	if lenOutGamers < 2 {
		return
	}
	//make the array even
	if lenOutGamers%2 == 1 {
		outOfGameConnections = outOfGameConnections[0 : lenOutGamers-2] //sorry, last player
	}

	//try match randomly chosen players by randomly sorting the array
	randomisedOutOfGameConnections := make([]*Connection, lenOutGamers)
	randPerm := rand.Perm(lenOutGamers)

	for i := 0; i < lenOutGamers; i++ {
		randomisedOutOfGameConnections[i] = outOfGameConnections[randPerm[i]]
	}

	//then pair with i and i+1
	for i := 0; i < lenOutGamers; i += 2 {
		GS.attemptMatchCreation(outOfGameConnections[i], outOfGameConnections[i+1])
	}
}

func (GS *GameServer) attemptMatchCreation(player1 *Connection, player2 *Connection) {
	fmt.Println("Server: Attempting to match " + player1.User.Username + " with " + player2.User.Username)
	player1.SendGameRequest(player2)
	player2.SendGameRequest(player1)
	//now create a provisional game
	g := NewGame()
	games := append(games, g)
}
