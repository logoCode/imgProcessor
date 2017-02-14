package settings

import(
    "io/ioutil"
    "encoding/json"
    "errors"
    "image/color"
    "strconv"
    "strings"
    "os"
    "fmt"
    "project-x/scanner"
    "log"
)

type Settings struct {
    FilenameIn, FilenameOut, Identifier, Separator string
    Accuracy int
    Colors []color.RGBA
}




//change settings
func (settings *Settings) SettingsMenu() error {
    fmt.Println("SETTINGS MENU:")
    for {
        fmt.Println("Enter help to get a list of options or type in any other command.")
        input := scanner.GetS("==","help","list","filenameIn","filenameOut","filenameOut","separator","identifier","accuracy","color","return")
        switch input {
        case "help":
            help()
        case "list":
            settings.listSettings()
        case "filenameIn":
            fmt.Println("Enter the filename of the input file (ending with .txt):")
            settings.FilenameIn = scanner.GetString()
        case "filenameOut":
            fmt.Println("Enter the filename of the output file (ending with .png):")
            settings.FilenameOut = scanner.GetString()
        case "separator":
            fmt.Println("Enter the Separator, the values are separated with:")
            settings.Separator = scanner.GetString()
        case "identifier":
            fmt.Println("Enter the Identifier, the lines containing data start with:")
            settings.Identifier = scanner.GetString()
        case "accuracy":
            fmt.Println("Enter the number of decimal places the coordinates are cut off after:")
            settings.Accuracy = scanner.GetI("><",0,10)
        case "color":
            fmt.Println("Enter the value of the color you want to change:")
            k := scanner.GetI(">=",0)
            fmt.Println("enter delete to delete it, or type in the rgba value in the syntax r/g/b/a:")
            v := scanner.GetString()
            if v == "delete" {
                //if there are still colors afterwards, set it to white
                if k < len(settings.Colors) - 1 {
                    settings.Colors[k] = color.RGBA{255,255,255,255}
                //otherwise delete it
                }else if k == len(settings.Colors) - 1 {
                    settings.Colors = settings.Colors[:len(settings.Colors) - 1]
                }else {
                    log.Println("color not found")
                }
            }else {
                err := settings.changeColor(k,v)
                if err != nil {
                    log.Println(err)
                }
            }
        case "return":
            //save settings to json and return
            err := settings.SaveSettings()
            if err != nil {
                return err
            }
            return nil
        }
    }
}

//change a color
func (settings *Settings) changeColor(kInt int, v string) error {
    col, err := getRGBA(v)
    if err != nil {
        return err
    }else {
        //in case the slice isn't long enought yet its size is increased
        if len(settings.Colors) <= kInt {
            neededLen := kInt + 1 - len(settings.Colors)
            for i := 0; i < neededLen; i++ {
                settings.Colors = append(settings.Colors, color.RGBA{255,255,255,255})
            }
        }
        settings.Colors[kInt] = col
    }
    return nil
}

