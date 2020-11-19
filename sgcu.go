package sgcu

import (
    "fmt"
    "strings"
    "strconv"
    "os"
)

func IndexAt (content string, sep string, startaddr int) int {
    idx := strings.Index (content[startaddr:], sep)
    if idx > -1 {
        idx += startaddr
    }
    return idx
}

func StringToFloat (str string) float64 {
    floatvalue, err := strconv.ParseFloat (str, 64)
    if err != nil {
        fmt.Println (err)
        os.Exit (2)
    }
    return floatvalue
}
