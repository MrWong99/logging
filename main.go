package main

import (
	"log"

	config "github.com/MrWong99/logging/config"
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
			log.Printf("Logged Text:%s\n", logText)
		} else {
			log.Println("I guess I was closed...")
			return
		}
	}
}

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	loader := loader.NewFolderLoader(config.Get().LogFolders)
	defer loader.Close()
	logChan, err := loader.StartWatching()
	if err != nil {
		log.Fatal(err)
	}
	reader := changeReader{logChannel: logChan, loader: loader}
	reader.listenForChange()
}
