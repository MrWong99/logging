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
