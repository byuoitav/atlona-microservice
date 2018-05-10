package main

import (
	//"log"
	"github.com/byuoitav/atlona-microservice/handlers"
	"github.com/byuoitav/atlona-microservice/handlersmatrix"
	"github.com/byuoitav/authmiddleware"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"time"
)

func init() {
	if os.Getenv("USERNAME") == "" {
		log.Fatal("$USER variable not set")
	}
	if os.Getenv("PASSWORD") == "" {
		log.Fatal("$PASSWORD variable not set")
	}
}

func main() {
	port := ":8015"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	// Functionality endpoints for Atlona Video over IP Switching
	secure.GET("/:address/input/:input", handlers.SwitchInput) //Format :input with 239.1.1.1!239.10.1.1
	secure.GET("/:address/status", handlers.CheckInput)
	//secure.GET("/:address/allstatus", handlers.AllInputs)
	//secure.GET("/:address/reboot", handlers.Reboot)

	// Functionality endpoints for Atlona Standard Switch
	secure.GET("/switch/:address/input/:input/:output", handlersmatrix.SwitchInput)
	secure.GET("/switch/:address/status/:output", handlersmatrix.CheckInput)
	secure.GET("/switch/:address/allstatus", handlersmatrix.AllInputs)

	// Status endpoints

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
