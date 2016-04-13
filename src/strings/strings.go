package main

import (
    "fmt"
    "strings"
    "unicode"
    "unicode/utf8"
    "math"
    "io"
)

func main(){
    /*
    printstring("vått og tørt")
    printstring("naïve, or just very smart?")
    printstring("røde og gule sløjfer")
    printstring("rå tørt\u2028vær")
    page96()
    */
    //page100()
    //page101()
    //page102()
    //page104()
    page107()
}

type polar struct {
    radius float64
    angle float64
}

func page107(){
    family := "remco*clara willem|pieter"
    names := strings.FieldsFunc(family, func(char rune) bool {
        switch char {
            case '*', ' ','|': return true
        }
        return false
    })
    fmt.Printf("%12s\n",names)
    fmt.Println(strings.Map(func(char rune) rune {
        if char%2==0 {
            return '?'
        }
        return char
    },"Niccolò•Noël•Geoffrey•Amélie••Turlough•José"))
    fmt.Println("\n--------\n")
    
    reader := strings.NewReader("Niccolò•Noël•Geoffrey•Amélie••Turlough•José drink Café")
    fmt.Printf("%T %v\n",reader ,reader)
    
    count:=0
    for {
        char, _, err := reader.ReadRune()
        if err !=nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }
        fmt.Printf("%c",char)
        count++
        if count==8 {
            reader.UnreadRune()
            count=0
        }
    }
    fmt.Println()
                   
}
func page104(){
    p := polar{10.10, 20.20}
    b := false
    i := 7607
    j := math.E
    c := 5+7i
    s := "Relativity"
    fmt.Printf("|%T|%p|%v|%#v|\n", p, &p, p, p)
    fmt.Printf("|%T|%p|%v|%t|\n", b, &b, b, b)
    fmt.Printf("|%T|%p|%v|%d|\n", i, &i, i, i)
    fmt.Printf("|%T|%p|%v|%f|\n", j, &j, j, j)
    fmt.Printf("|%T|%p|%v|%f|\n", c, &c, c, c)
    fmt.Printf("|%T|%p|%v|%s|%q|\n", s, &s, s, s, s)
    fmt.Println("\n\n\n")
    s = "Alias↔Synonym"
    chars := []rune(s)
    bytes := []byte(s)
    fmt.Printf("%T: %v\n%T: %v\n", chars, chars, bytes, bytes)
    fmt.Println("\n\n\nAddresses:\n")
    collection:=[]interface{}{p,b,i,j,c,s, chars, bytes}
    for _, item:= range collection {
        fmt.Printf("%20T: %p -> %v\n", item, &item, item)
    }
}

func printstring(phrase string) {
    fmt.Println("/===============================\\")
    fmt.Printf("string: \"%s\"\n", phrase)
    fmt.Println("|===============================|")
    fmt.Println("index  rune    char bytes")
    for index, char := range phrase {
         fmt.Printf("%-2d     %U  %c    % X\n",
         index, char, char,
         []byte(string(char)))
    }
    fmt.Println("|===============================|")
    out := ""
    for i:=0; i<len(phrase);i++ {
        out += string(phrase[i])
    }
    fmt.Println(out)
    fmt.Println("|===============================|")
    firstWord, lastWord := startend(phrase)
    fmt.Printf("First word: %v, last word: %v\n", firstWord, lastWord) 
    fmt.Println("\\===============================/\n")
}

func startend(line string) (firstWord, lastWord string) {
    i := strings.IndexFunc(line, unicode.IsSpace)
    firstWord = line[:i]
    j := strings.LastIndexFunc(line, unicode.IsSpace)
    _, size := utf8.DecodeRuneInString(line[j:])
    lastWord = line[j+size:]
    return firstWord, lastWord
}

func page96(){
    fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37,37,37,37,37)
    fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n",41,41,41,41,-41)
    i:=3931
    fmt.Printf("|%x|%X|%8x|%8X|%08x|%08X|%#04X|0x%04X|\n",i,i,i,i,i,i,i,i)
    i = 569
    fmt.Printf("|%d|%+d|%06d|%+06d|%s|\n",i, i,i,i,Pad(i,6,'*'))
    fmt.Printf("%d %#04x %08o %U %c\n", 0x3A6, 934, 934, '\u03A6','\U000003A6')
}
     
func Pad(number, width int, pad rune) string {
    s := fmt.Sprint(number)
    gap := width - utf8.RuneCountInString(s)
    if gap > 0 {
        return strings.Repeat(string(pad), gap) + s
    }
    return s
}

func page100(){
    for i, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
        fmt.Printf("%3d. %20.5e|%20.5f|%s|\n", i+1, x, x, Humanize(x, 20, 2, '*', ','))
    }
}
               
func Humanize(amount float64, width, decimals int, pad, separator rune) string {
    dollars, cents := math.Modf(amount)
    whole := fmt.Sprintf("%+.0f", dollars)[1:]
    fraction := ""
    if decimals >0 {
        fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:]
    }
    sep := string(separator)
    for i := len(whole)-3;i>0; i-=3 {
        whole = whole[:i] + sep + whole[i:]
    }
    if amount < 0.0 {
        whole = "-" + whole
    }
    number := whole + fraction
    gap := width - utf8.RuneCountInString(number)
    if gap > 0 {
        return strings.Repeat(string(pad), gap) + number
    }
    return number
}

func page101(){
    for _, x := range []complex128{2+3i, 172.6-58.3019i, -.8277e2+9.04831e-3i} {
        fmt.Printf("%15s|%9.3f|%.2f|%.1e|\n",
                   fmt.Sprintf("%6.2f%+.3fi", real(x), imag(x)), x,x,x)
    }
}

func page102(){
    slogan := "End Órétt5læti♥"
    bytes := []byte(slogan)
    fmt.Printf(" |%40s| \n %x \n %X \n % X \n %v \n", bytes, bytes, bytes, bytes, bytes)
    
    i := strings.Index(slogan, "t")
    fmt.Println(i)
    fmt.Printf("\n\n\n|%.10s|%.*s|%22.10s|%s|\n", slogan, i, slogan, slogan, slogan)
}