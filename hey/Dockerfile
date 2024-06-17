FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/classic-watchdog:0.3.1 as watchdog
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.16-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Required to enable Go modules
RUN apk add --no-cache git

# Allows you to add additional packages via build-arg
ARG ADDITIONAL_PACKAGE
ARG CGO_ENABLED=0
ARG GO111MODULE="off"
ARG GOPROXY=""
ARG GOFLAGS=""

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/rakyll/
RUN git clone https://github.com/rakyll/hey --depth=1
WORKDIR /go/src/github.com/rakyll/hey/
RUN mkdir -p bin

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=${CGO_ENABLED} \
    go build --ldflags "-s -w" -a -installsuffix cgo -o bin/hey .

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.19.1
RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=builder /go/src/github.com/rakyll/hey/bin/hey         .
COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN chown -R app /home/app

USER app

ENV fprocess="xargs ./hey"
EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
