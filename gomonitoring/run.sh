#!/bin/bash


go build -o gomonitoring cmd/web/*.go && ./gomonitoring \
-dbuser='postgres' \
-pusherHost='localhost' \
-pusherKey='abc123' \
-pusherSecret='123abc' \
-pusherApp="1"
-pusherPort="4001"
-pusherSecure=false