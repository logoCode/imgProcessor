package settings

import (
	"encoding/json"
	"errors"
	"image/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	FilenameIn, FilenameOut, Identifier, Separator string
	Accuracy                                       int
	Colors                                         map[int]color.RGBA
	Scaling                                        int
}

//change a color
func (settings *Settings) ChangeColor(kInt int, v string) error {
	col, err := getRGBA(v)
	if err != nil {
		return err
	} else {
		settings.Colors[kInt] = col
	}
	return nil
}

//set settings to default
func (settings *Settings) SetDefaultSettings() {
	settings.FilenameIn = "input.txt"
	settings.FilenameOut = "output.png"
	settings.Identifier = "$Data"
	settings.Separator = "/"
	settings.Accuracy = 0
	settings.Scaling = 1
	//build default color slice
	settings.Colors = make(map[int]color.RGBA)
	settings.Colors[0] = color.RGBA{255, 255, 255, 255}
	settings.Colors[1] = color.RGBA{255, 0, 0, 255}
	settings.Colors[2] = color.RGBA{0, 0, 255, 255}
	settings.Colors[3] = color.RGBA{0, 255, 0, 255}
	settings.Colors[4] = color.RGBA{0, 255, 255, 255}
	settings.Colors[5] = color.RGBA{255, 0, 255, 255}
	settings.Colors[6] = color.RGBA{255, 255, 0, 255}
	settings.Colors[7] = color.RGBA{0, 0, 0, 255}
}

//save settings to json file
func (settings *Settings) SaveSettings() error {
	//marshal settings to json format
	data, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	//save text in locales/settings.json and check for errors
	err = ioutil.WriteFile("locales/settings.json", data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//load settings from json file
func (settings *Settings) LoadSettings() error {
	data, err := ioutil.ReadFile("locales/settings.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, settings)
}

//translate color.RGBA to string
func GetString(c color.RGBA) string {
	return "rgba(" + strconv.Itoa(int(c.R)) + "," + strconv.Itoa(int(c.G)) + "," + strconv.Itoa(int(c.B)) + "," + strconv.Itoa(int(c.A)) + ")"
}

//translate string to color.RGBA
func getRGBA(colString string) (color.RGBA, error) {
	//split string which includes color information into its parameters
	rgbaStrings := strings.Split(colString, "/")
	//check if it are four
	if len(rgbaStrings) != 4 {
		return color.RGBA{255, 255, 255, 255}, errors.New("more or less than four arguments for rgba code")
	}
	//convert strings to ints and store them in slice rgbaInts
	rgbaInts := make([]uint8, len(rgbaStrings))
	for i := range rgbaStrings {
		integer, err := strconv.Atoi(rgbaStrings[i])
		rgbaInts[i] = uint8(integer)
		if err != nil {
			return color.RGBA{255, 255, 255, 255}, err
		}
	}
	//create and return rgba color
	return color.RGBA{rgbaInts[0], rgbaInts[1], rgbaInts[2], rgbaInts[3]}, nil
}

func GetMaximumKey(colors map[int]color.RGBA) (max int) {
	for i := range colors {
		if i > max {
			max = i
		}
	}
	return
}
