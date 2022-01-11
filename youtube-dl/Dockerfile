FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/classic-watchdog:0.2.2 as watchdog

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.17.3 as ship

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN apk add --no-cache py-pip bash ca-certificates

# youtube-dl appears to be broken
# https://github.com/ytdl-org/youtube-dl/issues/31530
# Switching to yt-dlp
RUN pip install --upgrade yt-dlp

# Add non root user
RUN mkdir -p /home/app
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

COPY entry.sh .
ENV fprocess="/bin/sh ./entry.sh"

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

EXPOSE 8080

CMD ["fwatchdog"]
