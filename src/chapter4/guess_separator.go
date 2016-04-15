package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func main(){
  if len(os.Args)==1 || os.Args[1] == "-h" || os.Args[1] == "--help"{
    fmt.Printf("usage: %s file\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }
  
  candidates := []string{"\t" ,  "*",  "," ,  "|" , " " , "/"}
  linesRead, lines := readUpToNLines(os.Args[1],300)
  if linesRead>0 {
    counts := createCount(lines, linesRead, candidates)
    nEqual := guessSeparator(counts, candidates)
    bestGuess := 0
    maxScore := 0
    for i, c := range nEqual {
      fmt.Printf("Separator %4q has same non-zero frequency in %3d lines (%7.2f%%)\n", candidates[i], c, float64(100*c)/float64(linesRead))
      if c>maxScore {
        bestGuess = i
        maxScore = c
      }
    }
    if bestGuess > 0 {
      fmt.Printf("\n*** Best guess is therefore %q.\n\n", candidates[bestGuess])
    } else {
      fmt.Println("\n*** Based on this no separator has been found, please try other options.\n")
    }
  } else {
    fmt.Println("This seems to be an empty file!")
  }
}

func readUpToNLines(filename string, limit int) (count int, result[]string) {
  if rawBytes, err := ioutil.ReadFile(filename); err != nil {
    log.Fatal(">>!!", err)
  } else {
    lines := strings.Split(string(rawBytes),"\n")
    for _, line:= range lines {
      line = strings.TrimSpace(line)
      if line!= "" {
        result = append(result, line)
        count++
      }
      if count>=limit {
        break
      }
    }
  }
  return count, result
}

func createCount (lines []string, linesRead int, candidates []string) [][]int{
  counts := make([][]int, len(candidates))
  for indexS, separator := range candidates {
    counts[indexS] = make([]int, linesRead)
    for lineIndex, line := range lines {
      counts[indexS][lineIndex] = strings.Count(line, separator)
    }
  }
  return counts
}

func guessSeparator(table [][]int, separator []string) []int {
  /* there is one situation overlooked here:
      if one of the separators is used on the last line only, it will not be considered (counted)
      it is possible to correct for this, but in any case, the chance of this separator to win is
      minimal
  */
  nLines := len(table[0])
  nEqual := make([]int, len(table))
  for indexSep, counts := range table {
    max := 0
    for i:=0;i<nLines-1;i++ {
      if counts[i]==0 { continue }
      c:=1
      for j:=i+1;j<nLines; j++ {
        if counts[i]==counts[j] {
          c++
        }
      }
      if c>max {
        max = c
      }
    }
    if max>nEqual[indexSep] {
      nEqual[indexSep] = max
    }
  }
  return nEqual
}
