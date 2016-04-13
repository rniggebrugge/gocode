package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
//    "math"
    "math/cmplx"    
//    "strings"
)

const (
    pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Quadratic Equation Solver</title>
<body><h3>Quadratic Equation Solver</h3>
<p>Solves equations of the form ax&sup2; +bx + c = 0</p>`
    form       = `<form action="/" method="POST">
<input type="text" name="a" size="3">x&sup2; +
<input type="text" name="b" size="3">x +
<input type="text" name="c" size="3">  
<input type="submit" value="Calculate">
</form>`
    pageBottom = `</body></html>`
    anError    = `<p class="error">%s</p>`
)

type equations struct {
    a float64
    b float64
    c float64
    answer1 complex128
    answer2 complex128
}

func main() {
    http.HandleFunc("/", homePage)
    if err := http.ListenAndServe(":9001",nil); err != nil {
        log.Fatal("failed to start server", err)
    } 
}

func homePage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    fmt.Fprint(writer, pageTop, form)
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        if numbers, message, ok := processRequest(request); ok {
            equation := getEquation(numbers)
            fmt.Fprint(writer, formatSolution(equation))
        } else if message != "" {
            fmt.Fprintf(writer, anError, message)
        }
    }
    fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
    var numbers []float64
    for _, fieldname := range []string{"a", "b", "c"} {
        if slice, found := request.Form[fieldname]; found && len(slice) > 0 {
            field:=slice[0];
            if a,err := strconv.ParseFloat(field, 64); err != nil {
                if (fieldname=="c") {
                    numbers = append(numbers, 0.0)
                } else {
                    numbers = append(numbers, 1.0)
                }
            } else {
                numbers = append(numbers, a)
            }
        } else {
            numbers = append(numbers,0.0)
        }
    }    
    return numbers, "", true
}

func formatSolution(equation equations) string {
    if equation.a==0 && equation.b==0 && equation.c == 0 {
        return ""
    }
    return fmt.Sprintf(`<p>%fx&sup2; + %fx + %f = 0 => x = %f or x = %f`, equation.a, equation.b, equation.c, equation.answer1, equation.answer2)
}

func getEquation(numbers []float64) (equation equations) {
    equation.a = numbers[0]
    equation.b = numbers[1]
    equation.c = numbers[2]    
    equation.answer1, equation.answer2 = solve(numbers)
    return equation
}

func solve (numbers []float64) (answer1, answer2 complex128){
    a:=numbers[0]
    b:=numbers[1]
    c:=numbers[2]
    fmt.Printf("solving %fx2; + %fx + %f = 0 \n", a,b,c)

    if a==0 && b==0 && c==0 {
        return 0,0
    }
    if a==0 {
        return complex(-c/b,0), complex(-c/b,0)
    }    

    ac:=complex(a,0)
    bc:=complex(b,0)
    cc:=complex(c,0)
    
    root := cmplx.Sqrt(cmplx.Pow(bc, 2) - (4 * ac * cc))
    x1 := (-bc + root) / (2 * ac)
    x2 := (-bc - root) / (2 * ac)
    return x1, x2
}
