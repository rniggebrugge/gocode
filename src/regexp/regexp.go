package main

import(
    "fmt"
    "regexp"
    "strings"
)

func main(){
    var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
    
    fmt.Println(validID.MatchString("adam[23]"))
    fmt.Println(validID.MatchString("eve[7]"))
    fmt.Println(validID.MatchString("remco.nggebrugge[35]"))    
    fmt.Println(validID.MatchString("Job[48]"))    
    fmt.Println(validID.MatchString("snakey"))
    
    names := []string{"Remco Niggebrugge", "Clara Moura Cruz", "Willem Gustav Alberto Niggebrugge", "Pieter Jan Thomas Niggebrugge","F.D.E. Niggebrugge, H.W. Niggebrugge-Rincker","Clara Cruz & Remco Niggebrugge"}
    
    
    nameRx := regexp.MustCompile(`(?P<forename>\pL+\.?(?:\s*\pL+\.?)*)\s+(?P<surname>[\pL\-]+)`)
    
    for i:=0; i<len(names);i++ {
        fmt.Println(names[i])
        names[i] = nameRx.ReplaceAllString(names[i], "${surname}, ${forename}")
        fmt.Println(names[i]+"\n")
    }
    
    pattern := `^[a-z].*\d$`
    line := "remco340"
    fmt.Println(regexp.MatchString(pattern, line))
    
    wordRx := regexp.MustCompile(`(\w+)`)
    for i:=0; i<len(names);i++ {
        fmt.Println(names[i])
        fmt.Println(wordRx.ReplaceAllString(names[i], "woord"))
        fmt.Println("----- ------ ------ ------")
    }
    
    text := "this is is a good text text with some some duplicates in it."
    if matches := wordRx.FindAllString(text, -1); matches != nil {
        fmt.Printf("%#v\n",matches)
        previous := ""
        for _, match := range matches {
            if match == previous {
                fmt.Println("Duplicate word: ", match)
            }
            previous = match
        }
    }
    
    lines := `
# global variables
pageTitle : "Main Menu"
bodyBgColor : #000000
tableBgColor : #000000
rowBgColor : #00ff00

[Customer]
pageTitle : "Customer Info"

[Login]
pageTitle : "Login"
focus :         "username"
Intro : "This is a value that spans more
           than one line. you must enclose
           it in triple quotes."

# hidden section
[.Database]
host:my.example.com
db:ADDRESSBOOK
user:php-user
pass:foobar
myname:     remco niggebrugge
`
    fmt.Println(strings.Repeat("-=0",25))
    valueForKey := make(map[string]string)
    kvRx := regexp.MustCompile(`\s*([[:alpha:]]\w*)\s*:\s*(.+)`)
    if matches := kvRx.FindAllStringSubmatch(lines, -1); matches != nil {
        for _, match:= range matches {
            valueForKey[match[1]] = strings.TrimRight(match[2], "\t")
        }
    }

    for k,v := range valueForKey {
        fmt.Printf("%20s = %-20s\n",k,v)
    }
    fmt.Println("\n\n\n"+strings.Repeat("-=0",25))
    for _, m := range kvRx.FindAllStringSubmatch(lines,-1){
        fmt.Println("---->", m)
        
    }
    fmt.Println(strings.Repeat("-=0",25))
    for _, m := range kvRx.FindAllString(lines,-1){
       fmt.Println("--->>",m)
        
    }
}