package main

import (
    "path/filepath"
    "os"
    "strings"
    "fmt"
    "time"
)

type VideoFile struct {
    name    string
    size    int
    status  string
}

func main() {
    //vids := []VideoFile
    filepath.Walk("Z:\\VideoArchives\\Archiver1",walkFunc)
}

var biggestsize int64

func walkFunc(path string, info os.FileInfo, err error) error{
    if filepath.Ext(path) == ".g64" {
        // Splitting path to file into a list
        listOfDirs :=(strings.Split(path,"\\"))
        // Getting the parent directory of the files we are interested in
        dateDir := listOfDirs[len(listOfDirs)-2:len(listOfDirs)-1]
        // To compare to todays date
        dt := time.Now()
        today := dt.Format("2006-01-02")
        if dateDir[0] == today {
            fmt.Println(path)
        }
        
    }
    return nil
}