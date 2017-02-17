package settings

import(
    "io/ioutil"
    "encoding/json"
    "errors"
    "image/color"
    "strconv"
    "strings"
    "os"
)

type Settings struct {
    FilenameIn, FilenameOut, Identifier, Separator string
    Accuracy int
    Colors map[int]color.RGBA
}

//change a color
func (settings *Settings) ChangeColor(kInt int, v string) error {
    col, err := getRGBA(v)
    if err != nil {
        return err
    }else {
        settings.Colors[kInt] = col
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
    //build default color slice
    settings.Colors = make(map[int]color.RGBA)
    settings.Colors[0] = color.RGBA{255,255,255,255}
    settings.Colors[1] = color.RGBA{255,0,0,255}
    settings.Colors[2] = color.RGBA{0,0,255,255}
    settings.Colors[3] = color.RGBA{0,255,0,255}
    settings.Colors[4] = color.RGBA{0,255,255,255}
    settings.Colors[5] = color.RGBA{255,0,255,255}
    settings.Colors[6] = color.RGBA{255,255,0,255}
    settings.Colors[7] = color.RGBA{0,0,0,255}
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
    maxKey := GetMaximumKey(settings.Colors)
    for i := 0; i <= maxKey; i++ {
        if i < maxKey {
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
            settings.ChangeColor(kInt, v)
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
func GetString(c color.RGBA) string {
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

func GetMaximumKey(colors map[int]color.RGBA) (max int) {
    for i := range colors {
        if i > max {
            max = i
        }
    }
    return
}
