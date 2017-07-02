package process

import (
	"extendedPlotter/data"
	"extendedPlotter/settings"
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
	rect := image.Rectangle{image.Point{0, 0}, image.Point{d.X, d.Y}}
	//build image form rectangle
	img := image.NewRGBA(rect)

	//range over x and y coordinates of data.Img
	for x := range d.Img {
		for y := range d.Img[x] {
			//set the color of the current pixel according to its value in d.Img (if it doesn't exist, it is set transparent)
			img.SetRGBA(x, d.Y-y, s.Colors[d.Img[x][y]])
		}
	}

	//finally save the image with the given name
	f, _ := os.Create(s.FilenameOut)
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}
