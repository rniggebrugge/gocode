package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "runtime"
    "sort"
    "strings"
    "unicode"
    "unicode/utf8"
)


func main(){
  if len(os.Args)==1 || os.Args[1] == "-h" || os.Args[1] == "--help"{
    fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }
  
  frequencyWord := map[string]int{}
  fmt.Println("\n-------START READING-----\n")
  for _, filename := range commandLineFiles(os.Args[1:]){
    updateFrequencies(filename, frequencyWord)
  }
  fmt.Println("\n-------DONE READING-----\n")
  
  reportByWords(frequencyWord)
  wordsForFrequency := invertMap(frequencyWord)
  reportByFrequency(wordsForFrequency)
}

func commandLineFiles(files []string) []string {
  if runtime.GOOS == "windows" {
    args := make([]string,0, len(files))
    for _, name := range files {
      if matches, err:= filepath.Glob(name); err != nil {
        args = append(args, name)
      } else if matches != nil {
        args = append(args, matches...)
      }
    }
    return args
  }
  return files
}

func updateFrequencies(filename string, frequencyWord map[string]int){
  var file *os.File
  var err error
  if file, err = os.Open(filename); err != nil {
    log.Println("failed to open the file: ", err)
    return
  }
  defer file.Close()
  fmt.Printf("--- reading file : %20s ---\n", filename)
  readAndUpdateFrequency(bufio.NewReader(file), frequencyWord)
}

func readAndUpdateFrequency(reader *bufio.Reader,frequencyWord map[string]int){
  for {
    line, err := reader.ReadString('\n')
    for _, word := range SplitOnNonLetters(strings.TrimSpace(line)){
      if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
        frequencyWord[word] += 1
      }
    }
    if err != nil {
      if err != io.EOF {
        log.Println("failed to finish reading the file: ", err)
      }
      break
    }
  }
}

func SplitOnNonLetters(s string) []string {
  notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
  return strings.FieldsFunc(s, notALetter)
}

func reportByWords(table map[string]int) {
  words := make([]string,0, len(table))
  wordWidth, freqWidth := 0, 0
  for word, frequency := range table {
    words = append(words, word)
    if width := utf8.RuneCountInString(word); width > wordWidth {
      wordWidth = width
    }
    if width := len(fmt.Sprint(frequency)); width > freqWidth {
      freqWidth = width
    }
  }
  sort.Strings(words)
  gap := wordWidth + freqWidth - len("Word") - len("Frequency")
  fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
  for _, word := range(words){
    fmt.Printf("%*s %*d\n", wordWidth, word, freqWidth, table[word])
  }
}

func invertMap(table map[string]int) map[int][]string {
  stringsPerInt := make(map[int][]string, len(table))
  for key, value := range table {
    stringsPerInt[value] = append(stringsPerInt[value], key)
  }
  return stringsPerInt
}

func reportByFrequency(table map[int][]string){
  frequencies := make([]int,0 , len(table))
  for frequency := range table {
    frequencies = append(frequencies, frequency)
  }
  sort.Ints(frequencies)
  width := len(fmt.Sprint(frequencies[len(frequencies)-1]))+7
  fmt.Println("Frequency -> Words")
  for _, frequency := range frequencies {
    words := table[frequency]
    sort.Strings(words)
    fmt.Printf("%*d    %s\n", width, frequency, strings.Join(words,", "))
  }
}