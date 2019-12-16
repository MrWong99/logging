package files

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	fsnotify "github.com/fsnotify/fsnotify"
)

// The FolderLoader can be used to keep watching for changes in a folder.
// It allows reading logs as they come in, by line or just the entire files.
type FolderLoader struct {
	LogFolders    []string
	fileBytesRead map[string]int64
	close         chan bool
}

// FileError is an error related to a file, it will be appended to the message.
type FileError struct {
	err      error
	filePath string
}

// Error returns the error that occured starting with the file that failed.
func (err FileError) Error() string {
	return "Path: " + err.filePath + "\n" + err.err.Error()
}

func (loader *FolderLoader) getLastText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return "", FileError{err: err, filePath: filePath}
	}
	info, err := file.Stat()
	if err != nil {
		return "", FileError{err: err, filePath: filePath}
	}
	// Check what we need to read from the file
	readBytes := loader.fileBytesRead[filePath]
	log.Printf("file size: %d\nread bytes: %d", info.Size(), readBytes)
	bytesToRead := info.Size() - readBytes
	// If the file became smaller it probably was deleted before
	if bytesToRead < 0 {
		loader.fileBytesRead[filePath] = info.Size()
		readBytes = 0
		bytesToRead = info.Size()
	} else {
		loader.fileBytesRead[filePath] = bytesToRead + readBytes
	}
	log.Printf("Next bytes to read: %d", loader.fileBytesRead[filePath])
	// Read from last location
	_, err = file.Seek(readBytes, 0)
	if err != nil {
		return "", FileError{err: err, filePath: filePath}
	}
	byteContent := make([]byte, bytesToRead)
	_, err = io.ReadAtLeast(file, byteContent, int(bytesToRead))
	if err != nil {
		return "", FileError{err: err, filePath: filePath}
	}
	return string(byteContent), nil
}

// StartWatching for logs an notify new line channels added with AddNewLineChan when
// any file changes.
func (loader *FolderLoader) StartWatching() (chan string, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		watcher.Close()
		return nil, err
	}
	outputChannel := make(chan string)
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
					log.Println("Sending to channel '" + text + "'")
					outputChannel <- text
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			case <-loader.close:
				log.Printf("closing watcher for folders %s", loader.LogFolders)
				close(outputChannel)
				watcher.Close()
			}
		}
	}()
	for _, folder := range loader.LogFolders {
		err := watcher.Add(folder)
		if err != nil {
			return nil, FileError{err: err, filePath: folder}
		}
	}
	return outputChannel, nil
}

// NewFolderLoader creates a new FolderLoader.
func NewFolderLoader(logLocations []string) *FolderLoader {
	// Create object, append "/" or "\" to log location if needed.
	loader := &FolderLoader{LogFolders: logLocations}
	for index, location := range loader.LogFolders {
		path, err := filepath.Abs(location)
		if err != nil {
			log.Printf("Could not parse path '%s': %s", location, err)
		}
		separator := string(os.PathSeparator)
		if !strings.HasSuffix(location, separator) {
			loader.LogFolders[index] = path + separator
		} else {
			loader.LogFolders[index] = path
		}
	}
	// Initialize map and channel
	loader.close = make(chan bool)
	loader.fileBytesRead = make(map[string]int64)
	return loader
}

// Close stops and closes this loader. It also notifies all channels to be closed.
func (loader *FolderLoader) Close() {
	loader.close <- true
}
