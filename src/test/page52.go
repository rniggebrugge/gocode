package main

import "fmt"

func main() {
    var temp bool
    temp = true
    true:=false
    false:=temp
    fmt.Println(true)
    fmt.Println(false)
}