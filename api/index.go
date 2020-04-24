package api

import (
	"fmt"
	"matrix-commander/routine"
	"net/http"

	"github.com/labstack/echo"
)

func Attach(server *echo.Echo, routines []routine.Routine) {
	channels := make([]chan routine.ID, 0)

	for _, r := range routines {
		channel := make(chan routine.ID)
		r.Apply(channel)
		channels = append(channels, channel)
	}

	server.POST("/api/pattern/:id", func(context echo.Context) error {

		ID := routine.ID(context.Param("id"))
		for i := range channels {
			fmt.Println("sending id", ID)
			channels[i] <- ID
		}

		return context.JSON(http.StatusOK, ID)
	})

}
