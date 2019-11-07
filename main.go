package main

import (
	"os"
	"log"

	"video_backup/email"
	"video_backup/config"
	"video_backup/fstracker"
)

func main() {
	//count := os.Args[1]
	// load my settings
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// create email notifier
	//notifier, err := email.NewNotifier(conf)
	if err != nil {
		log.Fatal(err)
	}

	// build tree of source directory
	// build tree of target directory
	// compare trees find differences
	// build queue of work
	// go routine to do work
	// while true
	// 	fsnotify
	//  when change send to work channel
}

func buildTree(conf config.Config) {
	root := conf.Directory.Name
	println(root)
	// enumerate items in path
}

func discoverBranch(path string) {

}
