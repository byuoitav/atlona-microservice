package helpers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
)

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

type SICommand struct {
	Creds
	SwitchConfigSet `json:"config_set"`
}

func SwitchInput(Video_input string, Audio_input string, Address string) (string, error) {
	// Building JSON Query
	fig := []Config{Config{Multicast: Multicast{Address: Video_input}, Name: "ip_input1"}, Config{Multicast: Multicast{Address: Audio_input}, Name: "ip_input3"}}
	SC := SICommand{Creds: Creds{Username: EnvUser, Password: EnvPassword}, SwitchConfigSet: SwitchConfigSet{Name: "ip_input", Config: fig}}
	comm, err := json.Marshal(SC)
	if err != nil {
		fmt.Printf(color.HiRedString("Error:", err))
		return "", err
	}

	// Get a new websocket connection to the decoder
	resp, err := OpenConnection(Address, comm)
	if err != nil {
		log.Printf(color.HiRedString("Error:", err))
		return "", err
	}

	return resp, nil
}
