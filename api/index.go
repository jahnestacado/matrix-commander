package api

import (
	"matrix-commander/routine"
	"net/http"

	"github.com/labstack/echo"
)

var resetID = routine.ID("")

func Attach(server *echo.Echo, routines []routine.Routine) {
	channels := make([]chan routine.ID, 0)

	for _, r := range routines {
		channel := make(chan routine.ID, 1)
		r.Apply(channel)
		channels = append(channels, channel)
	}

	server.POST("/api/pattern/:id", func(context echo.Context) error {
		ID := routine.ID(context.Param("id"))
		for i := range channels {
			channels[i] <- resetID
			channels[i] <- ID
		}

		return context.JSON(http.StatusOK, ID)
	})

}
