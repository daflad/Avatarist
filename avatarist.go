package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	avatar := initAvatar("MyFirstAvatar", 200, 200, 20, 20, 2)
	fmt.Println(avatar.String())
	avatar.GenerateRandomColours(98, 179, 229, 255, 175, 28)
	fmt.Println(avatar.String())
	avatar.BlankCanvas()
	avatar.Draw()
	avatar.Write()
}

//Avatar The paramaters required to ceeate a random Avatar
type Avatar struct {
	Name       string
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
func initAvatar(name string, width, height, blockSize, borderSize, numCols int) Avatar {
	var A Avatar
	A.Name = name
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
//Rnadom colours from a base colour
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

//BlankCanvas cover the image with the BaseColour
func (a *Avatar) BlankCanvas() {
	for y := a.Img.Rect.Min.Y; y < a.Img.Rect.Max.Y; y++ {
		for x := a.Img.Rect.Min.X; x < a.Img.Rect.Max.X; x++ {
			a.Img.Set(x, y, a.BaseColour)
		}
	}
}

//Draw the random avatar
func (a *Avatar) Draw() {
	//The colour to draw the current pixel
	var col color.RGBA
	top, bottom, left, right := a.InnerDimentions()
	//Scan image & update pixel colous
	//find corner of block and set the colour
	//change colour when new block is found
	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			if x%a.BlockSize == 0 {
				if y%a.BlockSize == 0 {
					col = a.Colours[rand.Intn(a.NumCols)]
				} else {
					if y > a.Img.Rect.Min.Y {
						r, g, b, a := a.Img.At(x, y-1).RGBA()
						col = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
					} else {
						col = a.Colours[rand.Intn(a.NumCols)]
					}
				}
			}
			a.Img.Set(x, y, col)
		}
	}
}

//InnerDimentions of randomised avatar
func (a *Avatar) InnerDimentions() (int, int, int, int) {
	top := a.Img.Rect.Min.Y + a.BorderSize
	bottom := a.Img.Rect.Max.Y - a.BorderSize
	left := a.Img.Rect.Min.X + a.BorderSize
	right := a.Img.Rect.Max.X - a.BorderSize
	return top, bottom, left, right
}

//Write the file to disk
func (a *Avatar) Write() {
	a.Name = a.Name + ".png"
	file, err := os.Create(a.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	png.Encode(file, a.Img)
}
