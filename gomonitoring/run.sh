#!/bin/bash


go build -o gomonitoring cmd/web/*.go && ./gomonitoring \
-dbuser='postgres' \
-pusherHost='localhost' \
-pusherKey='123abc' \
-pusherSecret='abc123' \
-pusherApp="1"
-pusherPort="4001"
-pusherSecure=false