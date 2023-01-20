package main

import (
	"fmt"
	"math/rand"
	"time"

	_ "image/png"

	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"

	"CarGamePixel/cars"
	"CarGamePixel/serialcon"
	"CarGamePixel/visuals"

	gs "CarGamePixel/globalsettings"
)

// Main program to run the game
func run() {
	rand.Seed(time.Now().UnixNano())
	gs.SetupGlobalSettings()
	visuals.InitVisuals()

	if gs.S.Serial {
		serialcon.InitSerial()
	}

	cars.SetupCars()
	resetVariables()

	for {
		Draw()
	}
}

// Set/reset variables every time the game is started (and when restarted after game over)
func resetVariables() {
	gs.S.IncomingSpeed = 6
	cars.DodgedCars = 0
	cars.PrevDodged = 0
}

// Runs every frame
func Draw() {
	if gs.S.GameOver {
		visuals.DisplayGameOverText(cars.DodgedCars)
		// fmt.Println("Game Over")
		if gs.S.Win.JustPressed(pixelgl.KeySpace) {
			fmt.Println("Start over")
			cars.Cars = nil
			gs.S.GameOver = false
			resetVariables()
			cars.SetupCars()
		}
	} else {
		gs.S.Win.Clear(colornames.Azure) //see https://upload.wikimedia.org/wikipedia/commons/e/e7/SVG1.1_Color_Swatch.svg for possible colors

		cars.SpawnNewCars()

		cars.MoveAllComingCars()
		visuals.DrawAllComingCars()

		// Get input, either from keyboard or serial
		if gs.S.Serial {
			cars.MoveCarSerial()
		} else {
			cars.MovePlayerCarKeys()
		}

		// Draw the player car
		visuals.DrawCar(cars.PlayerCar)

		cars.CollisionDetection(cars.Cars, cars.PlayerCar)

		// Increase the speed every time a car was dodged
		if cars.PrevDodged != cars.DodgedCars {
			cars.IncreaseSpeed()
			cars.PrevDodged = cars.DodgedCars
		}

	}
	gs.S.Win.Update()
}

func main() {
	pixelgl.Run(run)
}
