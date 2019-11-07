package main

import (
	"os"
	"log"

	"video_backup/email"
	"video_backup/config"
)

func main() {
	count := os.Args[1]
	// load my settings
	// create email notifier
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	notifier, err := email.NewNotifier(conf)
	if err != nil {
		log.Fatal(err)
	}
	notifier.Send("Want to see me count to 5?", count)
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
