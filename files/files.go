package files

import (
	"io"
	"log"
	"os"
	"strings"

	fsnotify "github.com/fsnotify/fsnotify"
)

// The FolderLoader can be used to keep watching for changes in a folder.
// It allows reading logs as they come in, by line or just the entire files.
type FolderLoader struct {
	LogFolder     string
	fileBytesRead map[string]int64
	newLines      []chan string
	close         chan bool
}

// Helper function to add to a channel
func extendChannelSlice(toExend []chan string, element chan string) []chan string {
	n := len(toExend)
	if n == cap(toExend) {
		newSlice := make([]chan string, len(toExend), 2*len(toExend)+1)
		copy(newSlice, toExend)
		toExend = newSlice
	}
	toExend = toExend[0 : n+1]
	toExend[n] = element
	return toExend
}

func (loader *FolderLoader) getLastText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}
	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	// Check what we need to read from the file
	readBytes := loader.fileBytesRead[filePath]
	log.Printf("file size: %d\nread bytes: %d", info.Size(), readBytes)
	bytesToRead := info.Size() - readBytes
	// If the file became smaller it probably was deleted before
	if bytesToRead > 0 {
		loader.fileBytesRead[filePath] = 0
		readBytes = 0
		bytesToRead = info.Size()
	}
	// Read from last location
	_, err = file.Seek(readBytes, 0)
	if err != nil {
		return "", err
	}
	byteContent := make([]byte, bytesToRead)
	_, err = io.ReadAtLeast(file, byteContent, int(bytesToRead))
	if err != nil {
		return "", err
	}
	return string(byteContent), nil
}

// StartWatching for logs an notify new line channels added with AddNewLineChan when
// any file changes.
func (loader *FolderLoader) StartWatching() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		watcher.Close()
		return err
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					filePath := event.Name
					log.Println("modified file:", filePath)
					text, err := loader.getLastText(filePath)
					if err != nil {
						log.Printf("Error while reading logs: %s", err)
					}
					for _, newLineChan := range loader.newLines {
						log.Println("Sending to channel '" + text + "'")
						newLineChan <- text
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			case <-loader.close:
				log.Println("closing watcher for folder " + loader.LogFolder)
				for _, newLineChan := range loader.newLines {
					close(newLineChan)
				}
				watcher.Close()
			}
		}
	}()
	return watcher.Add(loader.LogFolder)
}

// NewFolderLoader creates a new FolderLoader.
func NewFolderLoader(logLocation string) *FolderLoader {
	// Create object, append "/" to log location if needed.
	loader := &FolderLoader{LogFolder: logLocation}
	if !strings.HasSuffix(loader.LogFolder, "/") {
		loader.LogFolder = loader.LogFolder + "/"
	}
	// Initialize slice and channel
	loader.newLines = make([]chan string, 10)
	loader.close = make(chan bool)
	loader.fileBytesRead = make(map[string]int64)
	return loader
}

// AddNewLineChan adds a channel to FolderLoader that will receive any new log lines.
func (loader *FolderLoader) AddNewLineChan(newLineChan chan string) {
	loader.newLines = extendChannelSlice(loader.newLines, newLineChan)
}

// Close stops and closes this loader. It also notifies all channels to be closed.
func (loader *FolderLoader) Close() {
	loader.close <- true
}
