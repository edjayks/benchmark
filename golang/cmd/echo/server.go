package myecho

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func StartEchoServer() {
	// UNIX Time is faster and smaller than most timestamps
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	e := echo.New()

	// TODO: add middleware such as auth, logs and etc

	// add routes and handlers
	ConfigureRoutes(e)

	// TODO: configure server, should load parameters from env var
	s := http.Server{
		Addr:         ":8080",
		Handler:      e,
		ReadTimeout:  30 * time.Second, // customize http.Server timeouts
		WriteTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	fmt.Printf("⇨ http server started on %s\n", s.Addr)

	// start the server
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
