package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	avatar := initAvatar(200, 200, 20, 20, 2)
	fmt.Println(avatar)
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
	return fmt.Sprintf("Width \t:: %v\nHeight \t:: %v\nBlockSize \t:: %v\n"+
		"BorderSize \t:: %v\nNumCols \t:: %v\n", a.Width, a.Height, a.BlockSize,
		a.BorderSize, a.NumCols)
}
