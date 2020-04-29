package routine

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"time"

	life "github.com/jahnestacado/go-life"
	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"
)

const (
	opacity                       = 255
	numOfSeeds                    = 290
	gifDisplayDurationInSecs      = 30
	golGenerationIntervalInMillis = 500
)

var (
	redColor   = color.RGBA{255, 0, 0, opacity}
	cyanColor  = color.RGBA{25, 255, 255, opacity}
	greenColor = color.RGBA{0, 255, 0, opacity}
	blackColor = color.RGBA{0, 0, 0, opacity}
)

func imagePlayerAction(routineID ID, channel chan ID, matrixConfig *rgbmatrix.HardwareConfig) {
	start := func(toolkit *rgbmatrix.ToolKit) {
		for {

			gifFiles, err := ioutil.ReadDir("./images")
			if err != nil {
				log.Fatal(err)
			}

			for _, gifFile := range gifFiles {

				select {
				case emittedID := <-channel:
					if emittedID != routineID {
						toolkit.Close()
						return
					}
				default:
					gif, err := os.Open(fmt.Sprintf("./images/%s", gifFile.Name()))
					if err != nil {
						panic(err)
					}

					close, _ := toolkit.PlayGIF(gif)

					time.Sleep(time.Second * gifDisplayDurationInSecs)
					close <- true
				}

			}

		}
	}

	for emittedID := range channel {
		if emittedID == routineID {
			matrixController, err := rgbmatrix.NewRGBLedMatrix(matrixConfig)
			if err != nil {
				panic(err)
			}
			toolKit := rgbmatrix.NewToolKit(matrixController)
			start(toolKit)
		}
	}
}

func gameOfLifeAction(routineID ID, channel chan ID, matrixConfig *rgbmatrix.HardwareConfig) {
	var canvas *rgbmatrix.Canvas
	start := func(canvas *rgbmatrix.Canvas) {
		bounds := canvas.Bounds()

		config := life.Config{
			NumOfRows:  bounds.Max.Y,
			NumOfCols:  bounds.Max.X,
			NumOfSeeds: numOfSeeds,
		}
		l := life.New(config)

		for {
			l.Next()
			grid := l.GetGrid()
			for x := 0; x < bounds.Max.Y; x++ {
				for y := 0; y < bounds.Max.X; y++ {

					select {
					case emittedID := <-channel:
						if emittedID != routineID {
							canvas.Close()
							return
						}
					default:
						cell := grid[x][y]

						cellColor := blackColor
						if cell.Color == "green" {
							cellColor = greenColor
						}
						if cell.Color == "cyan" {
							cellColor = cyanColor
						}
						if cell.Color == "red" {
							cellColor = redColor
						}
						canvas.Set(y, x, cellColor)
					}
				}
			}
			canvas.Render()
			time.Sleep(time.Millisecond * golGenerationIntervalInMillis)
		}
	}

	for emittedID := range channel {
		if emittedID == routineID {
			matrixController, err := rgbmatrix.NewRGBLedMatrix(matrixConfig)
			if err != nil {
				panic(err)
			}
			canvas = rgbmatrix.NewCanvas(matrixController)
			start(canvas)
		}
	}
}
