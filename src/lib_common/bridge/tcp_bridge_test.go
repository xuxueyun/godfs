package bridge

import (
    "testing"
    "fmt"
)



func Test1(t *testing.T) {
    fmt.Print("123")
    fmt.Print(string([]byte{10,13}))
    fmt.Print("456")
    fmt.Print(rune('\n'))
    fmt.Print(rune('\r'))
    fmt.Print(rune('\r'))


}
