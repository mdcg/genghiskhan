# Genghiskhan

[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mdcg/genghiskhan/blob/master/LICENSE)

## Introduction

**Genghiskhan** is a TCP and UDP port scanner written in Go. Because the base language in which Genghiskhan was written has as its main functionality the easy handling of competing processes, the scanner is performed in a very fast.

Genghiskhan also makes an inference of the service that is running on a given port, and displays this in its final report. It is worth noting that the inference is based only on the "common sense" of the service that is running at the port, and no "banner capture" feature is used.

When using Genghiskhan, in addition to specifying the number of ports you want to scan, you can also choose between three types of scans:

* TCP ports; 
* UDP ports; 
* Both.

**PS:** *The number of ports you specify will result in a scan from port 1 to port N, where N is the number of ports requested. Unfortunately, we don't have a "port range scan" yet.*

Genghiskhan also allows you to scan based on IP or URL. If you do not enter a host, Genghiskhan will understand that you want to scan your own computer.

For details on how to use it, just compile Genghiskhan and use the `--help` or ` -h` flag. Below are some examples of how to create an executable based on your desired Operating System and Architecture.

## Generating the Genghiskhan executable

To generate the executable based on the Operating System and Architecture you want, just execute one of these commands:

``` 
env GOOS=windows GOARCH=amd64 go build -o genghiskhan .
env GOOS=linux GOARCH=amd64 go build -o genghiskhan .
env GOOS=darwin GOARCH=amd64 go build -o genghiskhan .
```

There are other ways to generate the executable based on an Operating System and Architecture. To get the complete list, just run:

``` 
go tool dist list
```

## Contributing

Feel free to do whatever you want with this project. :-)

*This program was developed for study purposes only. The author is not responsible for the illicit or indiscriminate use of this software.*
