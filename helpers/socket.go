package socket

import (
	"fmt"
	"log"
	"time"

	"github.com/byuoitav/event-router-microservice/base"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the router.
	pingWait = 90 * time.Second

	// Interval to wait between retry attempts
	retryInterval = 3 * time.Second
)

type Node struct {
	Name          string
	Conn          *websocket.Conn
	WriteQueue    chan base.Message
	ReadQueue     chan base.Message
	DecoderAddress string
	filters       map[string]bool
	readDone      chan bool
	writeDone     chan bool
	lastPingTime  time.Time
	state         string
}

func (n *Node) GetState() (string, interface{}) {
	values := make(map[string]interface{})

	values["router"] = n.DecoderAddress

	if n.Conn != nil {
		values["connection"] = fmt.Sprintf("%v => %v", n.Conn.LocalAddr().String(), n.Conn.RemoteAddr().String())
	} else {
		values["connection"] = fmt.Sprintf("%v => %v", "Local", n.DecoderAddress)
	}

	filters := []string{}
	for filter := range n.filters {
		filters = append(filters, filter)
	}
	values["filters"] = filters
	values["state"] = n.state
	values["last-ping-time"] = n.lastPingTime.Format(time.RFC3339)

	return n.Name, values
}

func (n *Node) Start(DecoderAddress string, filters []string, name string) error {

	log.Printf(color.HiGreenString("Connecting to decoder: %v", DecoderAddress))

	n.state = "initializing"
	n.DecoderAddress = DecoderAddress
	n.ReadQueue = make(chan base.Message, 4096)
	n.WriteQueue = make(chan base.Message, 4096)
	n.readDone = make(chan bool, 1)
	n.writeDone = make(chan bool, 1)
	n.filters = make(map[string]bool)
	n.Name = name

	for _, f := range filters {
		n.filters[f] = true
	}

	err := n.openConnection()
	if err != nil {
		log.Printf(color.YellowString("Opening connection failed, retrying..."))

		n.readDone <- true
		n.writeDone <- true

		go n.retryConnection()
		return nil
	}

	log.Printf(color.HiGreenString("Starting pumps..."))
	n.state = "good"
	go n.readPump()
	go n.writePump()

	return nil

}

func (n *Node) openConnection() error {
	//open connection to the decoder 
	dialer := &websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := dialer.Dial(fmt.Sprintf("ws://%s/wsapp", n.DecoderAddress), nil)
	if err != nil {
		log.Printf(color.HiRedString("There was a problem establishing the websocket with %v : %v", n.DecoderAddress, err.Error()))
		return err
	}
	n.Conn = conn

	return nil
}

func (n *Node) retryConnection() {

	//mark the connection as 'down'
	n.state = n.state + " retrying"

	log.Printf(color.HiMagentaString("[retry] Retrying connection, waiting for read and write pump to close before starting."))
	//wait for read to say i'm done.
	<-n.readDone
	log.Printf(color.HiMagentaString("[retry] Read pump closed"))

	//wait for write to be done.
	<-n.writeDone
	log.Printf(color.HiMagentaString("[retry] Write pump closed"))
	log.Printf(color.HiMagentaString("[retry] Retrying connection"))

	//we retry
	err := n.openConnection()

	for err != nil {
		log.Printf(color.HiMagentaString("[retry] Retry failed, trying again in 3 seconds."))
		time.Sleep(retryInterval)
		err = n.openConnection()
	}
	//start the pumps again
	log.Printf(color.HiGreenString("[Retry] Retry success. Starting pumps"))

	n.state = "good"
	go n.readPump()
	go n.writePump()

}

func (n *Node) readPump() {

	defer func() {
		n.Conn.Close()
		log.Printf(color.HiRedString("Connection to router %v is dying.", n.DecoderAddress))
		n.state = "down"

		n.readDone <- true
	}()

	n.Conn.SetPingHandler(
		func(string) error {
			log.Printf(color.HiCyanString("[%v] Ping! Ping was my best friend growin' up.", n.DecoderAddress))
			n.Conn.SetReadDeadline(time.Now().Add(pingWait))
			n.Conn.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(writeWait))

			//debugging purposes
			n.lastPingTime = time.Now()

			return nil
		})

	n.Conn.SetReadDeadline(time.Now().Add(pingWait))

	for {
		var message base.Message
		err := n.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		_, ok := n.filters[message.MessageHeader]
		if !ok {
			continue
		}

		//we need to check against the list of accepted values
		n.ReadQueue <- message
	}
}

func (n *Node) writePump() {

	defer func() {
		n.Conn.Close()
		log.Printf(color.HiRedString("Connection to router %v is dying. Trying to resurrect.", n.DecoderAddress))
		n.state = "down"

		n.writeDone <- true

		//try to reconnect
		n.retryConnection()
	}()

	for {
		select {
		case message, ok := <-n.WriteQueue:
			if !ok {
				n.Conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(writeWait))
				return
			}

			err := n.Conn.WriteJSON(message)
			if err != nil {
				return
			}
		case <-n.readDone:
			//put it back in
			n.readDone <- true
			return
		}
	}
}

func (n *Node) Write(message base.Message) error {
	n.WriteQueue <- message
	return nil

}

func (n *Node) Read() base.Message {
	msg := <-n.ReadQueue
	return msg
}

func (n *Node) Close() {

}
