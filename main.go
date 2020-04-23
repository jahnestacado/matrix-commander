package main

import (
	"fmt"
	"matrix-commander/api"

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

	// server.HTTPErrorHandler = func(err error, context echo.Context) {
	// 	httpError := err.(*echo.HTTPError)
	// 	exposedError := errors.Unwrap(err)
	// 	if exposedError == nil {
	// 		exposedError = echo.NewHTTPError(httpError.Code, errors.New("error"))
	// 	}
	//
	// 	server.DefaultHTTPErrorHandler(exposedError, context)
	// }
	fmt.Println("wtf")

	api.Attach(server)
	fmt.Println("listening")

	server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%s", host, port)))
}
