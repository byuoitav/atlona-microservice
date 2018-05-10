package helpers

import (
	"encoding/json"
	//"errors"
	"fmt"
	"log"
	//"net"
	"os"
	//"time"

	"github.com/gorilla/websocket"
)

var (
	envuser     = os.Getenv("USER")
	envpassword = os.Getenv("PASSWORD")
)

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SwitchConfigSet struct {
	Name   string   `json:"name"`
	Config []Config `json:"config"`
}

type Config struct {
	Multicast `json:"multicast"`
	Name      string `json:"name"`
}

type Multicast struct {
	Address string `json:"address"`
}

type Command struct {
	Creds
	SwitchConfigSet `json:"config_set"`
}

func SwitchInput(Video_input string, Audio_input string, Address string) error {
	// Building JSON Query
	Fig := []Config{Config{Multicast: Multicast{Address: Video_input}, Name: "ip_input1"}, Config{Multicast: Multicast{Address: Audio_input}, Name: "ip_input3"}}
	SC := Command{Creds: Creds{Username: envuser, Password: envpassword}, SwitchConfigSet: SwitchConfigSet{Name: "ip_input", Config: Fig}}
	fmt.Println(SC)
	Comm, err := json.Marshal(SC)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output to Console the factored JSON
	m := string(Comm)
	fmt.Println(m)

	// Get a new websocket connection to the decoder
	conn := OpenConnection(Address, Comm)
	return nil
}

//func
// func ChangePower(address, command string) string {
//   conn := getConnection(address)
//   command += string(CARRIAGE_RETURN)
//   conn.Write([]byte(command))
//   resp, err := readUntil(CARRIAGE_RETURN, conn, 1)
//   if err != nil {
//      log.Printf("Maybe it didn't read")
//   }
//   s := string(resp)
//   conn.Close()
//   return s
// }
