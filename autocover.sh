#!/bin/sh
WATCH=$(find . -name '*.go')
while inotifywait -e move_self -e modify $WATCH; do
    make cover
done
