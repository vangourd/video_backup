package main

import (
    "path/filepath"
    "log"
    "os"
    "io/ioutil"
    "fmt"
    "time"
)

type Item struct {
    name        string
    itemType    string
    abspath     string
    size        int64
    parent      *Item
    childItems  []*Item
}

func main() {
    var start_time = time.Now()
    Tree, err := buildTree("Z:\\VideoArchives\\Archiver1",nil)
    if err != nil {
        log.Fatal(err)
    }
    var end_time = time.Now()
    log.Print(Tree)
    log.Printf("Started: %s\r\nFinished: %s",start_time,end_time)
}

func buildTree(path string, parent *Item) (*Item, error) {

    var current Item

    stat, err := os.Stat(path)

    if err != nil{
        log.Fatalf("os.Stat error: %d",err)
    }

    current.name = stat.Name()
    current.abspath = path

    if stat.IsDir() {
        current.itemType = "dir"
    } else {
        current.itemType = "file"
        current.size = stat.Size()
    }

    current.parent = parent

    if current.itemType == "dir" {
        files, err := ioutil.ReadDir(current.abspath)
        if err != nil {
            log.Fatalf("ioiutil.ReadDir: %s",err)
        }
        for _, file := range files {
            child, err := buildTree(
                filepath.Join(current.abspath,file.Name()),
                &current)
            if err != nil {
                log.Fatal(err)
            } else {
                current.childItems = append(current.childItems, child)
            }
        }
    }

    return &current, nil
}

var biggestsize int64

func walkFunc(path string, info os.FileInfo, err error) error{
    if filepath.Ext(path) == ".g64" {
        // Splitting path to file into a list
        item, err := os.Stat(path)
        if err != nil {
            log.Fatal(err)
        }
        // To compare to todays date
        dt := time.Now()
        today := dt.Format("2006-01-02")
        if item.Name() == today {
            fmt.Println(path)
        }
        
    }
    return nil
}