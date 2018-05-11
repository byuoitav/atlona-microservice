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
	userName       string `json:"username"`
	passWord       string `json:"password"`
	switchInputGet string `json:"config_get"`
}

func GetInput(address string) (string, error) {
	gi := giCommand{userName: EnvUser, passWord: EnvPassword, switchInputGet: "ip_input"}
	fmt.Printf("output: %v/n", gi)
	comm, err := json.Marshal(gi)
	test := string(comm)
	fmt.Printf("Comm Output: %v/n", test)
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
