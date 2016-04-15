package main

import (
  "fmt"
  // "log"
)
  
type composer struct {
    name string
    year int
}

type rectangle struct {
  x0, y0, x1, y1 int
  fill string
}
  
func (c *composer)SetName(n string){
    c.name = n
}
  
func main(){
  x:=10
  y:=20
  product:=0
  fmt.Printf("%d * %d = %d  ????? \n", x, y, product)
  swapAndMultiply(&x, &y, &product)
  fmt.Printf("%d * %d = %d \n", x, y, product)
  x,y, product = swap2(x,y)
  fmt.Printf("%d * %d = %d \n", x, y, product)
  
  p := myname()
  fmt.Println(*p)
  
  antonio := composer{"antonio", 1707}
  agnes := new(composer)
  agnes.name, agnes.year = "agnes", 1845
  julia := &composer{}
  julia.name, julia.year = "julia", 1819
  augusta := &composer{"augusta", 1847}
  fmt.Printf("%v (%T) - %v (%T) - %v (%T) - %v (%T) \n", antonio, antonio, agnes, agnes, julia, julia, augusta, augusta)
  antonio.SetName("remco")
  agnes.SetName("clara")
  setname(augusta, "willem en pieter")
  fmt.Printf("%v (%T) - %v (%T) - %v (%T) - %v (%T) \n", antonio, antonio, agnes, agnes, julia, julia, augusta, augusta)

  grades := []int{32,35,1,87,3,12,53,4}
  inflate_grades(grades, 10)
  fmt.Println(grades)
  
  rect := rectangle{4, 8, 20, 10, "##ff00aa"}
  fmt.Println(rect)
  resizeRect(&rect, 5,5 )
  fmt.Println(rect)
  
  arr1 := [...]string{"Willem","Pieter","Clara","Remco"}
  arr2 := []string{"Willem","Pieter","Clara","Remco"}
  slice1 :=make([]string,4,10)
  fmt.Printf("arr1: %T : %v \n", arr1, arr1)
  fmt.Printf("arr2 %T : %v \n", arr2, arr2)
  fmt.Printf("slice1 %T : %v \n", slice1, slice1)
}

func resizeRect(rect *rectangle, dw, dh int){
  (*rect).x1 += dw
  rect.y1 += dh
}

func inflate_grades(grades []int, n int){
  for i:= range grades {
      grades[i]+=n
  }
}

func setname(c *composer, name string){
  c.name = name
}

func swapAndMultiply(x,y,p *int){
  *x, *y = *y, *x
  *p = *x**y
}

func swap2(x,y int) (int, int, int){
  return y,x,x*y
}

func myname() *string {
  name := "Remco Niggebrugge"
  return &name
}
