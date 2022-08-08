package sgcu

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "math"
)

func ShowArg () {
    var argc int = len (os.Args)
    fmt.Println ("argc is ", argc)
    for i:=0 ; i<argc ; i++ {
        fmt.Printf ("argv[%d] is %s\n", i, os.Args[i])
    }
}

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

func FloatToString (input_num float64, decimal_place int) string {
    if (input_num == 0) {
        return "0"
    }
    return strconv.FormatFloat (input_num, 'f', decimal_place, 64)
}

func FloatToIntStr (input_num float64) string {
    var str string
    var v float64 = math.Trunc (input_num)
    str = strconv.FormatFloat (v, 'f', 0, 64)
    return str
}

func FloatStringShrink (instr string) string {
    var dotaddr int = -1
    var outstr string = instr
    var outstrLen int = len (outstr)
    var outstrLen2 int = -1
    var endaddr int = outstrLen

    dotaddr = strings.Index (outstr, ".")

    if (dotaddr == -1) {
        return outstr
    }

    // ex: 10.
    if (dotaddr == (outstrLen - 1)) {
        return outstr[0:outstrLen-1]
    }

    for i:=outstrLen-1; i>dotaddr; i-- {
        //fmt.Println(i, string(outstr[i]))
        if (outstr[i] == '0') {
            endaddr = i
        } else {
            break
        }
    }

    // ex: 10.00 -> 10.
    outstrLen2 = len (outstr[0:endaddr])

    if (dotaddr == (outstrLen2 - 1)) {
        endaddr = dotaddr
    }

    return outstr[0:endaddr]
}

func FloatStringBeautySymbol (instr string) string {
    var retstr string
    var v float64

    if strings.Compare(instr, "0") == 0 {
        retstr = "+0"
        return retstr
    }

    v = StringToFloat (instr)
    //fmt.Println (v)

    if (v > 0) {
        retstr = "+" + instr
        return retstr
    }

    return instr
}
