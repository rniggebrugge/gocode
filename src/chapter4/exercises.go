package main

import (
  "fmt"
  "strings"
  "sort"
)

func main(){
  s:= []int{1,4,3,7,5,6,1,8,0,6,7,5,9,4,7,8,12,5,3,2,4,3,14}
  t:=UniqueInts(s)
  fmt.Printf("%v => %v\n\n", s, t)
  t=uniqueI(s)
  fmt.Printf("%v => %v\n\n", s, t)
  
  irm := [][]int{{1,4},{3,7,66,5,2},{4,7,8,5}}
  rem := Flatten(irm)
  fmt.Printf("%v => %v\n\n", irm, rem)
  
  s = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
  for i:=3; i<10;i++ {
    m2d := Make2D(s, i)
    fmt.Printf("%3d %v\n",i,m2d)
  }
  fmt.Println()
  
  iniData := []string{
    "; Cut down copy of Mozilla application.ini file",
    "",
    "This is an example file (invalid because of this line!)",
    "[App]",
    "Vendor=Mozilla",
    "Name=Iceweasel",
    "Profile=mozilla/firefox",
    "Version=3.5.16",
    "[Gecko]",
    "MinVersion=1.9.1",
    "MaxVersion=1.9.1.*",
    "[XRE]",
    "EnableProfileMigrator=0",
    "EnableExtensionManager=1",
  }
  parsed := ParseIni(iniData)
  prettyPrint(parsed)
  
  fmt.Println("--------\n\n---------")
  PrintIni(parsed)
}

func uniqueI(s []int) (result []int) {
  seen:= map[int]bool{}
  for _,x := range s {
    if _, found := seen[x]; !found {
      result = append(result, x)
      seen[x]=true
    }
  }
  return result
}

func UniqueInts(s []int) (result []int) {
  result = append(result, s[0])
  for i:=1;i<len(s);i++ {
    unique:=true
    for j:=0; j<i; j++ {
      if s[i]==s[j] {
        unique=false
        break
      }
    }
    if unique {
      result = append(result, s[i])
    }
  }
  return result
}

func Flatten(s [][]int) []int{
  result := make([]int, 0, len(s)*len(s[0]))
  for _, row := range s {
    for _, item := range row {
      result = append(result, item)
    }
  }
  return result
}

func Make2D(s []int, columns int) [][]int {
  rows := (len(s)+columns-1)/columns
  result := make([][]int, rows)
  item_number:=0
  for i:=0; i<rows; i++ {
    row:= make([]int, columns)
    for j:=0; j<columns; j++ {
      row[j]=s[item_number]
      item_number++
      if item_number>=len(s) {
        break
      }
    }
    result[i] = row
  }
  return result
}

func ParseIni(iniData []string) map[string]map[string]string {
  result := make(map[string]map[string]string, len(iniData))
  var current_group_name string
  current_group := make(map[string]string)
  for _, line := range iniData {
    line = strings.TrimSpace(line)
    if len(line) == 0 || line[0]== ';' {
      continue
    }
    if strings.HasPrefix(line, "["){
      if j:= strings.Index(line, "]"); j > -1 {
        if current_group_name != "" {
          result[current_group_name] = current_group
        }
        current_group_name = line[1:j]
        current_group = map[string]string{}
      }
    } else if current_group_name != ""{
      if j:= strings.Index(line,"="); j>0 {
        current_group[line[:j]]=line[j+1:]
      }
    }
  }
  result[current_group_name] = current_group
  return result
}

func prettyPrint(ini map[string]map[string]string){
  fmt.Println("map [ ")
  for key, group := range ini {
    fmt.Println("     "+key+" : map [ ")
    for key, value := range group {
      fmt.Printf("       %-10s: %s\n", key, value)
    }
    fmt.Println("     ]")
  }
  fmt.Println("]")
}

func PrintIni(ini map[string]map[string]string){
  groups := make([]string,0)
  for g := range ini {
    groups = append(groups, g)
  }
  sort.Strings(groups)
  for _,g := range groups {
    fmt.Printf("[%s]\n", g)
    names := make([]string,0)
    for n := range ini[g] {
      names = append(names, n)
    }
    sort.Strings(names)
    for _,n := range names {
      fmt.Printf("%s=%s\n",n, ini[g][n])
    }
  }
}
