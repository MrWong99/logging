package config

import (
	"flag"
	"log"
	"strings"
)

var globalConfig ApplicationConfig

// ApplicationConfig stores all of the input parameters.
type ApplicationConfig struct {
	LogFolders   []string // Folders that should be watched for changes.
	GrpcBackends []string // gRPC backends to send data to.
}

// Init the configuration from the cli parameters.
func Init() error {
	var logFolder string
	var grpcAddresses string

	flag.StringVar(&logFolder, "log-folders", "./logs,./log", "specify folders to read logs from. Default is ./logs,./log")
	flag.StringVar(&grpcAddresses, "grpc-addresses", "127.0.0.1:8080", "specify the backend to send logs to. Default is 127.0.0.1:8080")
	flag.Parse()

	folders := strings.Split(logFolder, ",")
	backends := strings.Split(grpcAddresses, ",")
	globalConfig = ApplicationConfig{
		LogFolders:   folders,
		GrpcBackends: backends,
	}
	log.Printf("Reading logs from folders %s", folders)
	log.Printf("Sending logs to backends %s", backends)
	return nil
}

// Get the configuration for this app. You should run 'Init()' before.
func Get() ApplicationConfig {
	return globalConfig
}
