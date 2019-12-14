package main

import (
	loader "github.com/MrWong99/logging/files"
	"flag"
	"log"
	"strings"
)

func main() {
	var logFolder string

	flag.StringVar(&logFolder, "log-folders", "./logs,./log", "specify folders to read logs from. Default is ./logs,./log")
	flag.Parse()

	folders := strings.Split(logFolder, ",")
	log.Printf("Reading logs from folders %s", folders)

	for _, folder := range folders {
		loader := loader.NewFolderLoader(folder)
		logChan := make(chan string)
		loader.AddNewLineChan(logChan)
		go func() {
			for {
				logText, more := <-logChan
				if more {
					log.Printf("Logged Text:\n%s\n", logText)
				} else {
					log.Println("I guess I was closed...")
					return
				}
			}
		}()
		loader.StartWatching()
	}

	done := make(chan bool)
	<-done
}
