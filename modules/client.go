package modules

import(
	"log"
)

type channel chan string
type Client struct {
	Ch   channel
	Type string
}

type Message struct {
	M    string
	Type string
}

var clients map[string]map[channel]bool
var msg chan Message
var connect chan *Client
var disconnect chan *Client

func ClientInit() {
	clients = make(map[string]map[channel]bool)
	msg = make(chan Message, 100000)
	connect = make(chan *Client, 10000)
	disconnect = make(chan *Client, 10000)
	go Broadcaster()
}

func Broadcaster() {
	defer func(){
		err := recover()
		if err != nil {
			log.Printf("broadcaster recover. err:%v", err)
		}
		go Broadcaster()
	}()
	for {
		select{
			case m := <-msg:
				for c := range clients[m.Type] {
					c <- m.M
				}
			case c := <-connect:
				if clients[c.Type] == nil {
					clients[c.Type] = make(map[channel]bool)
				}
				clients[c.Type][c.Ch] = true
			case c := <-disconnect:
				delete(clients[c.Type], c.Ch)
				close(c.Ch)
		}

	}
}

func Send(m Message) {
	msg <- m
}

func (c *Client)Connect() {
	connect <- c
}

func (c *Client)Disconnect() {
	disconnect <- c
}
