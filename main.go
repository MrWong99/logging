package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	loader "github.com/MrWong99/logging/files"
)

type changeReader struct {
	logChannel chan string
}

func (reader *changeReader) listenForChange() {
	go func() {
		for {
			logText, more := <-reader.logChannel
			if more {
				log.Printf("Logged Text:\n%s\n", logText)
			} else {
				log.Println("I guess I was closed...")
				return
			}
		}
	}()
}

func main() {
	var logFolder string

	flag.StringVar(&logFolder, "log-folders", "./logs,./log", "specify folders to read logs from. Default is ./logs,./log")
	flag.Parse()

	folders := strings.Split(logFolder, ",")
	log.Printf("Reading logs from folders %s", folders)

	for _, folder := range folders {
		path, err := filepath.Abs(folder)
		if err != nil {
			log.Printf("Could not parse path '%s': %s", folder, err)
		}
		loader := loader.NewFolderLoader(path)
		logChan := make(chan string)
		reader := changeReader{logChannel: logChan}
		reader.listenForChange()
		loader.AddNewLineChan(*logChan)
		loader.StartWatching()
	}

	done := make(chan bool)
	<-done
}
