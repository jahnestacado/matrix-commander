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

func init() {
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func imagePlayerAction(routineID ID, channel chan ID, matrixController rgbmatrix.Matrix) {
	start := func(toolkit *rgbmatrix.ToolKit) {
		for {

			files, err := ioutil.ReadDir("./images")
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {

				select {
				case emittedID := <-channel:
					fmt.Println("received", emittedID)
					if emittedID != routineID {
						toolkit.Close()
						return
					}
				default:
					gif, err := os.Open(fmt.Sprintf("./images/%s", file.Name()))
					if err != nil {
						panic(err)
					}

					// switch *rotate {
					// case 90:
					// 	tk.Transform = imaging.Rotate90
					// case 180:
					// 	tk.Transform = imaging.Rotate180
					// case 270:
					// 	tk.Transform = imaging.Rotate270
					// }

					close, _ := toolkit.PlayGIF(gif)

					time.Sleep(time.Second * 10)
					close <- true
				}

			}

		}
	}

	for emittedID := range channel {
		if emittedID == routineID {
			toolKit := rgbmatrix.NewToolKit(matrixController)
			start(toolKit)
		}
	}
}

func gameOfLifeAction(routineID ID, channel chan ID, matrixController rgbmatrix.Matrix) {
	const interval = 500
	var canvas *rgbmatrix.Canvas
	start := func(canvas *rgbmatrix.Canvas) {
		bounds := canvas.Bounds()

		config := life.Config{
			NumOfRows:  bounds.Max.Y,
			NumOfCols:  bounds.Max.X,
			NumOfSeeds: 290,
		}
		l := life.New(config)

		for {
			l.Next()
			grid := l.GetGrid()
			for x := 0; x < bounds.Max.Y; x++ {
				for y := 0; y < bounds.Max.X; y++ {

					select {
					case emittedID := <-channel:
						fmt.Println("received", emittedID)
						if emittedID != routineID {
							canvas.Close()
							return
						}
					default:
						cell := grid[x][y]

						cellColor := color.RGBA{0, 0, 0, 255}
						if cell.Color == "green" {
							cellColor = color.RGBA{0, 255, 0, 255}
						}
						if cell.Color == "cyan" {
							cellColor = color.RGBA{25, 255, 255, 255}
						}
						if cell.Color == "red" {
							cellColor = color.RGBA{255, 0, 0, 255}
						}
						canvas.Set(y, x, cellColor)
					}
				}
			}
			canvas.Render()
			time.Sleep(time.Millisecond * interval)
		}
	}

	for emittedID := range channel {
		if emittedID == routineID {
			canvas = rgbmatrix.NewCanvas(matrixController)
			start(canvas)
		}
	}
}
