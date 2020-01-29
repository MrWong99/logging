package config

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

var globalConfig ApplicationConfig

// ApplicationConfig stores all of the input parameters.
type ApplicationConfig struct {
	LogFolders   []string // Folders that should be watched for changes.
	GrpcBackends []string // gRPC backends to send data to.
	GrpcPort     int      // Port for the gRPC server
}

// Init the configuration from the cli parameters.
func Init() error {
	var logFolder string
	var grpcAddresses string
	var grpcPort string

	flag.StringVar(&logFolder, "log-folders", "./logs,./log", "specify folders to read logs from. Default is ./logs,./log")
	flag.StringVar(&grpcAddresses, "grpc-addresses", "", "specify the backends to send logs to. Default is none")
	flag.StringVar(&grpcPort, "grpc-port", "-1", "specify the port of the gRPC server to start. If none is set no server will be started. Default is none")
	flag.Parse()

	folders := strings.Split(logFolder, ",")
	backends := strings.Split(grpcAddresses, ",")
	port, err := strconv.Atoi(grpcPort)
	if err != nil {
		log.Printf("Could not read port argument '%s' error: %s", grpcPort, err)
		port = -1
	}
	foldersEnv := os.Getenv("LOG_FOLDERS")
	backendsEnv := os.Getenv("GRPC_ADDRESSES")
	portEnv := os.Getenv("GRPC_PORT")
	if foldersEnv != "" {
		folders = strings.Split(foldersEnv, ",")
	}
	if backendsEnv != "" {
		backends = strings.Split(backendsEnv, ",")
	}
	if portEnv != "" {
		port2, err := strconv.Atoi(portEnv)
		if err != nil {
			log.Printf("Could not read port argument '%s' error: %s", portEnv, err)
		} else {
			port = port2
		}
	}
	globalConfig = ApplicationConfig{
		LogFolders:   folders,
		GrpcBackends: backends,
		GrpcPort:     port,
	}
	log.Printf("Reading logs from folders %s", folders)
	log.Printf("Sending logs to backends %s", backends)
	if port > 0 {
		log.Printf("Starting gRPC server on port %s", strconv.Itoa(port))
	}
	return nil
}

// Get the configuration for this app. You should run 'Init()' before.
func Get() ApplicationConfig {
	return globalConfig
}
