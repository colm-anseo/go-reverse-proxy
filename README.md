# go-reverse-proxy
http reverse-proxy written in go.

Can connect to a TLS or non-TLS http (remote) address and proxy that to a (local) address using TLS or non-TLS.

## Install

```sh
go get github.com/colm-anseo/go-reverse-proxy
```
installs `go-reverse-proxy` in `GOBIN` directory (e.g. `${HOME}/go/bin` or `${GOPATH}/bin`)

## Run

### Locally

tunnel non-TLS `http` traffic over `https`:

```sh
REMOTE_ADDR="http://127.0.0.1:80" \
LOCAL_ADDR=":443" \
LOCAL_CERT_FILE="host.crt" \
LOCAL_CERT_KEY="host.key" \
    go-reverse-proxy
```

### Remote

If the target to be tunneled uses TLS, this connection can optionally be verified with a cert-file & name:

```sh
REMOTE_ADDR="https://some.example.com:443" \
REMOTE_CERT_FILE="example.crt" \
REMOTE_NAME="some.example.com" \
LOCAL_ADDR=":443" \
LOCAL_CERT_FILE="host.crt" \
LOCAL_CERT_KEY="host.key" \
    go-reverse-proxy
```

## Cross compiling

To compile the latest committed  version (`@master`) for Linux `x86_64`:
```sh
GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go get github.com/colm-anseo/go-reverse-proxy@master
```

binary will be here:
```sh
${HOME}/go/bin/linux_amd64/go-reverse-proxy
```
Note: `GOBIN` (used to set the binary destination) is not available with cross-compiles.
Hence the need to locate the binary in the `GOOS_GOARCH` directory. 




