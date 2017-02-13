package data

import(
    "os"
    "bufio"
    "errors"
    "strings"
    "strconv"
    "math"
)

//type Data is used to store all important imformation gained from an input file
type Data struct {
    Img [][]int //Img[x][y] = value
    X,Y int //x and y dimensions
}

//constructor for type Data
func NewData() Data {
    return Data{make([][]int,0),0,0}
}

//constructor for type Img
func NewImg(x,y int) [][]int {
    img := make([][]int,x)
    for i := range img {
        img[i] = make([]int,y)
    }
    return img
}

func (d *Data) CreateFromFile(filename, identifier, separator string, accuracy int) error {
    //stores all information
    var info [][]float64

    //open input file
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    //convert to readable scanner
    scanner := bufio.NewScanner(file)
    //scan over each line
    for scanner.Scan() {
        //remove whitespaces for different encodings
        text := removeWhiteSpace(scanner.Text())
        text = strings.TrimSpace(text)
        //check line for the given identifier, which labels it as important information
        if strings.HasPrefix(text, identifier) {
            //split line into the different parameters
            params := strings.Split(text, separator)
            //remove identifier
            params = params[1:]
            //store parameters as floats
            if len(params) < 3 {
                return errors.New("missing parameter")
            }
            values := make([]float64,len(params))
            var err error
            //range over parameters
            for i, param := range params {
                //convert parameters to floats and catch possible errors
                values[i], err = strconv.ParseFloat(param, 64)
                if err != nil {
                    return err
                }
            }
            //add values of the current line to all the other information
            info = append(info, values)
        }

    }
    //convert information to integers and scale it taking regards of the given accuracy
    infoAsInt := getInfoAsInt(info, accuracy)
    //get image dimensions and move the coordinates so that all of them are in the positive area
    x,y,infoAsInt := getImgDimensions(infoAsInt)
    //create new image
    d.Img = NewImg(x, y)
    d.X = x
    d.Y = y
    //assign the prepared information to the coordinates
    d.convertToImg(infoAsInt)

    return nil
}

//assign the prepared information to the coordinates stored in info[n][0] and info[n][1]
func (d *Data) convertToImg(info [][]int) {
    for i := range info {
        d.Img[info[i][0]][info[i][1]] = info[i][2]
    }
}

//convert information to integers and scale it taking regards of the given accuracy
func getInfoAsInt(info [][]float64, accuracy int) [][]int {
    var infoAsInt [][]int
    //range over info
    for i := range info {
        //add line to integer version of info
        infoAsInt = append(infoAsInt, make([]int, len(info[i])))
        //range over values
        for j := range info[i]{
            //if the value is no coordinate, just convert it to int
            if j > 1 {
                infoAsInt[i][j] = int(info[i][j])
            //else take regards of the given accuracy and then convert it to int
            }else {
                infoAsInt[i][j] = int(info[i][j] * math.Pow10(accuracy))
            }
        }
    }
    return infoAsInt
}

//get image dimensions and move the coordinates so that they are perfectly aligned to the images coordinate system
func getImgDimensions(info [][]int)(int, int, [][]int){
    //evaluate minimum x and y coordinates to align the according information to pixel(0|0) of the image
    x,y := 0,0
    for i := range info {
        if info[i][0] < x {
            x = info[i][0]
        }
        if info[i][1] < y {
            y = info[i][1]
        }
    }
    //add minimum x and y coordinates to all values
    for i := range info {
        info[i][0] = info[i][0] - x
        info[i][1] = info[i][1] - y
    }
    //evaluate correct dimensions after alignment
    for i := range info {
        if info[i][0] > x {
            x = info[i][0]
        }
        if info[i][1] > y {
            y = info[i][1]
        }
    }
    //add 1 to the dimensions for the 0 coordinates
    return x + 1, y + 1 ,info
}

//TrimSpace() doesn't work for every encoding
func removeWhiteSpace(text string) string {
    bytes := []byte(text)
    for i := len(bytes) - 1; i >= 0; i-- {
        if string(bytes[i]) == " " || string(bytes[i]) == " " {
            bytes = append(bytes[:i], bytes[i+1:]...)
        }
    }
    return string(bytes)
}
