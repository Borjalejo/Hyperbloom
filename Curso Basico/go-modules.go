package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	//Instanciar Echo
	e := echo.New()

	//Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Aloha Mundo")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
