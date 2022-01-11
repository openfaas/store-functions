#!/bin/sh

youtubeID=""
if [ -n "$1" ] ; then
  youtubeID=$1
else
  youtubeID=$(cat /dev/stdin)
fi

trimmedYoutubeID=$(echo "$youtubeID" | tr -d '\n')

# youtube-dl appears to be broken
# https://github.com/ytdl-org/youtube-dl/issues/31530
# Switching to yt-dlp

yt-dlp "$trimmedYoutubeID" --no-warnings --quiet -o -
