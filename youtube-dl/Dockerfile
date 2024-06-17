FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/classic-watchdog:0.3.1 as watchdog

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.19.1 as ship

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN apk add --no-cache curl bash ca-certificates

# youtube-dl appears to be broken
# https://github.com/ytdl-org/youtube-dl/issues/31530
# Switching to yt-dlp
# RUN pip install --upgrade yt-dlp

RUN curl -SLs https://get.arkade.dev | sh && \
    arkade get yt-dlp && \
    mv $HOME/.arkade/bin/yt-dlp /usr/local/bin/ && \
    chmod +x /usr/local/bin/yt-dlp && \
    rm -rf /tmp/* && rm -rf /usr/local/bin/arkade

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
