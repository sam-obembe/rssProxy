#!/bin/sh

rm rssProxy
GOOS=linux GOARCH=amd64 go build
podman build -t rssproxyapi .
