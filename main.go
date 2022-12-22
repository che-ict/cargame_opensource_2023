package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"

	"golang.org/x/image/colornames"

	"go.bug.st/serial"
)

var incomingSpeed float64 // wordt geinitieerd in de setup
var screenHeight float64 = 900
var screenWidth float64 = 900
var initHeight float64 = 550
var carSize float64 = 125
var carScale float64
var cars []*Car
var gameOver bool = false
var win *pixelgl.Window
var cfg pixelgl.WindowConfig
var carImg *pixel.Sprite
var playerCar *Car
var comingCar *Car
var dodgedCars int
var prevDodged int

// var port serial.Port
var prevSer int = 100

//var basicTxt *text.Text
//var basicAtlas *text.Atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

var serialInput chan int = make(chan int)

type Car struct {
	xpos      float64
	ypos      float64
	speed     float64
	direction string
}

func initComingCar() *Car {
	return &Car{
		//xpos:      screenWidth/2 - carSize/2,
		xpos:      rand.Float64()*(screenWidth-carSize) + carSize/2,
		ypos:      screenHeight,
		speed:     incomingSpeed,
		direction: "down",
	}
}

func initPlayerCar() *Car {
	return &Car{
		xpos:      screenWidth/2 - -carSize/2,
		ypos:      0 + carSize/2,
		speed:     15,
		direction: "right",
	}
}

func increaseSpeed() {
	incomingSpeed += 0.2
	for _, c := range cars {
		if c.direction == "down" {
			c.speed = incomingSpeed
		}
	}
}

func moveCar(c *Car) {
	switch c.direction {
	case "down":
		if (c.ypos - carSize/2) > 0-carSize {
			c.ypos -= c.speed
		} else {
			cars = cars[1:]
			dodgedCars++
			//c.ypos = screenHeight
			//c.xpos = rand.Float64() * (screenWidth - carSize/2)
		}
	case "left":
		if !(c.xpos-carSize/2 < 0) {
			c.xpos -= c.speed
		}
	case "right":
		if !(c.xpos+carSize/2 > screenWidth) {
			c.xpos += c.speed
		}
	}
}

func collisionDetection(cars []*Car, playerCar *Car) {
	for _, car := range cars {
		if playerCar.xpos > (car.xpos-carSize) && playerCar.xpos < (car.xpos+carSize) && playerCar.ypos > (car.ypos-(carSize*(145.0/190.0))) && playerCar.ypos < (car.ypos+(carSize*(145.0/190.0))) {
			fmt.Println("Collision detected")
			gameOver = true
		}
	}
}

func run() {
	var err error

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	port, err = serial.Open(getPorts(), mode)
	if err != nil {
		log.Fatal(err)
	}
	go getSerialInput()

	time.Sleep(time.Second * 2)

	rand.Seed(time.Now().UnixNano())
	cfg = pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	basicTxt = text.New(win.Bounds().Center(), basicAtlas)

	pic, err := loadPicture("car2.png")
	if err != nil {
		panic(err)
	}

	carImg = pixel.NewSprite(pic, pic.Bounds())
	carScale = carSize / pic.Bounds().W()

	Setup()

	for {
		Draw()
	}
}

func moveCarKeys() {
	if win.Pressed(pixelgl.KeyLeft) {
		playerCar.direction = "left"
		moveCar(playerCar)
	}
	if win.Pressed(pixelgl.KeyRight) {
		playerCar.direction = "right"
		moveCar(playerCar)
	}
}

func moveCarSerial() {
	//ser := getSerialInput()
	select {
	case ser := <-serialInput:
		// if ser != -1 {
		if ser < prevSer {
			playerCar.direction = "left"
			moveCar(playerCar)
			prevSer = ser
		}
		if ser > prevSer {
			playerCar.direction = "right"
			moveCar(playerCar)
			prevSer = ser
		}
		// }
	default:
		// fmt.Println("No data available")
		return
	}

}

func Setup() {
	incomingSpeed = 6
	playerCar = initPlayerCar()
	comingCar = initComingCar()
	cars = append(cars, comingCar)

	dodgedCars = 0
	prevDodged = 0

}

func Draw() {
	if gameOver {
		displayText()
		// fmt.Println("Game Over")
		if win.JustPressed(pixelgl.KeySpace) {
			fmt.Println("Start over")
			cars = nil
			gameOver = false
			Setup()
		}
	} else {
		//dt := time.Since(last).Seconds()
		//last = time.Now()

		win.Clear(colornames.Azure) //see https://upload.wikimedia.org/wikipedia/commons/e/e7/SVG1.1_Color_Swatch.svg for possible colors
		if len(cars) < 5 && cars[len(cars)-1].ypos < initHeight {
			comingCar = initComingCar()
			cars = append(cars, comingCar)
		}

		for _, car := range cars {
			moveCar(car)
			carImg.Draw(win, getMatrixToMove(car.xpos, car.ypos))
		}

		//moveCarKeys()
		moveCarSerial()

		carImg.Draw(win, getMatrixToMove(playerCar.xpos, playerCar.ypos))

		collisionDetection(cars, playerCar)
		//fmt.Println("cars", cars)

		if prevDodged != dodgedCars {
			increaseSpeed()
			//fmt.Println("dodgedCars", dodgedCars)
			//fmt.Println("incomingSpeed", incomingSpeed)
			prevDodged = dodgedCars
		}

	}
	win.Update()
}

func getMatrixToMove(x, y float64) pixel.Matrix {
	mat := pixel.IM
	//mat = mat.Scaled(pixel.ZV, 0.2)
	mat = mat.Scaled(pixel.ZV, carScale)
	mat = mat.Moved(pixel.V(x, y))
	return mat
}

func main() {
	pixelgl.Run(run)
}
