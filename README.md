# Genghiskhan

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