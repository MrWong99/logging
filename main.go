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
	var grpcAddresses string

	flag.StringVar(&logFolder, "log-folders", "./logs,./log", "specify folders to read logs from. Default is ./logs,./log")
	flag.StringVar(&grpcAddresses, "grpc-addresses", "127.0.0.1:8080", "specify the backend to send logs to. Default is 127.0.0.1:8080")
	flag.Parse()

	folders := strings.Split(logFolder, ",")
	backends := strings.Split(grpcAddresses, ",")
	log.Printf("Reading logs from folders %s", folders)
	log.Printf("Sending logs to backends %s", backends)

	loader := loader.NewFolderLoader(folders)
	logChan, err := loader.StartWatching()
	if err != nil {
		log.Fatal(err)
	}
	reader := changeReader{logChannel: logChan, loader: loader}
	reader.listenForChange()
}
