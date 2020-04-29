package routine

import (
	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"
)

type ID string

type Routine struct {
	MatrixConfig *rgbmatrix.HardwareConfig
	ID           ID
	Action       func(ID, chan ID, *rgbmatrix.HardwareConfig)
}

func (r *Routine) Apply(channel chan ID) {
	go r.Action(r.ID, channel, r.MatrixConfig)
}

func CreateRoutines(config *rgbmatrix.HardwareConfig) []Routine {
	var routines []Routine
	imagePlayerRoutine := Routine{
		MatrixConfig: config,
		ID:           ID("images"),
		Action:       imagePlayerAction,
	}
	gameOfLifeRoutine := Routine{
		MatrixConfig: config,
		ID:           ID("gol"),
		Action:       gameOfLifeAction,
	}

	routines = append(routines, imagePlayerRoutine)
	routines = append(routines, gameOfLifeRoutine)

	return routines

}
