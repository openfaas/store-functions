# store-functions

Common CLIs packaged in containers with a HTTP interface

About these functions:

* They are multi-arch - for `armhf`, `arm64` and `x86_64`
* They are built with `faas-cli`
* They are built and published to GHCR as container images to avoid the Docker Hub's rate limits

You may also like the bash template for OpenFaaS, so you can perform multiple tasks with CLIs or call your scripts over HTTPS.

* [Bash streaming templates](https://github.com/alexellis/openfaas-streaming-templates#example-with-bash)

## Using the functions

### With OpenFaaS

You can deploy them through openfaas on Kubernetes, or openfaas on a VM with [faasd](https://github.com/openfaas/faasd).

```bash
faas-cli store deploy curl

faas-cli store deploy nodeinfo

faas-cli deploy --image ghcr.io/openfaas/alpine:latest \
  --name env --fprocess="env"
```

Example usage:

```bash
faas-cli store deploy nmap
echo -n "-sP 192.168.0.0/24" faas-cli invoke nmap
```

### With Docker

Or run them ad-hoc with Docker:

```bash
docker run -p 8080:8080 --name hey \
  -d ghcr.io/openfaas/hey:latest

curl http://127.0.0.1:8080 -d "-c 5 -n 1000 https://your-service.com/"
```

## Included functions

* `alpine` - a base for running built-in bash or busybox commands like `env` (use it to debug headers) or `wc -l` to count text within a body
* `curl` - debug outgoing networking or internal services
* `figlet` - print ASCII logso
* `hey` - run a load-test against a HTTP API, website, or function with [hey](https://github.com/rakyll/hey)
* `nmap` - scan a network range
* `nodeinfo` - debug auto-scaling, find the container's memory, CPU information and uptime
* `sleep` - debug timeouts or async by sleeping for a set duration
* `sentimentanalysis` - use Python's textblob library to find out if a statement is positive or negative
* `shasum` - generate a SHA for a given input

Other functions: `imagemagick`, `ffmpeg`.

## License

MIT
