package sgcu

import (
    "strings"
)

func indexAt (content string, sep string, startaddr int) int {
    idx := strings.Index(content[startaddr:], sep)
    if idx > -1 {
        idx += startaddr
    }
    return idx
}
