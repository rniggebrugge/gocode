package main

import (
    "fmt"
    "math"
)

func main() {
    im := 4 - 1i
    in := 3 + 2i
    fmt.Println(im, in, im*in)
    x, y := 0.0, 0.0
    for i:=0; i<10; i++ {
        x+=0.1
        if i%2==0 {
            y+=0.2
        } else {
            fmt.Printf("%-5t %-5t %-5t %-5t", x==y,
                       EqualFloat(x,y,-1), EqualFloat(x,y,0.000001),
                       EqualFloatPrec(x,y,6))
            fmt.Println(x,y)
        }
    }
}

func EqualFloat(x,y, limit float64) bool {
    if limit <= 0.0 {
        limit = math.SmallestNonzeroFloat64
    }
    return math.Abs(x-y) <= (limit*math.Min(math.Abs(x), math.Abs(y)))
}

func EqualFloatPrec(x,y float64, decimals int) bool {
    a := fmt.Sprintf("%.*f", decimals, x)
    b := fmt.Sprintf("%.*f", decimals, y)
    fmt.Println(a,b)
    return len(a) == len(b) && a==b
}