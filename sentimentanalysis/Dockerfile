FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/classic-watchdog:0.1.4 as watchdog
FROM --platform=${TARGETPLATFORM:-linux/amd64} python:3.11

RUN mkdir -p /home/app

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Add non root user
RUN addgroup --system app \
  && adduser --system --ingroup app app --home /home/app
RUN mkdir -p /home/app
RUN chown app /home/app

WORKDIR /home/app


USER app
COPY requirements.txt   .
RUN pip install --user -r requirements.txt
RUN python3 -m textblob.download_corpora lite

COPY handler.py .
ENV fprocess="python3 ./handler.py"

ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
