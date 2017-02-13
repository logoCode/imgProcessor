package settings

import(
    "io/ioutil"
    "encoding/json"
    "errors"
    "image/color"
    "strconv"
    "strings"
)

type Settings struct {
    FilenameIn, FilenameOut, Identifier, Separator string
    Accuracy int
    Colors []color.RGBA
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
        if k == "Accuracy" {
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

//translate string to color.RGBA
func getRGBA(colString string) (color.RGBA, error) {
    rgbaStrings := strings.Split(colString, "/")
    if len(rgbaStrings) != 4 {
        return color.RGBA{255,255,255,255}, errors.New("more or less than four arguments for rgba code")
    }
    rgbaInts := make([]uint8,len(rgbaStrings))
    for i := range rgbaStrings {
        integer, err := strconv.Atoi(rgbaStrings[i])
        rgbaInts[i] = uint8(integer)
        if err != nil {
            return color.RGBA{255,255,255,255}, err
        }
    }
    return color.RGBA{rgbaInts[0],rgbaInts[1],rgbaInts[2],rgbaInts[3]}, nil
}
