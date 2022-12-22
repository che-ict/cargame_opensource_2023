package main

import (
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var basicTxt *text.Text
var basicAtlas *text.Atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

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

func displayText() {
	basicTxt.Clear()
	basicTxt.Color = colornames.Royalblue
	line := fmt.Sprintf("You have dodged %d cars", int(dodgedCars))
	//line := fmt.Sprintf("Game Over! You have dodged %d cars", int(dodgedCars))

	fmt.Fprintln(basicTxt, "Game over!")
	fmt.Fprintf(basicTxt, line)
	//basicTxt.Dot.X -= basicTxt.BoundsOf(line).W() / 2
	basicTxt.Dot.X -= (basicTxt.BoundsOf(line).W() / 2)
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Dot, 3))
	//basicTxt.Draw(win, pixel.IM)
}
