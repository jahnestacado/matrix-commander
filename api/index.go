package api

import (
	"fmt"
	"net/http"

	"matrix-commander/routines"

	"github.com/labstack/echo"
)

func Attach(server *echo.Echo) {
	channels := make([]chan routines.ID, 0)

	for _, r := range routines.Routines {
		channel := make(chan routines.ID)
		r.Apply(channel)
		channels = append(channels, channel)
	}

	fmt.Println(channels)
	fmt.Println(channels[0], len(channels))

	server.POST("/api/pattern/:id", func(context echo.Context) error {

		ID := routines.ID(context.Param("id"))
		for i := range channels {
			fmt.Println("sending id", ID)
			channels[i] <- ID
		}

		return context.JSON(http.StatusOK, ID)
	})

}
