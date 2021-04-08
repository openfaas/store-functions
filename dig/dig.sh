#!/usr/bin/env sh


# write log message to stderr
log() {
    echo "$@" 1>&2;
}

IFS='' read -d '' -r request
response=$(dig "$request" +short 2>&1)
status=$?

log "request=\"$request\" response=\"$response\" content_type=\"$content_type\""

if [ "$content_type" = "application/json" ]; then
    key="error"
    if [ $status -eq 0 ]; then
        key="${response_key:-response}"
    fi

    echo "{\"$key\": \"$response\"}"
else
    echo "$response"
fi