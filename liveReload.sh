#!/usr/bin/env bash

# in foreground, continously run app
while true; do
    cd internal/ui
    ../../bin/bubbleWaffle tui 
done
