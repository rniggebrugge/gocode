package main

import (
    "fmt"
   // "unicode/utf8"
    "strings"
    "unicode"
)

func main(){
    s:=`l presidente della re43535pubblica, Sergio Mattarella, si è detto particolarmente colpito dalla prematura scomparsa e ha ricordato l'"intellettuale, edito23re, protagonista politico innovativo e appassionato". Il presidente, in una lettera di cordoglio inviata alla famiglia, dice di averlo visto il 26 febbraio dello scorso anno al Q56767uirinale in occasione dell'incontro di una delegazione M5s, alla cui causa, ricorda il presidente, Casaleggio "aveva dedicato negli retwultimi anni tutto il suo impegno civile".`

    f:= func(char rune) rune {
        if(unicode.IsControl(char)) {
            return char
        }
        return ' '
    }
    
    fmt.Println(s+"\n"+strings.Repeat("-=",30))
    fmt.Println(strings.Map(f, s))
    fmt.Println(strings.Repeat("-=",30))
    fmt.Printf("%-20T: %#v\n",unicode.ASCII_Hex_Digit, unicode.ASCII_Hex_Digit)
}