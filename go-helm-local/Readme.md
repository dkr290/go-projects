# go-helm-local-registry

** Table Of Contents **

- [Overview](#Overview)
- [Build](#Build)
- [Usage](#Usage)
- [Examples](#Examples)

## Overview

golang application that will be able to pull the helm chart extract it and get from values.yaml and from Chart.yaml the images and repush them to a local repository.
It can find most the images if the pattern in helm is followed
It can be improved to be able to cover different helm cases and not only standard repository pattern
The idea is to avoind relying to docker.io and mainly because docker.io have now rate limits

## Build

To build the application you can use the following in windows and linux

```
env GOOS=windows GOARCH=amd64 go build -o bin/go-helm-local.exe
env GOOS=linux GOARCH=amd64 go build -o bin/go-helm-local
```

## Usage

```
go-helm-local -help
  -chart string the chart like grafana/loki , this is the actual helm chart from the repo
  -version  string version like 6.27.0  , the version to download
  -dest string like -dest ./temp , this is a temporary destination for unpacking the tar.gz helm chart
  -registry string like <internal container registry>
  -username the registry username
  -password the registry password
  -help Show this help message
```

## Examples

```
helm repo list
helm repo update
helm repo add [NAME] [URL] [flags]   #repository that you want to get images from helm chart

go-helm-local -chart grafana/loki -version 6.27.0 -dest /tmp/ -registry someregistry.azurecr.io -username user -password <some password here>

```
