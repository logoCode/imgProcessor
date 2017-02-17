package main

import(
    "fmt"
    "strconv"
    "os"
    "imgProcessor/process"
    "imgProcessor/settings"
    "log"
    "project-x/scanner"

)

var sett settings.Settings

//set default settings and start main menu
func main(){
    fmt.Println("-------------------------------------------------------------------------")
    fmt.Println("-------------------------------------------------------------------------")
    fmt.Println("----------------------- Welcome to ImageProcessor -----------------------")
    fmt.Println("-----------------------   (C)2017 Max Obermeier   -----------------------")
    fmt.Println("-------------------------------------------------------------------------")
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
        fmt.Println("----- MAIN MENU ---------------------------------------------------------")
        fmt.Println()
        fmt.Println("Enter help to get a list of options or type in any other command.")
        input := scanner.GetS("==","help","license","settings","process","exit")
        if input == "help" {
            help()
        }else if input == "license" {
            license()
        }else if input == "exit" {
            os.Exit(0)
        }else if input == "settings" {
            err := settingsMenu()
            if err != nil {
                log.Println(err)
            }
        }else if input == "process" {
            err := process.CreateImg(sett)
            if err != nil {
                log.Println(err)
            }else {
                fmt.Println()
                fmt.Println("Image was created successfully !")
            }
        }
    }
}

//print license
func license(){
    fmt.Println()
    fmt.Println("----- MIT LICENSE -------------------------------------------------------")
    fmt.Println()
    fmt.Println("Copyright (c) 2017 Max Obermeier")
    fmt.Println()
    fmt.Println(`Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:"`)
    fmt.Println()
    fmt.Println("The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.")
    fmt.Println()
    fmt.Println(`THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`)
}

//list all commands
func help(){
    fmt.Println()
    fmt.Println("----- LIST OF OPTIONS ---------------------------------------------------")
    fmt.Println()
    fmt.Println("  - help \t \t=> Show list of options")
    fmt.Println("  - license \t \t=> Show license")
    fmt.Println("  - settings \t \t=> Show and change processing parameters")
    fmt.Println("  - process \t \t=> Start image creating process")
    fmt.Println("  - exit \t \t=> Exit program")
}

//list all commands
func helpSettings(){
    fmt.Println()
    fmt.Println("----- LIST OF OPTIONS ---------------------------------------------------")
    fmt.Println()
    fmt.Println("  - help \t \t=> Show list of options")
    fmt.Println("  - list \t \t=> List all settings and according values")
    fmt.Println("  - filenameIn \t \t=> Change name of input file")
    fmt.Println("  - filenameOut \t=> Change name of output file")
    fmt.Println("  - separator \t \t=> Change separator")
    fmt.Println("  - identifier \t \t=> Change identifier")
    fmt.Println("  - accuracy \t \t=> Change accuracy")
    fmt.Println("  - color \t \t=> Change a color")
    fmt.Println("  - return \t \t=> Return to main menu and save settings")
}

//list settings
func listSettings(){
    fmt.Println()
    fmt.Println("----- LIST OF SETTINGS --------------------------------------------------")
    fmt.Println()
    fmt.Println("  - input filename: " + sett.FilenameIn)
    fmt.Println("  - output filename: " + sett.FilenameOut)
    fmt.Println("  - identifier: " + sett.Identifier)
    fmt.Println("  - separator: " + sett.Separator)
    fmt.Println("  - accuracy: " + strconv.Itoa(sett.Accuracy))
    fmt.Println("  - colors:")
    for i := 0; i <= settings.GetMaximumKey(sett.Colors); i++ {
        fmt.Println("      - " + strconv.Itoa(i) + " : " + settings.GetString(sett.Colors[i]))
    }
}

//change settings
func settingsMenu() error {
    for {
        fmt.Println()
        fmt.Println("----- SETTINGS MENU -----------------------------------------------------")
        fmt.Println()
        fmt.Println("Enter help to get a list of options or type in any other command.")
        input := scanner.GetS("==","help","list","filenameIn","filenameOut","filenameOut","separator","identifier","accuracy","color","return")
        switch input {
        case "help":
            helpSettings()
        case "list":
            listSettings()
        case "filenameIn":
            fmt.Println("Enter the filename of the input file (ending with .txt):")
            sett.FilenameIn = scanner.GetString()
        case "filenameOut":
            fmt.Println("Enter the filename of the output file (ending with .png):")
            sett.FilenameOut = scanner.GetString()
        case "separator":
            fmt.Println("Enter the Separator, the values are separated with:")
            sett.Separator = scanner.GetString()
        case "identifier":
            fmt.Println("Enter the Identifier, the lines containing data start with:")
            sett.Identifier = scanner.GetString()
        case "accuracy":
            fmt.Println("Enter the number of decimal places the coordinates are cut off after:")
            sett.Accuracy = scanner.GetI("><",0,10)
        case "color":
            fmt.Println("Enter the value of the color you want to change:")
            k := scanner.GetI(">=",0)
            fmt.Println("enter delete to delete it, or type in the rgba value in the syntax r/g/b/a:")
            v := scanner.GetString()
            if v == "delete" {
                _, ok := sett.Colors[k]
                if ok {
                    delete(sett.Colors, k)
                }else {
                    log.Println("color not found")
                }
            }else {
                err := sett.ChangeColor(k,v)
                if err != nil {
                    log.Println(err)
                }
            }
        case "return":
            //save settings to json and return
            err := sett.SaveSettings()
            if err != nil {
                return err
            }
            return nil
        }
    }
}
