package main

import (
    "fmt"
    "net/http"
    "log"
    "strings"
    "html"
    "io/ioutil"
    "sort"
)

const (
    pageTop = `<!DOCTYPE HTML><html><head><style>.error{color:red}</style><title>Soundex</title>
               </head><body><h1>Soundex</h1><p>Compute soundex codes for a list of names</p>
<p>Names (comma or space separated)</p>
<form method="post" action="/"><input type="text" name="names" size="60"> <input type="submit" value="Calculate">
</form>`
    pageBottom = `<hr><a href="/">Calculate</a> | <a href="/test">Testcases</a></body></html>`
    anError = `<p class="error">%s</p>`
)

var digitForLetter = []rune{
    0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5,
 // A  B  C  D  E  F  G  H  I  J  K  L  M
    5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2   }
 // N  O  P  Q  R  S  T  U  V  W  X  Y  Z

var testCases map[string]string

func main() {
    http.HandleFunc("/", homePage)
    var ok bool
    if testCases, ok = readTestCases("soundex-test-data.txt"); ok {
        http.HandleFunc("/test", testPage)
    }
    if err := http.ListenAndServe(":9001", nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}

func readTestCases(filename string) (map[string]string, bool){
    tests := make(map[string]string)
    if lines, err := ioutil.ReadFile(filename); err != nil {
        log.Println(err)
        return tests, false
    } else {
        for _, line := range strings.Split(string(lines), "\n") {
            if fields := strings.Fields(line); len(fields) == 2 {
                tests[fields[1]] = fields[0]
            }
        }
    }
    return tests, true    
}

func homePage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    fmt.Fprint(writer, pageTop)
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        if names := processRequest(request); len(names)>0 {
            soundexes := make([]string, len(names))
            for i, name := range names {
                soundexes[i] = soundex(name)
            }
            fmt.Fprint(writer, formatResults(names,soundexes))
        }
    }
    
    fmt.Fprint(writer, pageBottom)
}

func testPage(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, `<html><head><title>Soundex Test</title><style>.pass { color:green} .fail { color:red}</style>`)
    fmt.Fprint(writer, `</head><body><table border="1"><thead><th>Name</th><th>Soundex</th>`)
    fmt.Fprint(writer, `<th>Expected</th><th>Test</th></thead>`)
    names := []string{}
    for name, _ := range testCases {
        names = append(names, name)
    }
    sort.Strings(names)
    for _, name := range names {
        actual := soundex(name)
        expected := testCases[name]
        test := `<span class="fail">FAIL</span>`
        if actual == expected {
            test = `<span class="pass">PASS</span>`
        }
        fmt.Fprintf(writer, `<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>`, name, actual, expected, test)
    }
    fmt.Fprint(writer, `</table>`, pageBottom)
}

func soundex(name string) string {
    name = strings.ToUpper(name)
    chars := []rune(name)
    codes := []rune{}
    codes = append(codes, chars[0])
    for i:=1; i<len(chars); i++ {
        char := chars[i]
        if char == chars[i-1] {
            continue
        }
        if index := char - 'A'; index >= 0 && index < int32(len(digitForLetter)) && digitForLetter[index] != 0 {
            codes = append(codes, '0' + digitForLetter[index])
        }
    }
    for len(codes)<4 {
        codes = append(codes, '0')
    }
    return string(codes[:4])
}

func formatResults(names, soundexes []string) string {
    txt := `<table border="1"><thead><th>Name</th><th>Soundex</th></thead>`
    for i, name := range names {
        txt += "<tr><td>"+html.EscapeString(name)+"</td><td>"+html.EscapeString(soundexes[i])+"</td></tr>"
    }
    txt += "</table>"
    return txt
}

func processRequest(request *http.Request)(names []string){
    if slice, found := request.Form["names"]; found && len(slice)>0 {
        text := strings.Replace(slice[0],","," ",-1)
        names = strings.Fields(text)
    }
    return names
}
    