package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	address := context.Param("address")

	// Split out the video and audio streams
	finput := strings.Split(inputall, "!")
	video_input = finput[0]
	audio_input = finput[1]

	//Print out the command
	log.Printf("Routing %v and %v on %v", video_input, audio_input, address)

	//Call SwitchInput from helpers
	err := help.SwitchInput(address, video_input, audio_input)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Success")
	input = DecPort(input)
	output = DecPort(output)
	return context.JSON(http.StatusOK, se.Input{Input: fmt.Sprintf("%s:%s", input, output)})
}

func CheckInput(context echo.Context) error {
	//Set params, only need output to verify which input is connected.
	output := IncPort(context.Param("output"))
	address := context.Param("address")
	log.Printf("Verifiying which input is connected to %v", output)

	//Call help.GetInput func, it will return the input as a string
	input, err := help.GetInput(address, output)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Success")
	input = DecPort(input)
	output = DecPort(output)
	return context.JSON(http.StatusOK, se.Input{Input: fmt.Sprintf("%s:%s", input, output)})
}

// Endpoint to determine the status of all inputs and outputs. 4 total inputs, 5 total outputs.
func AllInputs(context echo.Context) error {
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

// Endpoint to reboot the Atlona decoder
func Reboot(context echo.Context) error {
	address := context.Param("address")
}

//Function to increment the port from a base 0 to a base 1 input/output
func IncPort(IOport string) string {
	x, err := strconv.Atoi(IOport)
	if err != nil {
		return ""
	}
	x++
	return strconv.Itoa(x)
}

//Function to decrement the port from a base 1 to a base 0 input/output
func DecPort(IOport string) string {
	x, err := strconv.Atoi(IOport)
	if err != nil {
		return ""
	}
	x--
	return strconv.Itoa(x)
}
