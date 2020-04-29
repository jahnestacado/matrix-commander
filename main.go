package main

import (
	"fmt"
	"matrix-commander/api"
	"matrix-commander/config"
	"matrix-commander/routine"

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

	matrixConfig := config.Matrix.GetConfig()
	routines := routine.CreateRoutines(matrixConfig)

	api.Attach(server, routines)

	server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%s", host, port)))
}
