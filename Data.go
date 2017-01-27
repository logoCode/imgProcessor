package simplexOut

import(
    "project-x/random"
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math"
)

type Data struct {
    Img [][]int
    MaxVal int
}

func NewData() Data {
    return Data{make([][]int,0),0}
}

func NewImg(x,y int) [][]int {
    d := make([][]int,x)
    for i := range d {
        d[i] = make([]int,y)
    }
    return d
}

func (d *Data) CreateFromFile(filename, identifier, separator string, accuracy int){
    var info [][]float64
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text := removeWhiteSpace(scanner.Text())
        text = strings.TrimSpace(text)
        if strings.HasPrefix(text, identifier) {
            params := strings.Split(text, separator)
            params = params[1:]
            values := make([]float64,len(params))
            var err error
            for i := range params {
                values[i], err = strconv.ParseFloat(params[i], 64)
                if err != nil {
                    fmt.Println(err)
                    return
                }else {
                }
            }
            info = append(info, values)
        }

    }
    infoAsInt := getInfoAsInt(info, accuracy)
    x,y,infoAsInt := getImgDimensions(infoAsInt)
    d.Img = NewImg(x, y)
    d.convertToImg(infoAsInt)
    d.MaxVal = d.getMaxVal()
}

func (d *Data) convertToImg(info [][]int) {
    for i := range info {
        d.Img[info[i][0]][info[i][1]] = info[i][2]
    }
}

func getInfoAsInt(info [][]float64, accuracy int) [][]int {
    var infoAsInt [][]int
    for i := range info {
        infoAsInt = append(infoAsInt, make([]int, len(info[i])))
        for j := range info[i]{
            if j > 1 {
                infoAsInt[i][j] = int(info[i][j])
            }else {
                infoAsInt[i][j] = int(info[i][j]) * int(math.Pow10(accuracy))
            }
        }
    }
    return infoAsInt
}

func (d *Data) GetWidth() (width int) {
    for i := range d.Img {
        if len(d.Img[i]) > width {
            width = len(d.Img[i])
        }
    }
    return
}

func getImgDimensions(info [][]int)(int, int, [][]int){
    x,y := 0,0
    for i := range info {
        if info[i][0] < x {
            x = info[i][0]
        }
        if info[i][1] < y {
            y = info[i][1]
        }
    }
    for i := range info {
        info[i][0] = info[i][0] + int(math.Abs(float64(x)))
        info[i][1] = info[i][1] + int(math.Abs(float64(y)))
    }
    for i := range info {
        if info[i][0] > x {
            x = info[i][0]
        }
        if info[i][1] > y {
            y = info[i][1]
        }
    }
    return x + 1, y + 1 ,info
}

func (d *Data) CreateRandom(x, y, maxVal int){
    random.SetSeed()
    for i := 0; i < x;i++ {
        fmt.Println(i)
        d.Img = append(d.Img, make([]int, y))
        for j := range d.Img[i] {
            d.Img[i][j] = random.GetInt(0, maxVal)
        }
    }
    d.MaxVal = d.getMaxVal()
}

/*
func (d *Data) GetPercentageAsRGBA(x,y int) uint8 {
    percentage := float64(d.Img[x][y]) / float64(d.MaxVal)
    return uint8(percentage * 255)
}
*/

func (d *Data) getMaxVal() int {
    var maxVal int
    for i := range d.Img {
        for j := range d.Img[i] {
            if d.Img[i][j] > maxVal {
                maxVal = d.Img[i][j]
            }
        }
    }
    return maxVal
}

func removeWhiteSpace(text string) string {
    bytes := []byte(text)
    for i := len(bytes) - 1; i >= 0; i-- {
        if string(bytes[i]) == " " || string(bytes[i]) == " " {
            bytes = append(bytes[:i], bytes[i+1:]...)
        }
    }
    return string(bytes)
}
