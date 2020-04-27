package main

import (
	"fmt"
	"matrix-commander/api"
	"matrix-commander/config"
	"matrix-commander/routine"

	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	host = "0.0.0.0"
	port = "5555"
)

func main() {
	server := echo.New()
	server.Pre(middleware.RemoveTrailingSlash())

	matrixController, err := rgbmatrix.NewRGBLedMatrix(config.Matrix.GetConfig())
	if err != nil {
		panic(err)
	}

	routines := routine.CreateRoutines(matrixController)

	api.Attach(server, routines)

	server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%s", host, port)))
}
