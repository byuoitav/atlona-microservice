package helpers

import (
	"fmt"
	//"log"
	//"github.com/gorilla/websocket"
)

/*
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

func GetInput(address, string) (string, error) {
	//command := ("Statusx" + output)
	conn := OpenConnection(address, Comm)
	//command += string(CARRIAGE_RETURN)
	//conn.Write([]byte(command))
	//resp, err := readUntil(CARRIAGE_RETURN, conn, 1)
	//if err != nil {
	//	log.Printf("%s", err)
	//}
	//log.Printf("Feedback: %s", resp)
	//s := string(resp[1])
	conn.Close()
	//return s, err
	return nil
}*/

func testing() {
	fmt.Printf("Testing......")
}
