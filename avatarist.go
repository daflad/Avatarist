package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

func main() {
	avatar := initAvatar(200, 200, 20, 20, 2)
	fmt.Println(avatar.String())
	avatar.GenerateRandomColours(98, 179, 229, 255, 175, 28)
	fmt.Println(avatar.String())
}

//Avatar The paramaters required to ceeate a random Avatar
type Avatar struct {
	Width      int
	Height     int
	BlockSize  int
	BorderSize int
	NumCols    int
	Img        *image.RGBA
	Colours    []color.RGBA
	BaseColour color.RGBA
}

//Set avatar sizes and init the image & colours
func initAvatar(width, height, blockSize, borderSize, numCols int) Avatar {
	var A Avatar
	A.Width = width
	A.Height = height
	A.BlockSize = blockSize
	A.BorderSize = borderSize
	A.NumCols = numCols
	A.Img = image.NewRGBA(image.Rect(0, 0, A.Width, A.Height))
	A.Colours = make([]color.RGBA, 0)
	return A
}

func (a *Avatar) String() string {
	return fmt.Sprintf("Width \t\t:: %v\nHeight \t\t:: %v\nBlockSize \t:: %v\n"+
		"BorderSize \t:: %v\nNumCols \t:: %v\nColours \t:: %v\nBaseColour \t:: %v\n", a.Width, a.Height, a.BlockSize,
		a.BorderSize, a.NumCols, a.Colours, a.BaseColour)
}

//GenerateRandomColours used in avatar.
func (a *Avatar) GenerateRandomColours(red, green, blue, alpha, offset, randomAmount int) {
	a.BaseColour = color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
	rand.Seed(time.Now().Unix())
	for i := 0; i < a.NumCols; i++ {
		r := uint8(rand.Intn(red) + rand.Intn(randomAmount) + offset)
		g := a.BaseColour.G + uint8(rand.Intn(randomAmount)+offset)
		b := a.BaseColour.B + uint8(rand.Intn(randomAmount)+offset)
		al := uint8(255)
		a.Colours = append(a.Colours, color.RGBA{r, g, b, al})
	}
}
