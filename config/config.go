package config

import rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"

var Matrix matrixConfig

type matrixConfig struct {
	config rgbmatrix.HardwareConfig
}

func (m *matrixConfig) GetConfig() *rgbmatrix.HardwareConfig {
	return &m.config
}

func init() {
	c := &rgbmatrix.DefaultConfig
	c.Rows = 32
	c.Cols = 64
	c.Parallel = 1
	c.ChainLength = 1
	c.Brightness = 100
	c.HardwareMapping = "regular"
	c.ShowRefreshRate = false
	c.InverseColors = false
	c.DisableHardwarePulsing = false

	Matrix = matrixConfig{*c}
}
