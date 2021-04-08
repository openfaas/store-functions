FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/of-watchdog:0.8.4 as watchdog

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.13

RUN mkdir -p /home/app

RUN apk add --no-cache bind-tools

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

COPY --chown=app:app dig.sh .
RUN chmod +x dig.sh

# customize the json key name for the success responses
ENV response_key="ip_address"

# configure of-watchdog
# https://github.com/openfaas/of-watchdog#configuration
ENV fprocess="/home/app/dig.sh"
ENV mode="streaming"
ENV content_type="application/json"
ENV suppress_lock="false"
ENV prefix_logs="false"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]