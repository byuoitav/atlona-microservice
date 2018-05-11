package helpers

import (
	"encoding/json"
	"fmt"
	//"log"
	//"os"

	"github.com/fatih/color"
	//"github.com/gorilla/websocket"
)

/*
type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
*/

type giCommand struct {
	Creds                 // Defined in common.get
	switchInputGet string `json:"config_get"`
}

func GetInput(address string) (string, error) {
	gi := giCommand{Creds: Creds{Username: EnvUser, Password: EnvPassword}, switchInputGet: "ip_input"}
	comm, err := json.Marshal(gi)
	if err != nil {
		fmt.Printf(color.HiRedString("Error:", err))
		return "", err
	}

	resp, err := OpenConnection(address, comm)
	if err != nil {
		fmt.Printf(color.HiRedString("Error Connecting to Decoder:", err))
		return "", err
	}

	return resp, nil
}

func testing() {
	fmt.Printf("Testing......")
}
