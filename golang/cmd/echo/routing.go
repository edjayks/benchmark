package myecho

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type User struct {
	Name  string `json:name`
	Email string `json:email`
}

func ConfigureRoutes(e *echo.Echo) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	e.GET("/echo", func(c echo.Context) error {
		log.Print("I am Ok.")
		return c.String(http.StatusOK, "I am Ok.")
	})

	e.POST("/simple/call", func(c echo.Context) error {
		var user User
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		log.Printf("Received user to save: %v", user)
		// log.Print("Received user to save: \n")

		return c.JSON(http.StatusOK, user)
	})
}
