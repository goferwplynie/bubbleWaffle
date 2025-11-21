#!/usr/bin/env bash

# watch code changes, trigger re-build, and kill process 
while true; do
    go build -o bin/bubbleWaffle ./main.go && pkill -f 'bin/bubbleWaffle'
    inotifywait -e attrib $(find . -name '*.go') || exit
done
