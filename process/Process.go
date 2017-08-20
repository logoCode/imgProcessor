package process

import (
	"extendedPlotter/data"
	"extendedPlotter/settings"
	"fmt"
	"image"
	"image/png"
	"os"
)

//create image
func CreateImg(s settings.Settings) error {
	//create data from file
	d := data.NewData()
	err := d.CreateFromFile(s.FilenameIn, s.Identifier, s.Separator, s.Accuracy)
	//check for any errors
	if err != nil {
		return err
	}
	//create rectangle with fitting dimensions (rectangle is used to build a rgbaImage)
	rect := image.Rectangle{image.Point{0, 0}, image.Point{s.Scaling * d.X, s.Scaling * d.Y}}
	//build image form rectangle
	img := image.NewRGBA(rect)

	//range over x and y coordinates of data.Img
	for x := range d.Img {
		for y := range d.Img[x] {
			for fx := 0; fx < s.Scaling; fx++ {
				for fy := 0; fy < s.Scaling; fy++ {
					//set the color of the current pixel according to its value in d.Img (if it doesn't exist, it is set to transparent)
					img.SetRGBA(s.Scaling*x+fx, s.Scaling*(d.Y-y-1)+fy, s.Colors[d.Img[x][y]])
				}
			}
		}
		fmt.Printf("\r%d%% finished", int(100*(x+1)/d.X))
	}
	fmt.Println()

	//finally save the image with the given name
	f, _ := os.Create(s.FilenameOut)
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}
