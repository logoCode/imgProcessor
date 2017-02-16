package main

import(
    "fmt"
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
            err := sett.SettingsMenu()
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
