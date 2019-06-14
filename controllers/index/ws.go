package index

import (
	"log"
	"time"
	"net/http"

	"hackathon/modules"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Ws(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	types := vars["type"]
	cType := ""
	if types != nil {
		cType = types[0]
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	channel := make(chan string, 100000)
	client := modules.Client{
		Ch:   channel,
		Type: cType,
	}
	client.Connect()
	go send(client, conn)
	go read(client, conn)

}

func send(c modules.Client, conn *websocket.Conn) {
	defer func(){
		c.Disconnect()
		conn.Close()
	}()

	for {
		message, ok := <-c.Ch
		conn.SetWriteDeadline(time.Now().Add(10*time.Second))
		if !ok {
			conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			return
		}
	}
}

func read(c modules.Client, conn *websocket.Conn) {
	defer func(){
		c.Disconnect()
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error:%v", err)
			}
			break
		}
		msg := modules.Message {
			M:    string(message),
			Type: c.Type,
		}
		modules.Send(msg)
	}
}
