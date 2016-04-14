package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

type Song struct {
    Title string
    Filename string
    Seconds int
}

func (song Song) AsPls(n int) string {
    txt := fmt.Sprintf("File%d=%s\nTitle%d=%s\nLength%d=%d", n, song.Filename, n, song.Title, n, song.Seconds)
    return txt
}

func (song Song) AsM3u() string {
    return fmt.Sprintf("#EXTINF:%d,%s\n%s", song.Seconds, song.Title, song.Filename)
}

func main(){
  
    b:= "Remco"
    a:= []rune(b)
    fmt.Printf("a=%v(%T); b=%v(%T) \n",a,a,b,b)
    if len(os.Args) == 1 || (!strings.HasSuffix(os.Args[1],".m3u") && !strings.HasSuffix(os.Args[1],".pls")){
        fmt.Printf("usage: %s <file.[pls|m3u]>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }
    if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
        log.Fatal(err)
    } else {
        fromM3U := strings.HasSuffix(os.Args[1], ".m3u")
        songs := readPlaylist(string(rawBytes), fromM3U)
        if fromM3U {
            writePlsPlaylist(songs)
        } else {
            writeM3uPlaylist(songs)
        }
    }
}

func readPlaylist(data string, fromM3U bool) (songs []Song) {
    var song Song
    for _, line := range strings.Split(data, "\n") {
        line = strings.TrimSpace(line)
        if line =="" || line == "#EXTM3U" || line == "[playlist]" {
            continue
        }
        if fromM3U {
            if strings.HasPrefix(line, "#EXTINF:"){
                song.Title, song.Seconds = parseExtInfLine(line)
            } else {
                song.Filename = strings.Map(platformDir, line)
            }
        } else {
            if strings.HasPrefix(line, "File") {
                song.Filename = strings.Map(platformDir, readPlsLine(line))
            } else if strings.HasPrefix(line, "Title") {
                song.Title = readPlsLine(line)
            } else if strings.HasPrefix(line, "Length") {
                temp := readPlsLine(line)
                if seconds , err := strconv.Atoi(temp); err != nil {
                    log.Printf("failed to read the duration: %v\n", err )
                    song.Seconds = -1
                } else {
                    song.Seconds = seconds
                }
            }
        }
        if song.Filename != "" && song.Seconds != 0 && song.Title != "" {
            songs = append(songs, song)
            song = Song{}
        }
    }
    return songs
}

func readPlsLine(line string) string {
    if i:=strings.Index(line, "="); i>-1 {
        return line[i+1:]
    }
    return ""
}

func parseExtInfLine(line string) (title string, seconds int) {
    if i:= strings.IndexAny(line, "-0123456789"); i>-1 {
        const separator = ","
        line = line[i:]
        if j := strings.Index(line, separator); j >-1 {
            title = line[j+len(separator):]
            var err error
            if seconds, err = strconv.Atoi(line[:j]); err != nil {
                log.Printf("failed to read the duration of '%s': %v\n", title, err)
                seconds = -1
            }
        }
    }
    return title, seconds
}

func platformDir(char rune) rune {
    if char == '/' || char == '\\' {
        return filepath.Separator
    }
    return char
}

func writePlsPlaylist(songs []Song){
    fmt.Println("[Playlist]")
    for i, song := range songs {
        i++
        fmt.Println(song.AsPls(i))
    }
    fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func writeM3uPlaylist(songs []Song){
    fmt.Println("#EXTM3U")
    for _, song := range songs {
        fmt.Println(song.AsM3u())
    }
}

    

