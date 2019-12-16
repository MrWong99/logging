package main

import (
	"flag"
	"log"
	"strings"

	loader "github.com/MrWong99/logging/files"
)

type changeReader struct {
	logChannel chan string
	loader     *loader.FolderLoader
}

func (reader *changeReader) listenForChange() {
	for {
		logText, more := <-reader.logChannel
		if more {
			log.Printf("Logged Text:\n%s\n", logText)
			if strings.Contains(logText, "Close()") {
				log.Println("Text contained 'Close()' so im closing")
				reader.loader.Close()
			}
		} else {
			log.Println("I guess I was closed...")
			return
		}
	}
}

func main() {
	var logFolder string

	flag.StringVar(&logFolder, "log-folders", "./logs,./log", "specify folders to read logs from. Default is ./logs,./log")
	flag.Parse()

	folders := strings.Split(logFolder, ",")
	log.Printf("Reading logs from folders %s", folders)

	loader := loader.NewFolderLoader(folders)
	logChan, err := loader.StartWatching()
	if err != nil {
		log.Fatal(err)
	}
	reader := changeReader{logChannel: logChan, loader: loader}
	reader.listenForChange()
}