//list all commands
func help(){
    fmt.Println("List of options:")
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
func (settings *Settings) listSettings(){
    fmt.Println("List of settings:")
    fmt.Println("  - input filename: " + settings.FilenameIn)
    fmt.Println("  - output filename: " + settings.FilenameOut)
    fmt.Println("  - identifier: " + settings.Identifier)
    fmt.Println("  - separator: " + settings.Separator)
    fmt.Println("  - accuracy: " + strconv.Itoa(settings.Accuracy))
    fmt.Println("  - colors:")
    for i, color := range settings.Colors {
        fmt.Println("      - " + strconv.Itoa(i) + " : " + getString(color))
    }
}

//set settings to default
func (settings *Settings) SetDefaultSettings(){
    settings.FilenameIn = "input.txt"
    settings.FilenameOut = "output.png"
    settings.Identifier = "$Data"
    settings.Separator = "/"
    settings.Accuracy = 0
    //build default color slice
    settings.Colors = append(settings.Colors, color.RGBA{255,255,255,255})
    settings.Colors = append(settings.Colors, color.RGBA{255,0,0,255})
    settings.Colors = append(settings.Colors, color.RGBA{0,0,255,255})
    settings.Colors = append(settings.Colors, color.RGBA{0,255,0,255})
    settings.Colors = append(settings.Colors, color.RGBA{0,255,255,255})
    settings.Colors = append(settings.Colors, color.RGBA{255,0,255,255})
    settings.Colors = append(settings.Colors, color.RGBA{255,255,0,255})
    settings.Colors = append(settings.Colors, color.RGBA{0,0,0,255})
}

//save settings to json file
func (settings *Settings) SaveSettings() error {
    //build text string
    var text string
    text = text + `{
    `
    text = text + `"filenameIn": "` + settings.FilenameIn + `",
    `
    text = text + `"filenameOut": "` + settings.FilenameOut + `",
    `
    text = text + `"separator": "` + settings.Separator + `",
    `
    text = text + `"identifier": "` + settings.Identifier + `",
    `
    text = text + `"accuracy": "` + strconv.Itoa(settings.Accuracy) + `",
    `
    //add all colors to text
    for i := range settings.Colors {
        if i < len(settings.Colors) - 1 {
            text = text + `"` + strconv.Itoa(i) + `": "` + getJsonString(settings.Colors[i]) + `",
    `
        //take care of the missing comma in the last line
        }else {
            text = text + `"` + strconv.Itoa(i) + `": "` + getJsonString(settings.Colors[i]) + `"
`
        }
    }

    text = text + `}`

    //save text in locales/settings.json and check for errors
    err := ioutil.WriteFile("locales/settings.json",[]byte(text), os.ModePerm)
    if err != nil {
        return err
    }

    return nil
}

//load settings from json file
func (settings *Settings) LoadSettings() error {
    //firstly store the settings in a map of strings
    settingsAsMap := make(map[string]string)
    //get file content
    setFile, err := ioutil.ReadFile("locales/settings.json")
    if err != nil {
        return err
    }
    //convert file content to map
    err = json.Unmarshal(setFile, &settingsAsMap)
    if err != nil {
        return err
    }
    //read settings from map and assign values to settings struct
    for k, v := range settingsAsMap {
        kInt, err := strconv.Atoi(k)
        //for ints
        if k == "accuracy" {
            var setAc int
            setAc, err = strconv.Atoi(v)
            if err != nil {
                return err
            }else{
                settings.Accuracy = setAc
            }
        //for colors
        }else if err == nil {
            settings.changeColor(kInt, v)
        //for strings
        }else {
            switch k {
            case "filenameIn":
                settings.FilenameIn = v
            case "filenameOut":
                settings.FilenameOut = v
            case "separator":
                settings.Separator = v
            case "identifier":
                settings.Identifier = v
            default:
                return errors.New("json key " + k + " not found !")
            }
        }
    }
    return nil
}

//translate color.RGBA to string
func getString(c color.RGBA) string {
    return "rgba(" + strconv.Itoa(int(c.R)) + "," + strconv.Itoa(int(c.G)) + "," + strconv.Itoa(int(c.B)) + "," + strconv.Itoa(int(c.A)) + ")"
}

//translate color.RGBA to string for json file
func getJsonString(c color.RGBA) string {
    return strconv.Itoa(int(c.R)) + "/" + strconv.Itoa(int(c.G)) + "/" + strconv.Itoa(int(c.B)) + "/" + strconv.Itoa(int(c.A))
}

//translate string to color.RGBA
func getRGBA(colString string) (color.RGBA, error) {
    //split string which includes color information into its parameters
    rgbaStrings := strings.Split(colString, "/")
    //check if it are four
    if len(rgbaStrings) != 4 {
        return color.RGBA{255,255,255,255}, errors.New("more or less than four arguments for rgba code")
    }
    //convert strings to ints and store them in slice rgbaInts
    rgbaInts := make([]uint8,len(rgbaStrings))
    for i := range rgbaStrings {
        integer, err := strconv.Atoi(rgbaStrings[i])
        rgbaInts[i] = uint8(integer)
        if err != nil {
            return color.RGBA{255,255,255,255}, err
        }
    }
    //create and return rgba color
    return color.RGBA{rgbaInts[0],rgbaInts[1],rgbaInts[2],rgbaInts[3]}, nil
}
