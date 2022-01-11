#!/bin/sh

youtubeID=""
if [ -n "$1" ] ; then
  youtubeID=$1
else
  youtubeID=$(cat /dev/stdin)
fi

trimmedYoutubeID=$(echo "$youtubeID" | tr -d '\n')

youtube-dl "$trimmedYoutubeID" --no-warnings --quiet -o -
