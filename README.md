:exclamation: **THIS IS IN DEVELOPMENT AND NOT EVEN REMOTELY FINISHED!** :exclamation:

# gLogCollector

This is a simple tool to collect and forward logs. When finished it should be able to efficiently read
new logs written in any given folder and then forward these logs via a gRPC connection.

## Usage

The executable has two parameters:

* `--log-folders`: specify folders to read logs from. Default is ./logs,./log
* `--grpc-addresses`: specify the backend to send logs to. Default is 127.0.0.1:8080

The log files need to be appended to via software or in the editor.
**Most modern IDEs and editors will perform [two write operations on save](https://github.com/fsnotify/fsnotify/issues/304), which results in the entire file being logged once only one line changes!**

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
