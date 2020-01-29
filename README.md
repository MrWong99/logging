![Go](https://github.com/MrWong99/logging/workflows/Go/badge.svg) [![Release](https://github.com/MrWong99/logging/workflows/Release/badge.svg)](https://github.com/MrWong99/logging/releases) [![Docker](https://github.com/MrWong99/logging/workflows/Docker/badge.svg)](https://hub.docker.com/repository/docker/mrwong99/logging)

# gLogCollector

This is a simple tool to collect and forward logs. It is able to efficiently read
new logs written in any given folder and then forward these logs via a gRPC connection.

## Usage

The executable has three **parameters**:

* `--log-folders`: specify folders to read logs from. **The folders need to exist!** Default is ./logs,./log
* `--grpc-addresses`: specify the backends to send logs to. Seperate multiple servers with `,` e.g. `localhost:8080,example.web.com:234`. Default is none
* `--grpc-port`: specify the port of the gRPC server to start. If none is set no server will be started. Default is none

You can also provide the same configuration via **environment variables**:

* `LOG_FOLDERS`
* `GRPC_ADDRESSES`
* `GRPC_PORT`

The environment variables override the cli parameters.

The log files need to be appended to via software or in the editor.
**Most modern IDEs and editors will perform [two write operations on save](https://github.com/fsnotify/fsnotify/issues/304), which results in the entire file being logged even if only one line changes!**

To start the application as server for a specific port:

**Windows:**
```batch
logging.exe --log-folders . --grpc-port 8080
```

**Linux:**
```bash
./logging --log-folders . --grpc-port 8080
```

you can then connect a client that sends logs in folder `./logs` to the server:

**Windows:**
```batch
logging.exe --log-folders ./logs --grpc-addresses 127.0.0.1:8080
```

**Linux:**
```batch
./logging --log-folders ./logs --grpc-addresses 127.0.0.1:8080
```

## Docker

You can also pull the image from Dockerhub:

https://hub.docker.com/repository/docker/mrwong99/logging

## Generate from Protobuf file

To generate the gRPC code you need to [download and install the Google Protocol Buffer Compiler](https://developers.google.com/protocol-buffers/docs/downloads).

Afterwars you have to install the following go packages:
```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

Finally you can regenerate the [communication/log.pb.go] file using

**Linux:**
```bash
protoc -I communication/ communication/log.proto --go_out=plugins=grpc:communication
```

**Windows:**
```batch
protoc -I .\communication\ .\communication\log.proto --go_out=plugins=grpc:communication
```
