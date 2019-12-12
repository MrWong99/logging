package main

import (
	"flag"
	"log"
	fsnotify "github.com/fsnotify/fsnotify"
)

func main() {
	var logFolder string

	flag.StringVar(&logFolder, "log-folder", "./logs", "specify folder to read logs from. Default is ./logs")
	flag.Parse()

	log.Printf("Reading logs from folder %s", logFolder)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(logFolder)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
