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
	GI := giCommand{userName: EnvUser, passWord: EnvPassword, switchInputGet: "ip_input"}
	fmt.Printf("output: %v\n", GI)
	comm, err := json.Marshal(GI)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %v\n", err))
		return "", err
	}
	//test := string(comm)
	fmt.Printf("Comm Output: %v\n", comm)
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
