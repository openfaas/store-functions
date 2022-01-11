FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/classic-watchdog:0.1.4 as watchdog
FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.15

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN apk add --no-cache py-pip bash ca-certificates
RUN pip install --upgrade youtube_dl

# Add non root user
RUN mkdir -p /home/app
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

COPY entry.sh .
ENV fprocess="./entry.sh"

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
