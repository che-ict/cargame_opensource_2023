package globalsettings

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GlobalSettings struct {
	ScreenHeight  float64
	ScreenWidth   float64
	InitHeight    float64
	CarSize       float64
	IncomingSpeed float64
	Win           *pixelgl.Window
	GameOver      bool
	Serial        bool
}

var S *GlobalSettings

func SetupGlobalSettings() *GlobalSettings {
	S = &GlobalSettings{
		ScreenHeight:  900,
		ScreenWidth:   900,
		InitHeight:    550,
		CarSize:       125,
		IncomingSpeed: 6,
		Win:           nil,
		GameOver:      false,
		Serial:        true,
	}

	S.Win = windowSetup(S)
	return S
}

func windowSetup(s *GlobalSettings) *pixelgl.Window {
	// Create visual window
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, s.ScreenWidth, s.ScreenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	return win
}
