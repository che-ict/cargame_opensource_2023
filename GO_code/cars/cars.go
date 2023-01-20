package cars

import (
	"fmt"
	"math/rand"

	gs "CarGamePixel/globalsettings"
	"CarGamePixel/serialcon"

	"github.com/faiface/pixel/pixelgl"
)

var Cars []*Car

var PlayerCar *Car
var comingCar *Car
var DodgedCars int
var PrevDodged int

type Car struct {
	Xpos      float64
	Ypos      float64
	Speed     float64
	Direction string
}

func SetupCars() {
	PlayerCar = initPlayerCar()
	comingCar = InitComingCar()
	Cars = append(Cars, comingCar)
}

// Initializes a coming car
func InitComingCar() *Car {
	return &Car{
		Xpos:      rand.Float64()*(gs.S.ScreenWidth-gs.S.CarSize) + gs.S.CarSize/2,
		Ypos:      gs.S.ScreenHeight,
		Speed:     gs.S.IncomingSpeed,
		Direction: "down",
	}
}

// Initializes the player car in the middle of the screen, at the bottom
func initPlayerCar() *Car {
	return &Car{
		Xpos:      gs.S.ScreenWidth/2 - gs.S.CarSize/2,
		Ypos:      0 + gs.S.CarSize/2,
		Speed:     15,
		Direction: "right",
	}
}

func SpawnNewCars() {
	// Init cars if there are less than 5
	if len(Cars) < 5 && Cars[len(Cars)-1].Ypos < gs.S.InitHeight {
		comingCar := InitComingCar()
		Cars = append(Cars, comingCar)
	}
}

// Increases the speed of the coming cars
func IncreaseSpeed() {
	gs.S.IncomingSpeed += 0.2
	for _, c := range Cars {
		if c.Direction == "down" {
			c.Speed = gs.S.IncomingSpeed
		}
	}
}

func MoveAllComingCars(){
	for _, car := range Cars {
		moveCar(car)
	}
}

// Function to move all the present cars
func moveCar(c *Car) {
	switch c.Direction {
	// Coming cars
	case "down":
		if (c.Ypos - gs.S.CarSize/2) > 0-gs.S.CarSize {
			c.Ypos -= c.Speed
		} else {
			Cars = Cars[1:]
			DodgedCars++
		}

	// Player car
	case "left":
		if !(c.Xpos-gs.S.CarSize/2 < 0) {
			c.Xpos -= c.Speed
		}
	case "right":
		if !(c.Xpos+gs.S.CarSize/2 > gs.S.ScreenWidth) {
			c.Xpos += c.Speed
		}
	}
}

// Checks if the player car has collided with one of the coming cars
func CollisionDetection(cars []*Car, playerCar *Car) {
	for _, car := range cars {
		if playerCar.Xpos > (car.Xpos-gs.S.CarSize) && playerCar.Xpos < (car.Xpos+gs.S.CarSize) && playerCar.Ypos > (car.Ypos-(gs.S.CarSize*(145.0/190.0))) && playerCar.Ypos < (car.Ypos+(gs.S.CarSize*(145.0/190.0))) {
			fmt.Println("Collision detected")
			gs.S.GameOver = true
		}
	}
}

// Move the player car to play the game with keyboard keys
func MovePlayerCarKeys() {
	if gs.S.Win.Pressed(pixelgl.KeyLeft) {
		PlayerCar.Direction = "left"
		moveCar(PlayerCar)
	}
	if gs.S.Win.Pressed(pixelgl.KeyRight) {
		PlayerCar.Direction = "right"
		moveCar(PlayerCar)
	}
}


// Move the players car to play the game with sensors in Arduino
func MoveCarSerial() {
	//ser := serialcon.GetSerialInput()
	select {
	case ser := <- serialcon.SerialInput: // Read the serial channel
		if ser < serialcon.PrevSer {
			PlayerCar.Direction = "left"
			moveCar(PlayerCar)
			serialcon.PrevSer = ser
		}
		if ser > serialcon.PrevSer {
			PlayerCar.Direction = "right"
			moveCar(PlayerCar)
			serialcon.PrevSer = ser
		}
	default:
		// Continue if the channel is empty
		return
	}

}