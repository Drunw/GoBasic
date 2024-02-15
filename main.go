package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var uuidString string

func main() {
	e := echo.New()
	e.POST("cambioNombre", func(c echo.Context) error {
		r := new(Request)
		if err := c.Bind(r); err != nil {
			return err
		}
		uuidString = uuid.New().String()
		imprimirLogEntradaApi(uuidString, r.String())
		return c.JSON(http.StatusCreated, convertirMayuscula(*r))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
