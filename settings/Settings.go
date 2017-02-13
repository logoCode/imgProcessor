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
)

type Settings struct {
    FilenameIn, FilenameOut, Identifier, Separator string
    Accuracy int
    Colors []color.RGBA
}




//change settings
func (settings *Settings) ChangeSettings() error {
    fmt.Println("Enter the filename of the input file (ending with .txt):")
    settings.FilenameIn = scanner.GetString()
    fmt.Println("Enter the Identifier, the lines containing data start with:")
    settings.Identifier = scanner.GetString()
    fmt.Println("Enter the Separator, the values are separated with:")
    settings.Separator = scanner.GetString()
    fmt.Println("Enter the number of decimal places the coordinates are cut off after:")
    settings.Accuracy = scanner.GetI("><",0,10)
    fmt.Println("Enter the filename of the output file (ending with .png):")
    settings.FilenameOut = scanner.GetString()
    //save settings to json
    err := settings.SaveSettings()
    if err != nil {
        return err
    }
    return nil
}

//set settings to default
func (settings *Settings) SetDefaultSettings(){
    settings.FilenameIn = "input.txt"
    settings.FilenameOut = "output.png"
    settings.Identifier = "$Data"
    settings.Separator = "/"
    settings.Accuracy = 0
    //build color slice
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
            text = text + `"` + strconv.Itoa(i) + `": "` + getString(settings.Colors[i]) + `",
    `
        //take care of the missing comma in the last line
        }else {
            text = text + `"` + strconv.Itoa(i) + `": "` + getString(settings.Colors[i]) + `"
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
            var col color.RGBA
            col, err = getRGBA(v)
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
func getString(color.RGBA) string {
    return "255/255/255/255"
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
