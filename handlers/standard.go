package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	//"strconv"

	help "github.com/byuoitav/atlona-microservice/helpers"
	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/labstack/echo"
)

// Switching between possible inputs for any device
func SwitchInput(context echo.Context) error {
	log.Printf("Changing inputs")

	//Set parameters
	inputall := context.Param("input")
	//output := IncPort(context.Param("output"))
	Address := context.Param("address")

	// Split out the video and audio streams
	finput := strings.Split(inputall, "!")
	Video_input := finput[0]
	Audio_input := finput[1]

	//Print out the command
	log.Printf("Routing %v and %v on %v", Video_input, Audio_input, Address)

	//Call SwitchInput from helpers
	//resp, err := help.SwitchInput(Video_input, Audio_input, Address)
	err := help.SwitchInput(Video_input, Audio_input, Address)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	} /* else if resp.Error == true {
		log.Printf("There was a problem: %v", resp.ErrorMessage)
	}*/
	log.Printf("Success")
	return context.JSON(http.StatusOK, se.Input{Input: fmt.Sprintf("%s:%s", Address, inputall)})
}

func CheckInput(context echo.Context) error {
	//Set params, only need output to verify which input is connected.
	address := context.Param("address")
	inputall := context.Param("input")

	//log.Printf("Verifiying which input is connected to %v", output)

	//Call help.GetInput func, it will return the input as a string
	//input, err := help.GetInput(address, output)
	//if err != nil {
	//	log.Printf("There was a problem: %v", err.Error())
	//	return context.JSON(http.StatusInternalServerError, err.Error())
	//}

	log.Printf("Success")
	return context.JSON(http.StatusOK, se.Input{Input: fmt.Sprintf("%s:%s", address, inputall)})
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
