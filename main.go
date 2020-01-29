package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	communication "github.com/MrWong99/logging/communication"
	config "github.com/MrWong99/logging/config"
	loader "github.com/MrWong99/logging/files"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type changeReader struct {
	logChannel chan loader.LogMessage
	loader     *loader.FolderLoader
}

func (reader *changeReader) listenForChange() {
	for {
		logmsg, more := <-reader.logChannel
		if more {
			log.Printf("Logged Text:%s\n", logmsg.Text)
		} else {
			log.Println("I guess I was closed...")
			return
		}
	}
}

func (reader *changeReader) listenForChangeAndForward(servers []string) {
	clients := make([]communication.LogReceiverClient, len(servers))
	for index, server := range servers {
		conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Printf("Could not connect to server '%s', error: %s", server, err)
			clients[index] = nil
			continue
		}
		defer conn.Close()
		client := communication.NewLogReceiverClient(conn)
		clients[index] = client
	}
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		select {
		case logmsg, more := <-reader.logChannel:
			if more {
				for _, client := range clients {
					// Send new log to this client in goroutine
					go func(client communication.LogReceiverClient) {
						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()
						message := communication.LogText{
							LogFile: &communication.LogPath{
								Path: logmsg.FilePath,
							},
							LoggedAt: &timestamp.Timestamp{
								Seconds: logmsg.Timestamp.Unix(),
								Nanos:   int32(logmsg.Timestamp.Nanosecond()),
							},
							LogMessage: logmsg.Text,
						}
						_, err := client.ReceiveLoggedText(ctx, &message)
						if err != nil {
							log.Printf("Could not send to a gRPC server. Error: %s", err)
						}
					}(client)
				}
			} else {
				log.Println("I guess I was closed...")
				return
			}
		case <-c:
			fmt.Println("Closing clients...")
			return
		}
	}
}

// LogFileService can be used as LogCollector server.
type LogFileService struct {
	folderLoader *loader.FolderLoader
}

// GetFileList returns a list of existing files for a given directory.
func (server LogFileService) GetFileList(ctx context.Context, dir *communication.LogPath) (*communication.FileList, error) {
	return nil, errors.New("Not implemented")
}

// ReadLogFile reads the given file and return its contents.
func (server LogFileService) ReadLogFile(ctx context.Context, file *communication.LogPath) (*communication.LogFile, error) {
	content, err := server.folderLoader.ReadFile(file.GetPath())
	if err != nil {
		return nil, err
	}
	return &communication.LogFile{
		Path:    file,
		Content: content,
	}, nil
}

// LogTextReceiver can be used to start a log receiver gRPC server
type LogTextReceiver struct {
}

// ReceiveLoggedText receives a logs message via gRPC and displays it in STDOUT.
func (server LogTextReceiver) ReceiveLoggedText(ctx context.Context, text *communication.LogText) (*any.Any, error) {
	fmt.Printf("Received log in file %s timestamp %s with content: %s\n", text.GetLogFile().GetPath(), text.GetLoggedAt(), text.GetLogMessage())
	return &any.Any{}, nil
}

func main() {
	log.Println("Starting applcation...")
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

	port := config.Get().GrpcPort
	if port > 0 {
		startGrpcServer(port, loader)
	} else {
		serversToConnect := config.Get().GrpcBackends
		if len(serversToConnect) <= 0 || serversToConnect[0] == "" {
			reader.listenForChange()
		} else {
			reader.listenForChangeAndForward(serversToConnect)
		}
	}
}

func startGrpcServer(port int, loader *loader.FolderLoader) {
	ctx := context.Background()
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Panicf("Could not start gRPC server! %s", err)
	}
	server := grpc.NewServer()
	communication.RegisterLogCollectorServer(server, LogFileService{
		folderLoader: loader,
	})
	communication.RegisterLogReceiverServer(server, LogTextReceiver{})
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	server.Serve(listen)
}
