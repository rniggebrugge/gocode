package main

import (
    "fmt"
    "math"
)

func main(){
    var b=byte(3)
    i:=0
    for {
       i++
       if i>15 {
            break
        }
       fmt.Printf("bits=%b, b=%d \n",b,b)
       b<<=1
    }
    dofloat()
}

func dofloat(){
    a:=math.Abs(-34)
    fmt.Println(a+30)
    x:=305.0
    y:=-346.0
    fmt.Println(math.Copysign(x,y))
    z:=34334.5235
    fmt.Printf("z=%f and in bits z=%b or %s \n",z,z, math.Float64bits(z))
}