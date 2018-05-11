package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

type giCommand struct {
	Creds
	SwitchInputGet string `json:"config_get"`
}

// Get a read out of the current configuration of the inputs for a decoder.
func GetInput(address string) (string, error) {
	gi := giCommand{Creds: Creds{Username: EnvUser, Password: EnvPassword}, SwitchInputGet: "ip_input"}
	comm, err := json.Marshal(gi)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %v\n", err))
		return "", err
	}
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
