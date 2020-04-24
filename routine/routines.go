package routine

import (
	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"
)

type ID string

type Routine struct {
	MatrixController rgbmatrix.Matrix
	ID               ID
	Action           func(ID, chan ID, rgbmatrix.Matrix)
}

func (r *Routine) Apply(channel chan ID) {
	go r.Action(r.ID, channel, r.MatrixController)
}

func CreateRoutines(matrixController rgbmatrix.Matrix) []Routine {
	var routines []Routine
	imagePlayerRoutine := Routine{
		MatrixController: matrixController,
		ID:               ID("image-player"),
		Action:           imagePlayerAction,
	}
	gameOfLifeRoutine := Routine{
		MatrixController: matrixController,
		ID:               ID("gol"),
		Action:           gameOfLifeAction,
	}

	routines = append(routines, imagePlayerRoutine)
	routines = append(routines, gameOfLifeRoutine)

	return routines

}
