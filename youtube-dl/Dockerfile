FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/classic-watchdog:0.3.3 as watchdog

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.21.3 as ship

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Download yt-dlp, an updated alternative to youtube-dl, from Alpine repository
# since it is packaged as static binary compatible with Alpine.
RUN apk add --no-cache ca-certificates yt-dlp

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
