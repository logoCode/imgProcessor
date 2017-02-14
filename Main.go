package main

import(
    "image"
    "fmt"
    "image/png"
    "os"
    "imgProcessor/data"
    "imgProcessor/settings"
    "log"
    "project-x/scanner"

)

var sett settings.Settings

//set default settings and start main menu
func main(){
    fmt.Println("-------------------------------------------------------------------------")
    fmt.Println("----------------------- Welcome to ImageProcessor -----------------------")
    fmt.Println("-----------------------   (C)2017 Max Obermeier   -----------------------")
    fmt.Println("-------------------------------------------------------------------------")
    sett.SetDefaultSettings()
    err := sett.LoadSettings()
    if err != nil {
        log.Println(err)
    }
    menu()
}

//main menu
func menu(){
    for {
        fmt.Println()
        fmt.Println("MAIN MENU:")
        fmt.Println("Enter help to get a list of options or type in any other command.")
        input := scanner.GetS("==","help","license","settings","process","exit")
        if input == "help" {
            help()
        }else if input == "license" {
            license()
        }else if input == "exit" {
            os.Exit(0)
        }else if input == "settings" {
            err := sett.SettingsMenu()
            if err != nil {
                log.Println(err)
            }
        }else if input == "process" {
            createImg(sett.FilenameIn, sett.FilenameOut, sett.Identifier, sett.Separator, sett.Accuracy)
        }
    }
}

//print license
func license(){
    fmt.Println("MIT License")
    fmt.Println("Copyright (c) 2017 Max Obermeier")
    fmt.Println("")
    fmt.Println(`Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:"`)
    fmt.Println("")
    fmt.Println("The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.")
    fmt.Println("")
    fmt.Println(`THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`)
    fmt.Println("")
}

//list all commands
func help(){
    fmt.Println("List of options:")
    fmt.Println("  - help \t \t=> Show list of options")
    fmt.Println("  - license \t \t=> Show license")
    fmt.Println("  - settings \t \t=> Show and change processing parameters")
    fmt.Println("  - process \t \t=> Start image creating process")
    fmt.Println("  - exit \t \t=> Exit program")
}

//create image
func createImg(FilenameIn, FilenameOut, Identifier, Separator string, Accuracy int){
    //create data from file
    d := data.NewData()
    err := d.CreateFromFile(sett.FilenameIn, sett.Identifier, sett.Separator, sett.Accuracy)
    //check for any errors
    if err != nil {
        log.Println(err)
        return
    }
    //create rectangle with fitting dimensions (rectangle is used to build a rgbaImage)
    rect := image.Rectangle{image.Point{0, 0}, image.Point{d.X, d.Y}}
    //build image form rectangle
    img := image.NewRGBA(rect)

    //range over x and y coordinates of data.Img
    for x := range d.Img {
        for y := range d.Img[x]{
            //if there is a specific color for this value in the colors slice use it
            if d.Img[x][y] < len(sett.Colors) {
                //invert y - coordinates, because (0|0) of a image/png is at the top left corner and not at the bottom left as in a coordinate system
                img.SetRGBA(x, d.Y - y, sett.Colors[d.Img[x][y]])
            //if there is no color specified just use white
            }else {
                //invert y - coordinates, because (0|0) of a image/png is at the top left corner and not at the bottom left as in a coordinate system
                img.SetRGBA(x, d.Y - y, sett.Colors[0])
            }
        }
    }

    //finally save the image with the given name
    f, _ := os.Create(sett.FilenameOut)
    defer f.Close()
    err = png.Encode(f, img)
    if err != nil {
        log.Println(err)
        return
    }

    fmt.Println("Image was created successfully !")

}
