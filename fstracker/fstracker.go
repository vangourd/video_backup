package main

import (
    "io/ioutil"
    "io"
    "log"
    "os"
    "path/filepath"
)

func main() {
    source := "Z:\\VideoArchives\\Archiver1"
    target := "Y:\\"
    scanDir(source, target)
}

func itemMissing(filepath string) bool{
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        return true
    }

    return false
}

func copyFile(sourcePath string,targetPath string) error {

    from, err := os.Open(sourcePath)
    if err != nil {
        log.Fatal(err)
        return err
    }
    defer from.Close()

    log.Printf("Cloning: %s to %s", sourcePath,targetPath)
    to, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)    
    if err != nil {
        log.Fatal(err)
        return err
    }
    defer to.Close()

    _, err = io.Copy(to, from)
    if err != nil {
        log.Fatal(err)
        return err
    }
    return nil
}

func getChildItems(filepath string) ([]os.FileInfo, error){
    files, err := ioutil.ReadDir(filepath)
    if err != nil {
        log.Fatalf("ioutil.ReadAll: %v", err)
        return nil, err
    }
    return files, nil
}

func copyItem(sourcePath string, targetPath string) error {

    fileInfo, err := os.Stat(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	switch fileInfo.Mode() & os.ModeType{
    case os.ModeDir:
        log.Println("copying directory")
		os.Mkdir(targetPath, os.ModeDir )
        scanDir(sourcePath,targetPath)
        return nil
	default:
        log.Println("copy regular file")
        copyFile(sourcePath, targetPath)
        return nil
	}
}

func scanDir(sourceBase string,targetBase string) error {

    // Get all items of current node 
    items, err := getChildItems(sourceBase)
    if err != nil {
        log.Fatal(err)
    }

    // For each item in node
    for _, file := range items {

        // Build references for current item source and target operations
        sourcePath := filepath.Join(sourceBase,file.Name())
        targetPath := filepath.Join(targetBase, file.Name())

        // See if current item exists in target, 
        if itemMissing(targetPath){
            // copyItem handles directories and files
            copyItem(sourcePath, targetPath)
        } else {
            scanDir(sourcePath,targetPath)
        }
    }

    return nil

    // if current.itemType == "dir" {
    //     files, err := ioutil.ReadDir(current.abspath)
    //     if err != nil {
    //         log.Fatalf("ioiutil.ReadDir: %s",err)
    //     }
    //     for _, file := range files {
    //         child, err := buildTree(
    //             filepath.Join(current.abspath,file.Name()),
    //             &current)
    //         if err != nil {
    //             log.Fatal(err)
    //         } else {
    //             current.childItems = append(current.childItems, child)
    //         }
    //     }
    // }

}
