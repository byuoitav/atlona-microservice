package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	//"github.com/byuoitav/event-router-microservice/base"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

var (
	EnvUser     = os.Getenv("ATLONA_USERNAME")
	EnvPassword = os.Getenv("ATLONA_PASSWORD")
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the router.
	pingWait = 90 * time.Second

	// Interval to wait between retry attempts
	retryInterval = 3 * time.Second
)

func OpenConnection(Address string, Comm []byte) (string, error) {
	// Open connection to the decoder
	dialer := &websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := dialer.Dial(fmt.Sprintf("ws://%s/wsapp/", Address), nil)
	if err != nil {
		log.Printf(color.HiRedString("There was a problem establishing the websocket with 192.168.0.7: %v", err.Error()))
		return "", err
	}
	// Write JSON to Connected Decoder
	err = conn.WriteMessage(websocket.TextMessage, Comm)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Read Back any message that is returned from Writing the Message
	_, msgd, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	Msg := string(msgd)
	// pass the return message back up to handler to build the final failure

	fmt.Println(Msg)
	conn.Close()

	return Msg, err
}
