package routines

import (
	"fmt"
	"image/color"
	"matrix-commander/config"
	"time"

	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"
)

type ID string

type Routine struct {
	ID     ID
	Action func(chan ID)
}

func (r *Routine) Apply(channel chan ID) {
	fmt.Println("Apply Action", channel)

	go r.Action(channel)
}

var Routines []Routine

func init() {
	var r1 Routine
	r1 = Routine{
		ID: ID("r1"),
		Action: func(channel chan ID) {

			start := func() {
				m, err := rgbmatrix.NewRGBLedMatrix(config.Matrix.GetConfig())
				if err != nil {
					panic(err)
				}
				c := rgbmatrix.NewCanvas(m)

				defer c.Close()

				bounds := c.Bounds()

				for {

					for x := bounds.Min.X; x < bounds.Max.X; x++ {
						for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

							select {
							case emittedID := <-channel:
								fmt.Println("received", emittedID)
								if emittedID != r1.ID {
									c.Close()

									return
								}
							default:
								c.Set(y, x, color.RGBA{25, 255, 255, 255})
								c.Render()
								time.Sleep(50 * time.Millisecond)
							}
						}
					}

				}
			}
			for {
				select {
				case emittedID := <-channel:
					fmt.Println("received", emittedID)
					if emittedID == r1.ID {
						start()
					}
				}
			}

		},
	}

	var r2 Routine
	r2 = Routine{
		ID: ID("r2"),
		Action: func(channel chan ID) {

			start := func() {
				m, err := rgbmatrix.NewRGBLedMatrix(config.Matrix.GetConfig())
				if err != nil {
					panic(err)
				}
				c := rgbmatrix.NewCanvas(m)

				defer c.Close()

				bounds := c.Bounds()

				for {

					for x := bounds.Min.X; x < bounds.Max.X; x++ {
						for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

							select {
							case emittedID := <-channel:
								fmt.Println("received", emittedID)
								if emittedID != r2.ID {
									c.Close()
									return
								}
							default:
								c.Set(y, x, color.RGBA{250, 0, 0, 255})
								c.Render()
								time.Sleep(50 * time.Millisecond)
							}
						}
					}

				}
			}
			for {
				select {
				case emittedID := <-channel:
					fmt.Println("received", emittedID)
					if emittedID == r2.ID {
						start()
					}
				}
			}

		},
	}

	Routines = append(Routines, r1)
	Routines = append(Routines, r2)

	fmt.Println(len(Routines))
}
