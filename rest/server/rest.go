package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

const addr = ":4000"

func main() {
	log.Printf("server running %s\n", addr)

	e := echo.New()

	e.POST("/", func(c *echo.Context) error {
		var persons Persons
		var proccesed []Person

		if err := json.NewDecoder(c.Request().Body).Decode(&persons); err != nil {
			return c.String(http.StatusBadRequest, "error in the payload")
		}

		for _, person := range persons.Persons {
			proccesed = append(proccesed, person)
		}

		return c.String(http.StatusOK, fmt.Sprintf("%d records procced", len(proccesed)))
	})

	e.Start(addr)
}
