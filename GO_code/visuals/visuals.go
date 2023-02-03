package visuals

import (
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"CarGamePixel/cars"
	gs "CarGamePixel/globalsettings"
)

var basicTxt *text.Text
var basicAtlas *text.Atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

var carImg *pixel.Sprite
var carScale float64

func InitVisuals() {
	basicTxt = text.New(gs.S.Win.Bounds().Center(), basicAtlas) //Create a text element with basicAtlas type

	pic, err := loadPicture("car.png") //Load the picture from file
	if err != nil {
		panic(err)
	}

	carImg = pixel.NewSprite(pic, pic.Bounds()) // Create a car sprite
	carScale = gs.S.CarSize / pic.Bounds().W()
}

// Loads the car picture from the file, and returns a pixel.Picture element
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// Displays game over text
func DisplayGameOverText(dodged int) {
	basicTxt.Clear()
	basicTxt.Color = colornames.Darkcyan
	line := fmt.Sprintf("Je hebt %d auto's ontweken", dodged)

	fmt.Fprintln(basicTxt, "Game over!")
	fmt.Fprintf(basicTxt, line)

	basicTxt.Dot.X -= (basicTxt.BoundsOf(line).W() / 2)
	basicTxt.Draw(gs.S.Win, pixel.IM.Scaled(basicTxt.Dot, 3))

	basicTxt.Clear()
	lineSpace := "Druk op de spatiebalk om opnieuw te spelen"
	fmt.Fprintln(basicTxt, lineSpace)
	fmt.Fprintln(basicTxt, "Typ CTRL+C om het spel te sluiten")
	basicTxt.Dot.X += (basicTxt.BoundsOf(lineSpace).W() - 50)
	basicTxt.Dot.Y += basicTxt.BoundsOf(line).H() + 100
	basicTxt.Draw(gs.S.Win, pixel.IM.Scaled(basicTxt.Dot, 2.7))
}

// Scales the car img to a smaller size. Based on the global variable carScale
func getMatrixToMove(x, y float64) pixel.Matrix {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, carScale)
	mat = mat.Moved(pixel.V(x, y))
	return mat
}

func DrawCar(car *cars.Car) {
	carImg.Draw(gs.S.Win, getMatrixToMove(car.Xpos, car.Ypos))
}

// Move and draw all the coming cars
func DrawAllComingCars() {
	for _, car := range cars.Cars {
		DrawCar(car)
	}
}
