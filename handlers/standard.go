package handlers

import (
	"log"
	"net/http"

	help "github.com/byuoitav/atlona-microservice/helpers"
	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/labstack/echo"
)

func SwitchInput(context echo.Context) error {
	log.Printf("Changing inputs")
	//Set parameters
	input := context.Param("input")
	output := context.Param("output")
	address := context.Param("address")
	//Print out the command
	log.Printf("Routing %v to %v on %v", input, output, address)
	//Call SwitchInput from helpers
	err := help.SwitchInput(address, input, output)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Success")
	return context.JSON(http.StatusOK, statusevaluators.Input{Input: input})
}

func CheckInput(context echo.Context) error {
	//Set params, only need output to verify which input is connected.
	output := context.Param("output")
	address := context.Param("address")
	log.Printf("Verifiying which input is connected to %v", output)
	//Call help.GetInput func, it will return the input as a string
	resp, err := help.GetInput(address, output)
	if err != nil {
		log.Printf("There was a problem: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Success")
	return context.JSON(http.StatusOK, statusevaluators.Input{Input: resp})
}

// func PowerSwitch(address, command string) string {
//   response := help.ChangePower(address, command)
//   return response
// }
