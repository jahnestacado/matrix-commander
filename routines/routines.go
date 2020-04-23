package routines

import (
	"fmt"
	"image/color"
	"matrix-commander/config"
	"sync"
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

	var mutex = &sync.Mutex{}

	var r1 Routine
	m, err := rgbmatrix.NewRGBLedMatrix(config.Matrix.GetConfig())
	if err != nil {
		panic(err)
	}
	c := rgbmatrix.NewCanvas(m)

	bounds := c.Bounds()
	r1 = Routine{
		ID: ID("r1"),
		Action: func(channel chan ID) {

			start := func() {

				for {

					for x := 0; x < bounds.Max.Y; x++ {
						for y := 0; y < bounds.Max.X; y++ {

							select {
							case emittedID := <-channel:
								fmt.Println("received", emittedID)
								if emittedID != r1.ID {
									return
								}
							default:
								mutex.Lock()

								c.Set(y, x, color.RGBA{25, 255, 255, 255})
								c.Render()
								mutex.Unlock()

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

				for {

					for x := 0; x < bounds.Max.Y; x++ {
						for y := 0; y < bounds.Max.X; y++ {

							select {
							case emittedID := <-channel:
								fmt.Println("received", emittedID)
								if emittedID != r2.ID {

									return
								}
							default:
								mutex.Lock()
								c.Set(y, x, color.RGBA{250, 0, 0, 255})
								c.Render()
								mutex.Unlock()
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
