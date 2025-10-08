# hexbook

## client

build image with buildx.

```sh
# install if not available
sudo apt-get install docker-buildx-plugin

docker buildx version
```

## server

_dev_

run for local module.

```sh
# before publishing
go mod edit -replace github.com/hexbook/auth=./auth
```
