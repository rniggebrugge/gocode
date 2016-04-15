package main

import (
  "fmt"
  "sort"
  "strings"
)

type FoldedStrings []string

func (slice FoldedStrings) Len() int { return len(slice) }

func (slice FoldedStrings) Less(i, j int) bool {
  return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedStrings) Swap(i, j int) {
  slice[i], slice[j] = slice[j], slice[i]
}

type Rectangle struct {
  width float64
  height float64
}

type Areas []Rectangle

func (s Areas) Len() int { return len(s) }

func (s Areas) Less(i,j int) bool {
  return s[i].width*s[i].height < s[j].width*s[j].height
}
func (s Areas) Swap(i, j int){
  s[i], s[j] = s[j], s[i]
}

func (r Rectangle) String() string {
  return fmt.Sprintf("width %.2f x height %.2f = area %.2f\n", r.width, r.height, r.width*r.height)
}

func main(){
  files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
  SortFoldedStrings(files)
  fmt.Printf("Case insensitive: %q\n", files)
  
  shapes := []Rectangle{{34.3,55}, {1,30}, {0,10}, {32.5,99.5}, {0.2,200.4}}
  SortAreas(shapes)
  fmt.Printf("Small to big: %s]\n", shapes)
  

}

func SortFoldedStrings(slice []string){
  sort.Sort(FoldedStrings(slice))
}

func SortAreas(slice []Rectangle){
  sort.Sort(Areas(slice))
}
