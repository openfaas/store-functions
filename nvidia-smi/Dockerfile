FROM ghcr.io/openfaas/classic-watchdog:0.3.2 AS watchdog

FROM ubuntu:24.04

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

USER 1000

ENV fprocess="nvidia-smi"

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1
CMD ["fwatchdog"]
