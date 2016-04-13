package main

import (
    "fmt"
    "strings"
    "math"
    "os"
)

type BitFlag int

func (flag *BitFlag) Rood() {
    *flag = rood
}

func (flag *BitFlag) Nederland() (string, error) {
    *flag = rood|wit|blauw
    return "Het is gelukt hoor!", nil
}


func (flag BitFlag) String() string {
    var flags []string
    if flag&rood == rood {
        flags = append(flags, "Rood")
    }
    if flag&wit == wit {
        flags = append(flags, "Wit")
    }
    if flag&blauw == blauw {
        flags = append(flags, "Blauw")
    }
    if flag&oranje == oranje {
        flags = append(flags, "Oranje")
    }
    if len(flags)>0 {
        return fmt.Sprintf("%d (%s)", int(flag), strings.Join(flags,"|"))
    }
    return "0()"
}

const (
    rood BitFlag =  1 << iota
    wit
    blauw
    oranje
)

func main() {
    flag := rood|oranje
    fmt.Printf("Kleurnummers: %d %d %d %d.\n", rood, wit, blauw, oranje)
    line()
    fmt.Println(flag)
    flag.Rood()
    line()
    fmt.Println(flag)
    msg, err := flag.Nederland()
    if err == nil {
        line()
        fmt.Println(msg)
        fmt.Println(flag)
    }
    line()
    fmt.Println(BitFlag(0))
    line()
    page_58()
    line()
}

func Uint8FromInt(x int) (uint8, error) {
    if 0<=x&&x<=math.MaxUint8{
        return uint8(x),nil
    }
    return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
func page_58(){
    const factor = 3 // factor is compatible with any numeric type
    i := 20000 // i is of type int by inference
    i *= factor
    j := int16(20) // j is of type int16; same as: var j int16 = 20
    i += int(j) // Types must match so conversion is required
    k := uint8(0) // Same as: var k uint8
    k = uint8(i) // Succeeds, but k's value is truncated to 8 bits ✗
    fmt.Println(i, j, k) // Prints: 60020 20 116   
    fmt.Println(math.MaxUint16)
    if u8, err := Uint8FromInt(int(26)); err != nil {
        fmt.Println(err)
        os.Exit(1)
    } else {
        fmt.Println(u8)   
    }
    line()
    const (
        ee int64 = 1000000000
        ef = 16.0/9.0
        eg = complex(-2, 3.5) * ef
        eh = 0.0 <= ef && ef < 2.0
    )
    fmt.Println(ee,ef,eg,eh)
}
func line(){
    fmt.Println("\n"+strings.Repeat("-=",20)+"-\n")
}

