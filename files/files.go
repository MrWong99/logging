package files

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	fsnotify "github.com/fsnotify/fsnotify"
)

// A LogMessage is created when something is logged.
type LogMessage struct {
	FilePath  string    // The file in which the message was logged.
	Timestamp time.Time // The time when it was logged.
	Text      string    // The text that was logged.
}

// The FolderLoader can be used to keep watching for changes in a folder.
// It allows reading logs as they come in, by line or just the entire files.
type FolderLoader struct {
	LogFolders    []string
	fileBytesRead map[string]int64  // Map of bytes read for a file
	close         chan chan<- error // Channel that will receive close call
	outputChannel chan LogMessage   // Logs are written to this channel
	isClosed      bool              // Set to true when Close() is first called
	mu            sync.Mutex        // Map access
}

// FileError is an error related to a file, it will be appended to the message.
type FileError struct {
	error
	filePath string
}

// Error returns the error that occured starting with the file that failed.
func (err FileError) Error() string {
	return "Path: " + err.filePath + "\n" + err.Error()
}

// NewFolderLoader creates a new FolderLoader.
func NewFolderLoader(logLocations []string) *FolderLoader {
	// Create object, append "/" or "\" to log location if needed.
	loader := &FolderLoader{
		LogFolders:    logLocations,
		close:         make(chan chan<- error, 1),
		fileBytesRead: make(map[string]int64),
	}
	separator := string(os.PathSeparator)
	for index, location := range loader.LogFolders {
		path, err := filepath.Abs(location)
		if err != nil {
			log.Printf("Could not parse path '%s': %s", location, err)
		}
		if !strings.HasSuffix(location, separator) {
			loader.LogFolders[index] = path + separator
		} else {
			loader.LogFolders[index] = path
		}
	}
	return loader
}

// ReadFile reads the given log file if it exists and is in the scope of this reader.
func (loader *FolderLoader) ReadFile(filePath string) (string, error) {
	path, err := filepath.Abs(filePath)
	if err != nil {
		return "", FileError{err, filePath}
	}
	contained := false
	for _, location := range loader.LogFolders {
		if strings.HasPrefix(path, location) {
			contained = true
		}
	}
	if !contained {
		return "", FileError{errors.New("File is not within scope of current folders"), filePath}
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", FileError{err, filePath}
	}
	return string(content), nil
}

// Close stops and closes this loader. It also notifies all channels to be closed.
func (loader *FolderLoader) Close() error {
	if loader.isClosed {
		return nil
	}
	loader.isClosed = true
	ch := make(chan error)
	loader.close <- ch
	return <-ch
}

// readLastText writes the last read text for a file to the output channel.
func (loader *FolderLoader) readLastText(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatal(FileError{err, filePath})
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal(FileError{err, filePath})
	}
	// Check what we need to read from the file
	loader.mu.Lock()
	readBytes := loader.fileBytesRead[filePath]
	bytesToRead := info.Size() - readBytes
	// If the file became smaller it probably was deleted before
	if bytesToRead < 0 {
		loader.fileBytesRead[filePath] = info.Size()
		readBytes = 0
		bytesToRead = info.Size()
	} else {
		loader.fileBytesRead[filePath] = bytesToRead + readBytes
	}
	loader.mu.Unlock()
	// Read from last location
	_, err = file.Seek(readBytes, 0)
	if err != nil {
		log.Fatal(FileError{err, filePath})
	}
	byteContent := make([]byte, bytesToRead)
	_, err = io.ReadAtLeast(file, byteContent, int(bytesToRead))
	if err != nil {
		log.Fatal(FileError{err, filePath})
	}
	text := string(byteContent)
	if len(text) > 0 {
		loader.outputChannel <- LogMessage{
			FilePath:  filePath,
			Timestamp: time.Now(),
			Text:      text,
		}
	}
}

// StartWatching for logs and return a channel with the log output.
func (loader *FolderLoader) StartWatching() (chan LogMessage, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		watcher.Close()
		return nil, err
	}
	outputChannel := make(chan LogMessage, 50)
	loader.outputChannel = outputChannel
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					filePath := event.Name
					loader.readLastText(filePath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			case ch := <-loader.close:
				close(outputChannel)
				ch <- watcher.Close()
			}
		}
	}()
	for _, folder := range loader.LogFolders {
		err := watcher.Add(folder)
		if err != nil {
			return nil, FileError{err, folder}
		}
	}
	return outputChannel, nil
}
