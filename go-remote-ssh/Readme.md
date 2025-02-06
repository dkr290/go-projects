# Introduction

## What is this tool for

- go-remoteconn tool Windows and Linux version.
- This is the tool that mange ssh connections based on the json file

# Build

### 1. Build Executable per operating system for example Linux amd64

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o go-remoteconn

### 2. For windows

env GOOS=windows GOARCH=amd64 go build -o go-remoteconn.exe

# Usage

```
go-remoteconn --help
go-remoteconn.exe --help


```

1. Exmaple usage

```

go-remoteconn -f connections.json
```
