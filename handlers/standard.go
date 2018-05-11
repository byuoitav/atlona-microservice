package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	help "github.com/byuoitav/atlona-microservice/helpers"
	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/fatih/color"
	"github.com/labstack/echo"
)

type Response struct {
	errorResp    bool
	errorMessage string
}

// Switching between possible inputs for any device
func SwitchInput(context echo.Context) error {
	log.Printf("Changing inputs")

	//Set parameters
	inputall := context.Param("input")
	//output := IncPort(context.Param("output"))
	address := context.Param("address")

	// Split out the video and audio streams
	finput := strings.Split(inputall, "!")
	video_input := finput[0]
	audio_input := finput[1]

	//Print out the command
	log.Printf("Routing %v and %v on %v", video_input, audio_input, address)

	//Call SwitchInput from helpers
	//resp, err := help.SwitchInput(Video_input, Audio_input, Address)
	resp, err := help.SwitchInput(video_input, audio_input, address)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	var mresp Response
	if resp != "" {
		err = json.Unmarshal([]byte(resp), &mresp)
		if err != nil {
			log.Printf(color.HiRedString("Error: %v", err))
			return context.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	if mresp.errorResp == true {
		log.Printf("There was a problem switching: %v", mresp.errorMessage)
		return context.JSON(http.StatusInternalServerError, mresp.errorMessage)
	}
	log.Printf("Success")
	return context.JSON(http.StatusOK, se.Input{Input: fmt.Sprintf("%s:%s", address, inputall)})
}

func CheckInput(context echo.Context) error {
	//Set params, only need output to verify which input is connected.
	address := context.Param("address")
	//inputall := context.Param("input")

	log.Printf("Verifiying which inputs are connected to %v decoder", address)

	//Call help.GetInput func, it will return the input as a string
	input, err := help.GetInput(address)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Success")
	return context.JSON(http.StatusOK, se.Input{Input: fmt.Sprintf("%s:%s", address, input)})
}

// Endpoint to determine the status of all inputs and outputs. 4 total inputs, 5 total outputs.
/*func AllInputs(context echo.Context) error {
	address := context.Param("address")
	feedback := make(map[string]string)
	log.Printf("Verifying the input of all output ports.")
	for i := 1; i <= 5; i++ {
		out := strconv.Itoa(i)
		in, err := help.GetInput(address, out)
		if err != nil {
			log.Printf("There was a problem: %v", err.Error())
			return context.JSON(http.StatusInternalServerError, err.Error())
		}
		out = DecPort(out)
		in = DecPort(in)
		feedback[out] = in
	}
	log.Printf("Success")
	return context.JSON(http.StatusOK, feedback)
}
*/

// Endpoint to reboot the Atlona decoder
/*func Reboot(context echo.Context) error {
	address := context.Param("address")
	log.Printf("Rebooting the decoder %v", address)
	reboot, err := help.DecoderReboot(address)
	if err != nil {
		log.Printf("There was a problem rebooting the decoder %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Printf("Success")
	return context.JSON(http.StatusOK)
}*/
